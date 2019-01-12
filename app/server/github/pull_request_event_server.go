package github

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	github_pb "github.com/izumin5210/ghrsync/api/github"
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
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
