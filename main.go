package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/xoebus/shutdown"
)

const query = "org:cloudfoundry org:pivotal-cf-experimental org:pivotal-cf org:cloudfoundry-incubator unsafe extension:yml extension:erb path:jobs NOT rejoin"

var bold = color.New(color.Bold)

func main() {
	ctx := shutdown.WithShutdown(context.Background())
	client := github.NewClient(nil)

	res, _, err := client.Search.Code(ctx, query, nil)
	if err != nil {
		log.Fatal(err)
	}
	for i, result := range res.CodeResults {
		// Simple heuristic for false positives.
		if !strings.HasPrefix(result.GetName(), "bpm") {
			continue
		}
		display(i, result)
	}
}

func display(idx int, res github.CodeResult) {
	repo := res.GetRepository().GetName()
	path := res.GetPath()
	url := res.GetHTMLURL()
	fmt.Printf("[%d] %s (%s)\n    %s\n\n", idx, bold.Sprint(repo), path, url)
}
