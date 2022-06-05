package main

import "fmt"


type point struct {	
	x, y int
}

func (p *point) add(a,b int) {

	p.x += a
	p.y += b
}

func (p *point) display() string {
	return fmt.Sprintf("x: %d, y: %d", p.x, p.y)
}

func main() {
	p := point{1,2}
	p.add(3,4)
	fmt.Println(p.display())
}