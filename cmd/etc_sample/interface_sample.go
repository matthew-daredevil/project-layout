package etc_sample

import (
	"fmt"
)

type SampleInterface interface {
	SampleMethod()
}
type Human interface {
	Walk() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Walk() string {
	return fmt.Sprintf("%s can walk", s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func Etc_sample() {
	var sample_interface SampleInterface
	fmt.Println(sample_interface)

	s := Student{Name: "John", Age: 20}
	var h Human

	h = s
	fmt.Println(h.Walk())
	// fmt.Println(h.GetAge()) // ERROR

	// // User struct from the user package
	// u := user.User{Name: "John"}
	// fmt.Println(u)

	// // Assigning User to Human interface
	// h := Human(u)
	// fmt.Println(h)

	// // Calling methods from the interface
	// h.Walk()
	// h.Talk()
}

// Assuming the 'user' package has the following content:
// package user
//
// import "fmt"
//
// type User struct {
//     Name string
// }
//
// func (u User) Walk() {
//     fmt.Println("Walking")
// }
// func (u User) Talk() {
//     fmt.Println("Talking")
// }
