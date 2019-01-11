package github

import (
	"context"
	"testing"


	github_pb "github.com/izumin5210/ghrsync/api/github"
)

func Test_InstallationEventServiceServer_CreateInstallationEvent(t *testing.T) {
	svr := NewInstallationEventServiceServer()

	ctx := context.Background()
	req := &github_pb.CreateInstallationEventRequest{}

	resp, err := svr.CreateInstallationEvent(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
