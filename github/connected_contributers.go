package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"
	"stay/graph"
)

// Discover GitHub users that are strongly connected via their Pull-Request relationships.
// For example if UserA requests a PR review from UserB and UserB requests a review in another PR from UserC and then
// UserC requests a review from UserA, that completes a cycle and users A, B and C are said to be strongly connected
// components in the graph.
//
// Provide the repository owner, name, auth token and the number of past Pull-Requests to look at
func FindConnectedUsers(owner string, repo string, auth string, nPrs int) map[int][]string {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: auth},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	listOpts := &github.PullRequestListOptions{State: "all", ListOptions: github.ListOptions{PerPage: 100}}
	var allPrs []*github.PullRequest
	for {
		prs, resp, err := client.PullRequests.List(ctx, owner, repo, listOpts)
		if err != nil {
			break
		}
		allPrs = append(allPrs, prs...)
		if resp.NextPage == 0 || len(allPrs) >= nPrs {
			break
		}
		listOpts.Page = resp.NextPage
	}
	fmt.Println(len(allPrs))
	return findStronglyConnected(allPrs)
}

func findStronglyConnected(prs []*github.PullRequest) map[int][]string {
	g := graph.NewGraph()
	for _, pr := range prs {
		from, to := authorAndReviewers(*pr)
		g.Push(from)
		for _, t := range to {
			g.Push(t)
			if err := g.Connect(from, t); err != nil {
				fmt.Println(err)
			}
		}
	}
	return graph.FindStronglyConnectedComponents(g)
}

func authorAndReviewers(pr github.PullRequest) (from string, to []string) {
	assignees := make([]string, 0)
	for _, a := range pr.RequestedReviewers {
		assignees = append(assignees, *a.Login)
	}
	return *pr.User.Login, assignees
}
