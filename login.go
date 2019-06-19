package main

import (
	"fmt"
)

var(
	name = "james_sam"
	pin = 1234
	email = "jumachambala@gmail.com"
)


func main() {
	    fmt.Println("Enter Your Name . .")
		fmt.Scanln(&name)
		fmt.Println("Enter Your Email . .")
		fmt.Scanln(&email)
		fmt.Println("Enter Your Password . .")
     	fmt.Scanln(&pin)
	if pin == 1234 && email == "jumachambala@gmail.com" && name == "james_sam" {
		fmt.Println("Your Successful Log In . .")

	}else if name == "james_sam" && email != "jumachambala@gmail.com" && pin != 1234 || name != "james_sam" && email == "jumachambala@gmail.com" && pin != 1234 || name != "james_sam" && email != "jumachambala@gmail.com" && pin == 1234 || name == "james_sam" && email == "jumachambala@gmail.com" && pin != 1234 || name != "james_sam" && email == "jumachambala@gmail.com" && pin == 1234 || name != "james_sam" && email != "jumachambala@gmail.com" && pin != 1234 || name == "james_sam" && email != "jumachambala@gmail.com" && pin == 1234 {




		/*name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||

		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||
		name != "james sam" && email != "jumachambala@gmail.com" && pin != 1234 ||*/






		fmt.Println("404 Access denied wrong password . .")

	}

}
