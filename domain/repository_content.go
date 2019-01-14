package domain

import "context"

type RepositoryContent struct {
	Path    string
	Content string
	SHA     SHA
}

type RepositoryContentRepository interface {
	List(ctx context.Context, path string, ref Ref) ([]*RepositoryContent, error)
	Get(ctx context.Context, path string, ref Ref) (*RepositoryContent, error)
	CreateOrUpdateAll(ctx context.Context, contents []*RepositoryContent, ref Ref) error
}
