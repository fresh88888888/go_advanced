package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"umbrella.github.com/advanced_go/advanced_5/github/git"
)

var apiToken = os.Getenv("GETHUB_API_TOKEN")

func main() {
	ctx := context.Background()
	client := git.NewClient(ctx, apiToken)
	repos, _, err := client.Repositories.List(ctx, "fresh888888")
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Println(repo)
	}
}
