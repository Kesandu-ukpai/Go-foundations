package main

import(

	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex
	
func count(wg *sync.WaitGroup){

	mu.Lock()
	defer wg.Done()
	defer mu.Unlock()

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

/**we want to create a lock for in count
establish the type 

var mu sync.Mutex

the use it

func count(wg *sync.WaitGroup){
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()

	for i:=0; i<1000000; i++ {
		counter++
	}
}


*/