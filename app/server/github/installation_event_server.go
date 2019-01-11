package github

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	github_pb "github.com/izumin5210/ghrsync/api/github"
)

// InstallationEventServiceServer is a composite interface of github_pb.InstallationEventServiceServer and grapiserver.Server.
type InstallationEventServiceServer interface {
	github_pb.InstallationEventServiceServer
	grapiserver.Server
}

// NewInstallationEventServiceServer creates a new InstallationEventServiceServer instance.
func NewInstallationEventServiceServer() InstallationEventServiceServer {
	return &installationEventServiceServerImpl{}
}

type installationEventServiceServerImpl struct {
}

func (s *installationEventServiceServerImpl) CreateInstallationEvent(ctx context.Context, req *github_pb.CreateInstallationEventRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
