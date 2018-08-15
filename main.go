package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/xoebus/shutdown"

	"github.com/xoebus/osha/querybuild"
)

var (
	query = querybuild.Build(
		querybuild.Org("cloudfoundry"),
		querybuild.Org("cloudfoundry-incubator"),
		querybuild.Org("pivotal-cf"),
		querybuild.Org("pivotal-cf-experimental"),
		querybuild.Filename("bpm.yml"),
		querybuild.Query("unsafe"),
	)

	bold = color.New(color.Bold)
)

func main() {
	ctx := shutdown.WithShutdown(context.Background())
	client := github.NewClient(nil)

	res, _, err := client.Search.Code(ctx, query, nil)
	if err != nil {
		log.Fatal(err)
	}
	// BUG(cb): pagination is not handled
	for i, result := range res.CodeResults {
		display(i+1, result)
	}
}

func display(idx int, res github.CodeResult) {
	repo := res.GetRepository().GetName()
	path := res.GetPath()
	url := res.GetHTMLURL()
	fmt.Printf("[%d] %s (%s)\n    %s\n\n", idx, bold.Sprint(repo), path, url)
}
