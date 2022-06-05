package main

import "fmt"

type I interface {

	talk() string
	}
type cat struct {}

func (c *cat) talk() string {
	return "meow"
}

func main() {
	fmt.Println("interfaec implementation")
	fmt.Println(check(struct{y int}{y:1}))

}
func check(i interface{}) string{
	a,ok := i.(I)
	if ok {
		return fmt.Sprintf(a.talk())
	}
	return "not ok"
}