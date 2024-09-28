package main

import (
    "strconv"
    "fmt"
    "strings"
    "unicode"
)

func Valid(id string) bool {
	sum := 0
	id = strings.Replace(id, " ", "", -1)

    for _, k := range id {
        if unicode.IsLetter(k){
            return false
        }
    }

	if len(id) <= 1 {
		return false
	}

    runes := []rune(id)
    length := len(runes)
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    reversedId := string(runes)

	for i := len(reversedId)-1; i >=0; i-- {
		n, _ := strconv.Atoi(string(reversedId[i]))
		if i%2 != 0 {
			fmt.Printf("n before anything: %d\n", n)
			n = n * 2
			if n > 9 {
				fmt.Printf("n after multiplying: %d\n", n)
				n = (n - 9)
			} else {
				fmt.Printf("n after multiplying is same: %d\n", n)
			}
		} else {
			fmt.Printf("n is unchanged: %d\n", n)
		}
		sum = sum + n
		fmt.Printf("\n sum after adding %d: %d\n", n, sum)
	}

	fmt.Printf("\n sum : %d", sum)

	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}


func main() {
	fmt.Println(Valid("59"))
}
