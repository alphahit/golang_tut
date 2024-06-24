package main

import "fmt"

//1. We use pointers when we need to update a state
//pointer 8 bytes
//2. When we want want to optimize the memory for large objects. Pointer will
// be used when you have a big object but you want to copy only one thing and not whole object

type User struct {
	email string
	username string
	age int
	file []byte
}

func (u User) Email() string {
	return u.email
}
//The top and below functions are completely same
func Email(user User) string {
	return user.email
}//x amount of bytes => sizeOf(user)

//8 bytes
func (u *User) UpdateEmail(email string){
	u.email = email
}

func getUser() (*User, error) {
	return nil, fmt.Errorf("foo error")
}

func getUserTwo() (*User, error) {
	// Simulate a successful case
	return &User{
		email: "success@example.com",
	}, nil
}


func main() {
	// Create a User instance
	user := User{
		email: "prat@example.com",
	}
	
	// Print the initial email
	fmt.Println(user.Email())
	
	// Update the email using the pointer receiver method
	user.UpdateEmail("foo@example.com")
	
	// Print the updated email
	fmt.Println(user.Email())

	// Using the getUser function
	usr, err := getUser()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", usr.Email())
	}


	// Using the getUser function
	usr2, err := getUserTwo()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", usr2.Email())
	}
}