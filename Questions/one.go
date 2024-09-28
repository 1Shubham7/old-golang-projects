package main

import (
	"fmt"
)

func main(){
	arr := []int32{7 ,69 ,2 ,221 ,8974}
	miniMaxSum(arr)
}

func miniMaxSum(arr []int32) {
	// var max int
	
	// for a:=1; a<5; a++ {
	// 	if (arr[a-1] > arr[a]){ //1
	// 		temp := arr[a-1]
	// 		arr[a-1] = arr[a]
	// 		arr[a] = temp
	// 	}
		
	// }


	for i := 0; i < 4; i++ {
        for j := 0; j < 4-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }


	for i:=0;i<=4;i++{
		fmt.Printf("\n%d\n",arr[i])
	}
	fmt.Printf("---------------------------------")

	maxSum := arr[1]+arr[2]+arr[3]+arr[4]
	minSum := arr[0]+arr[1]+arr[2]+arr[3]
	fmt.Printf("%d %d", minSum, maxSum)
}