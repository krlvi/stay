package main

import (
	"fmt"
	"os"
	"stay/github"
)

func main() {
	owner := os.Args[1]
	repo := os.Args[2]
	scc := github.FindConnectedUsers(owner, repo, os.Getenv("GH_TOKEN"), 300)
	for _, x := range scc {
		fmt.Println(x)
	}
}
