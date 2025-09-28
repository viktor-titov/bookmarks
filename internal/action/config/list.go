package config

type ConfigRepository interface {
	List() map[string]string
}

type List struct {
	repo ConfigRepository
}

func NewList(repo ConfigRepository) *List {
	return &List{
		repo: repo,
	}
}

func (l *List) Do() map[string]string {
	return l.repo.List()
}
