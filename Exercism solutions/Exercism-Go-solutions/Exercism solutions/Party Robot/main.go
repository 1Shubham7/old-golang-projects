package main

import (
	// "strconv"
    "fmt"
)

func main (){
	a := AssignTable("Christiane", 4, "Frank", "on the left", 23.7834298)
	fmt.Print(a)
}
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	myTable := fmt.Sprint(table)
if (table / 100 < 1 && table/10 >1){
	myTable = "0" + myTable 
}
if (table/10 <1){
	myTable = "00" + myTable 
}

tableMsg := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, myTable, direction, distance, neighbor) 
return tableMsg
}