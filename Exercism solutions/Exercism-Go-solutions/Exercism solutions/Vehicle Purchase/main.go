package main

import (
	"fmt"
)

func main(){
	vehicle := CalculateResellPrice(1000, 1)
	fmt.Println(vehicle)
}

func NeedsLicense(kind string) bool {
    if (kind == "car" || kind == "truck"){
        return true
    }
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var a string 
	if (option1 < option2){
        a = option1
    } else {
    	a = option2
    }
	b := fmt.Sprintf("%v is clearly the better choice.", a)
    return b
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var b float64
    if (age <3 && age >=0){
        b = (0.8 * (originalPrice))
        return b
    }
        

	if (age >= 3 && age <10){
        b= (0.7 * (originalPrice))
        return b
    }
if (age >= 10){
    b=(0.5 * originalPrice)
	return b
}
	return 0
}