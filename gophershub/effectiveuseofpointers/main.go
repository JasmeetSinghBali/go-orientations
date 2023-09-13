package main

import "fmt"

type User struct {
	email    string
	username string
	age      int
	file     []byte // ?? small/medium/large ??
}

// 📌 This is a getter type function that just returns the email of the User type struct
// hence, no use of pointer
// 💡 will be x amount of bytes => sizeOf(user)
// 💫 NOTE- though we are only accesing the email still the entire user will be copied into this function with may be the size of file is 1gb + sizeOf(email)....
func (u User) Email() string {
	return u.email
}

// 📌 This is a setter type function
// hence, use of pointer
// 💡 only the 8-bytes size pointer gets copied here unlike Email() function
func (u *User) updateEmail(email string) {
	u.email = email
}

func main() {
	user := User{
		email: "noice@bar.com",
	}
	user.updateEmail("noice@foo.com")
	fmt.Println(user.Email())
}
