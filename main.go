package main

import (
	"fmt"
	"sync"
)

var exclude = []string{
	"node_modules",
	".go",
}

func main() {
	args := getArgs()
	src, err := NewSrc(args.Path)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, repo := range src.repos {
		wg.Add(1)
		repo := repo
		go func() {
			// pull
			res, err := repo.Git("pull")
			fmt.Println("PULL", repo.id)
			fmt.Println(err)
			fmt.Println(res)

			// pull submodules
			res, err = repo.Git("submodule", "foreach", "git", "pull")
			fmt.Println("PULL", repo.id)
			fmt.Println(err)
			fmt.Println(res)

			// push
			res, err = repo.Git("push")
			fmt.Println("PUSH", repo.id)
			fmt.Println(err)
			fmt.Println(res)
			wg.Done()
		}()
	}
	wg.Wait()
}
