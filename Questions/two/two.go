package main

import (
	"fmt"
)

func main(){


	arr := []int32{13,23,13,13,12}
	gg(arr)
}

func gg(arr []int32){
	n := len(arr)
    var posi int = 0
    var negi int =0
    var zero int =0
    
    for i:=0; i<n-1;i++{
        if (arr[i]>0){
            posi++
            fmt.Printf("posi count is : %d", posi)
        }
        if (arr[i]<0){
            negi++
        }
        if (arr[i]==0){
            zero++
        }
    }

    posirato:= posi/n
    negirato:= negi/n
    zerorato:= zero/n
    fmt.Printf("%d %d %d", posirato, negirato, zerorato)
}