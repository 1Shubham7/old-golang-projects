package cal

import "errors"

func Divide(a ,b int) (int, error){
	var c int = a/b

	if (b ==0 ){
		return 0, errors.New("division by zero")
	}
	return c,nil
}