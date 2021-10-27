package commands

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	fmt.Println("vim-go")
	res, err := git.PlainClone("./foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
