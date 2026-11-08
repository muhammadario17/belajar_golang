package main

import "fmt"

//Definisi model struct
type User struct {
	Username string
	Email    string
}

//Method dengan Pointer Receiver (u *User)
//Wajib pakai tanda * agar perubahan email benar-benar tersimpan permanen

func (u *User) GantiEmail(emailBaru string) {
	u.Email = emailBaru
}

func main() {
	user1 := &User{
		Username: "muhammadario17",
		Email:    "muhammadario017@gmail.com",
	}

	user1.GantiEmail("ario_baru@gmail.com")

	fmt.Println("Email Terbaru: ", user1.Email)
}
