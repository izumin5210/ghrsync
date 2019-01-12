package github

import (
	"context"
	"testing"


	github_pb "github.com/izumin5210/ghrsync/api/github"
)

func Test_PullRequestEventServiceServer_CreatePullRequestEvent(t *testing.T) {
	svr := NewPullRequestEventServiceServer()

	ctx := context.Background()
	req := &github_pb.CreatePullRequestEventRequest{}

	resp, err := svr.CreatePullRequestEvent(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
