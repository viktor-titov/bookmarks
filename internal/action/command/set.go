package command

type Set struct {
	repo CommandRepository
}

func NewSet(repo CommandRepository) *Set {
	return &Set{repo: repo}
}

func (l *Set) Do(key, value string) error {
	return l.repo.Set(key, value)
}
