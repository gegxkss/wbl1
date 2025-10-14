package main

import (
	"fmt"
)

// структура Human
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// метод структуры Human
func (h Human) GetAge() int {
	return h.Age

}

func (h Human) GetName() string {
	return h.FirstName + " " + h.LastName

}

// структура Action с встроенной структурой Human
type Action struct {
	Human
	ActionType string
}

// метод структуры Action
func (a Action) GetInfo() string {
	return fmt.Sprintf("Name - %s, Age - %d, Action - %s", a.GetName(), a.GetAge(), a.ActionType)
}

func main() {
	newPerson := Action{
		Human: Human{
			FirstName: "Oksana",
			LastName:  "Andreeva",
			Age:       23},
		ActionType: "walking",
	}
	fmt.Println(newPerson.GetInfo())
}
