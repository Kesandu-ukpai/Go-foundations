package main

import(

	"fmt"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex

	
func count(wg *sync.WaitGroup){

	
	defer wg.Done()
	

	for i:=0; i<1000000; i++ {
		mu.Lock()
		
		counter++

		mu.Unlock()
	}
}

func main(){
start := time.Now()
var wg sync.WaitGroup

for i:=0; i<10; i++{
wg.Add(1)
go count(&wg)

}

wg.Wait()
duration := time.Since(start)
	
	fmt.Println("Counter:", counter)
	fmt.Println("Time taken:", duration)
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