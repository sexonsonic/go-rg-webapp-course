package main

import (
	"encoding/json"
	"fmt"
)

var strJSson = ``

type User struct {
	FirstName 	string 		`json:"first_name"`
	LastName	string		`json:"last_name"`
	Age			int			`json:"age"`
	Hobbies		[]string	`json:"hobi`
	Arr			[]int		`json:"arr"`
	Address		Address		`json:"alamat"`
}

type Address struct {
	Street		string		`json:"jalan"`
	City		string		`json:"kota"`
	Province	string		`json:"provinsi"`
}

func main()  {
	bagas := User{}

	// decode strJson to struct user
	// use json.Unmarshall[]byte
	err := json.Unmarshal([]byte(strJSson), &bagas)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}

	fmt.Println("First Name:", bagas.FirstName)
	fmt.Println("Last Name:", bagas.LastName)

	fmt.Println("First Name:", bagas.FirstName)
	fmt.Println("Last Name:", bagas.LastName)
	fmt.Println("Age:", bagas.Age)
	fmt.Println("Hobbies:", bagas.Hobbies)
	fmt.Println("Arr:", bagas.Arr)
	fmt.Println("Address:", bagas.Address)
	fmt.Println()

	salsa := User{
		FirstName: "Salsa",
		LastName:  "Ainun",
		Age:       22,
		Hobbies:   []string{"membaca", "menulis"},
		Address: Address{
			Street:   "Jalan 2",
			City:     "Jakarta",
			Province: "DKI Jakarta",
		},
	}

	// encode struct User to JSON
	// use json.Marshal(struct)
	jsonSalsa, err := json.Marshal(salsa)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonSalsa))

}