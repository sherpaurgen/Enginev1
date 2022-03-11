package main

import "fmt"

type Dog struct {
	sound string
	legs int
	speed int
}
type Car struct {
	sound string
	speed int
}
type Check interface {
	CheckSound() string
	CheckSpeed() int
}
func main(){
	car1:=Car{sound: "vrooom",speed:999}
	tommy:=Dog{sound: "woof",speed: 10,legs: 4}
	var checkInterface1 Check = car1
	fmt.Println(checkInterface1.CheckSpeed())
	var checkInterface2 Check = tommy
	fmt.Println(checkInterface2.CheckSpeed())

}
func (d Dog) CheckSound() string{
	return "woof"
}
func (c Car) CheckSound() string {
	return "vroom"
}
func (d Dog) CheckSpeed() int {
	return 10
}
func (c Car) CheckSpeed() int {
	return 999
}


