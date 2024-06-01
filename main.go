package main

import (
	"fmt"
	"nits-tips-api/db"
)

func main() {
	db.NewDB()
	fmt.Printf("Hello, World!\n")
}
