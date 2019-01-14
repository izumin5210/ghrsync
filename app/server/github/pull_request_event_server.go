package github

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"

	github_pb "github.com/izumin5210/ghrsync/api/github"
	"github.com/izumin5210/ghrsync/domain"
)

// PullRequestEventServiceServer is a composite interface of github_pb.PullRequestEventServiceServer and grapiserver.Server.
type PullRequestEventServiceServer interface {
	github_pb.PullRequestEventServiceServer
	grapiserver.Server
}

// NewPullRequestEventServiceServer creates a new PullRequestEventServiceServer instance.
func NewPullRequestEventServiceServer() PullRequestEventServiceServer {
	return &pullRequestEventServiceServerImpl{}
}

type pullRequestEventServiceServerImpl struct {
}

func (s *pullRequestEventServiceServerImpl) CreatePullRequestEvent(ctx context.Context, req *github_pb.CreatePullRequestEventRequest) (*empty.Empty, error) {
	ev := req.GetPullRequestEvent()
	_ = &domain.Installation{
		ID:   ev.GetInstallation().GetId(),
		Slug: domain.RepositorySlug(ev.GetRepository().GetFullName()),
	}
	switch ev.GetAction() {
	case "opened":
		// TODO
	case "synchronized":
		// TODO
	case "closed":
		if ev.GetPullRequest().GetMerged() {
			// TODO
		}
	}
	return new(empty.Empty), nil
}
