package main

import "fmt"

// ID не стал выносить в person, т.к. посчитал, что у пользователей и сотрудников должны быть свои идентификаторы
type person struct {
	Name  string
	Addss string
	Phone string
}

type user struct {
	ID int
	person
}
type employee struct {
	ID int
	person
}

func main() {
	u := user{
		ID: 123,
		person: person{
			Name:  "Vasya",
			Addss: "Street",
			Phone: "111-11-11",
		},
	}

	e := employee{
		ID: 234,
		person: person{
			Name:  "Petya",
			Addss: "Boulevard",
			Phone: "222-22-22",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
