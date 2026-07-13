package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Sound string
}

func main() {
	fmt.Println("Defining Functions As Methods Of Custom Types")

	dogAlpha := Dog{"Poodle", "Woof"}
	// fmt.Printf("The %v Says %v!\n", dogAlpha.Breed, dogAlpha.Sound)

	dogAlpha.Speak()

	dogAlpha.Sound = "Arf"
	dogAlpha.Speak()

	println(dogAlpha.SpeakThreeTime())
}

func (dogAlpha Dog) Speak() {
	fmt.Printf("The %v Says %v!\n", dogAlpha.Breed, dogAlpha.Sound)
}

func (dogAlpha Dog) SpeakThreeTime() string {
	return fmt.Sprintf("%v! %v! %v!", dogAlpha.Sound, dogAlpha.Sound, dogAlpha.Sound)
}
