package main

import (
	"fmt"
)

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	key := "qwertyuiopasdfghjklzxcvbnm"
	pasangan := map[string]string{}

	for i := 0; i < 26; i++ {
		pasangan[string(97+i)] = string(key[i])
	}

	for i := range s.name {
		if value, exist := pasangan[string(s.name[i])]; exist {
			nameEncode += value
		} else {
			nameEncode += string(s.name[i])
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	key := "qwertyuiopasdfghjklzxcvbnm"
	pasangan := map[string]string{}

	for i := 0; i < 26; i++ {
		pasangan[string(key[i])] = string(97+i)
	}
	for i := range s.name {
		if value, exists := pasangan[string(s.name[i])]; exists {
			nameDecode += value
		} else {
			nameDecode += string(s.name[i])
		}
	}

	return nameDecode
}

func main() {
	var menu int
  	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())

	} else {
		fmt.Println("Menu doesn't exist!")
	}
}