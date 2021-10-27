package main

import (
	"fmt"

	"github.com/hoenn/clone/cmd/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
	}
}
