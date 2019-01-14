package domain

import "context"

type SHA string
type Ref string
type BranchName string

func (bn BranchName) Ref() Ref {
	return Ref("heads/" + string(bn))
}

type Branch struct {
	Name BranchName
	Ref  Ref
	SHA  SHA
}

func (b *Branch) GetRef() Ref {
	if b.Ref != "" {
		return b.Ref
	}
	return b.Name.Ref()
}

type PullRequest struct {
	Number uint32
	HEAD   *Branch
}

type PullRequestRepository interface {
	Create(ctx context.Context, slug RepositorySlug, head, base *Branch, title, body string) error
}
