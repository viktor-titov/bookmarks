package config

type ConfigRepositoryForSet interface {
	Set(key, value string) error
}

type Set struct {
	repo ConfigRepositoryForSet
}

func NewSet(repo ConfigRepositoryForSet) *Set {
	return &Set{
		repo: repo,
	}
}

func (s *Set) Do(key, value string) error {
	return s.repo.Set(key, value)
}
