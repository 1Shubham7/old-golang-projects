package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	database "github.com/1shubham7/jwt/database"
	helper "github.com/1shubham7/jwt/helpers"
	models "github.com/1shubham7/jwt/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
func HashPassword(password string) string {
	hashed, err:=bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!=nil{
		log.Panic(err)
	}
	return string(hashed)
}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		check = false
		msg = fmt.Sprintf("email or password is incorrect.")
	}
	return check, msg
}

func SignUp()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		// this is used to validate, but what? see the User struct, and see those validate struct fields
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// we are using count to help us validate. if you find documents with the user email already
		// then count would be more than 0, and we can then handle that err
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone number"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
		}

		// by "c.BindJSON(&user)" user already have the information from the website user
		user.Created_at, _  = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _  = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
		
		// giving value that we generated to user
		user.Token = &token
		user.Refresh_token = &refreshToken

		// now let's insert it to the database
		resultInsertionNumber, insertErr :=  userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		// giving the user data to user variable
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// finding the user through email and if found, storing it in foundUser variable
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()

		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return 
		}

		// we need pointer to acess the origina user and foundUser,
		// if we only pass user and foundUser, it will create a new instance of user and foundUser
		isPasswordValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if isPasswordValid != true{
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id":foundUser.User_id}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}

// GetUsers can only be accessed by the admin.
func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		if err := helper.CheckUserType(c, "ADMIN"); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		//  setting how many records you want per page.
		// we are taking the recodePerPage from c and converting it to int
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))

		// if error or recordPerPage is less than 1, by default we will have 9 records per page
		if recordPerPage<1||err != nil {
			recordPerPage = 9
		}

		// this is just like page number
		page, err1 := strconv.Atoi(c.Query("page"))
		// we want to start with the page number 1 by default.
		if err1 !=nil || page < 1{
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}},
	}}
		// group all the data based on id, and then count them using $sum. then
		// pushing all the data to the root.
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}},
	}}}

		// in project stage we deside which data should go to the user and which not.
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})

		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing user items"})
		}

		// creating a slice called allUser and giving the result value
		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		} 
		c.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUserById() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id") // we are taking the user_id given by the user in json
		// with the help of gin.context we can access the json data send by postman or curl or user
		
		if err := helper.MatchUserTypeToUserId(c, userId); err != nil{

			//checking if the user in admin or not.
			// we will create that func in helper package.

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return 
			}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// if everything goes ok, pass the data of the user (UserModel.go)
		c.JSON(http.StatusOK, user)

	}
}
