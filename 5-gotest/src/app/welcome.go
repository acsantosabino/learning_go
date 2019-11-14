package main

import "fmt"

func welcome(name string) string{
	if(len(name) == 0){
		return "Be welcome!"
	} else {
		return fmt.Sprintf("Be welcome %v!", name)
	}
}