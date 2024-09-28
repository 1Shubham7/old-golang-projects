package main

import "fmt"

func main(){
	// a := InterestRate(1000)
	// fmt.Println(a)

	b := Interest(200.75)
	fmt.Println(b)
}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

    if (balance < 0){
        return 3.213
    } else if (balance>=0 && balance < 1000){
    	return 0.5
    } else if (balance>=1000 && balance < 5000){
    	return 1.621
    } else{
    	return 2.475
    }
    
}

// Interest calculates the interest for the provided balance.
// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    interest := float32(balance) + (InterestRate(balance) * float32(balance))
    if (balance < 0){
        y:= float64(interest)
    return (-y)
    } else {
    return float64(interest)
    }

}