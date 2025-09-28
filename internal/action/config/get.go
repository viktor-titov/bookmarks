package config

type ConfigRepositoryForGet interface {
	Get(key string) (string, error)
}

type Get struct {
	repo ConfigRepositoryForGet
}

func NewGet(repo ConfigRepositoryForGet) *Get {
	return &Get{
		repo: repo,
	}
}

func (s *Get) Do(key string) (string, error) {
	return s.repo.Get(key)
}
