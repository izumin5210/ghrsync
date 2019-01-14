package domain

type RepositorySlug string

type Installation struct {
	ID   uint32
	Slug RepositorySlug
}
