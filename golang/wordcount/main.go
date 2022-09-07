package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter String: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	arr := strings.Split(str, " ")

	hmap := make(map[string]int32)

	for i := 0; i < len(arr); i++ {

		hmap[arr[i]] += 1
	}
	fmt.Println(hmap)
}
