package main

import "testing"

func TestWelcome( t *testing.T ){
	
	result := welcome("")
	waiting_for := "Be welcome!"

	if (result != waiting_for){
		t.Errorf("welcome(\"\") failed, expected %v, got %v", waiting_for, result)
	}
	
	result = welcome("Zé")
	waiting_for = "Be welcome Zé!"

	if (result != waiting_for){
		t.Errorf("welcome(\"\") failed, expected %v, got %v", waiting_for, result)
	}
}