package link

type Set struct {
	repo LinkRepository
}

func NewSet(repo LinkRepository) *Set {
	return &Set{repo: repo}
}

func (l *Set) Do(key, value string) error {
	return l.repo.Set(key, value)
}
