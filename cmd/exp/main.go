package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}
type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Smith",
		Bio:  `<script>alert("Haha, you've been done");<script>`,
		Age:  123,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
