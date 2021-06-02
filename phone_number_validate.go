package main

import (
	"fmt"
	"regexp"
)

func main()  {
	phoneNumber := "01777188552a"
	reg := regexp.MustCompile(`^(01)[3-9][0-9]{8}$`)
	if !reg.MatchString(phoneNumber) {
		fmt.Println("phone number is not valid")
		return
	}

	fmt.Println("valid phn number")

}
