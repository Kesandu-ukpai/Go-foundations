package main

import(

	"fmt"
	"sync"
)

var counter int
	
func count(wg *sync.WaitGroup){

	wg.Done()

	for i=0, i<1000, i++{
		counter++
	}
}

func main(){

var wg sync.WaitGroup

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)

wg.Add(1)
go count(&wg)
}