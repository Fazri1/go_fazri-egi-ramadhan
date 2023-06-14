package controllers

import (
	"praktikum/users"
	"strconv"
)

var usersData = users.Users


func GetUsersController() interface{} {
	return usersData
}

func GetUserController(id int) (string, interface{}) {
	var message string
	var user interface{}
	index := BinarySearch(usersData, id)

	if index == -1 {
		message = "user not found"
		user = nil
	} else {
		message = "success get user by id: " + strconv.Itoa(id)
		user = usersData[index]
	}

	return message, user
}

func DeleteUserController(id int) (string, interface{}) {
	var message string
	var user interface{}
	
	index := BinarySearch(usersData, id)

	if index == -1 {
		message = "user not found"
		user = nil
	} else {
		message = "success delete user with id: " + strconv.Itoa(id)
		user = usersData[index]
		usersData = append(usersData[:index], usersData[index+1:]...)
	}

	return message, user
}

func UpdateUserController(id int, updateUser users.User) (string, interface{}) {
	var message string
	var user interface{}
	
	index := BinarySearch(usersData, id)

	if index == -1 {
		message = "user not found"
		user = nil
	} else {
		message = "success update user"
		user = updateUser
		usersData[index] = updateUser
	}

	return message, user
}

func CreateUserController(user users.User) (string, interface{}){
	var message string

	if len(usersData) == 0 {
		user.Id = 1
	} else {
		newId := usersData[len(usersData)-1].Id + 1
		user.Id = newId
	}

	switch {
	case user.Name == "":
		message = "name cannot be empty"
	case user.Email == "":
		message = "email cannot be empty"
	case user.Password == "":
		message = "password cannot be empty"
	default:
		message = "success create user"
		usersData = append(usersData, user)
	}

	return message, user
}

func BinarySearch(user []users.User, id int) int {
	left := 0
	right := len(user) - 1

	for left <= right {
		mid := (left + right) / 2
		if user[mid].Id == id {
			return mid
		} else if user[mid].Id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
