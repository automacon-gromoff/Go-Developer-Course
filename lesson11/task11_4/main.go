package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {
	myErr := myFirstError{"моя ошибка"}
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	fmt.Printf("Ошибка \"%s\" была: %t", myErr, errors.Is(err3, myErr))
}
