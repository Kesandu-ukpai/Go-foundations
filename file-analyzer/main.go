package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("big.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return
}
