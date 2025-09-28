package command

type List struct {
	repo CommandRepository
}

func NewList(repo CommandRepository) *List {
	return &List{repo: repo}
}

func (a *List) Do() (map[string]string, error) {
	return a.repo.List()
}
