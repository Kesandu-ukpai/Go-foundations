package main

import(

	"fmt"
	"sync"
)

var counter int
	
func count(wg *sync.WaitGroup){

	defer wg.Done()

	for i:=0; i<1000000; i++ {
		counter++
	}
}

func main(){

var wg sync.WaitGroup

for i:=0; i<10; i++{
wg.Add(1)
go count(&wg)

}

wg.Wait()

fmt.Println(counter)
}