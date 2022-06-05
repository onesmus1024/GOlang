package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("concurrency testing")

	 k := Increment{n: 0}
	 k.Lock()
	 k.n++
	 k.Unlock()
	 go k.Inc()
	 time.Sleep(time.Second)
	 fmt.Println(k.n)
	 
}

type Increment struct{
	n int
	sync.Mutex
}

func (i *Increment) Inc() {
	i.Lock()
	i.n++
}