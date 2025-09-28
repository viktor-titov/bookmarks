package link

type List struct {
	repo LinkRepository
}

func NewList(repo LinkRepository) *List {
	return &List{repo: repo}
}

func (a *List) Do() (map[string]string, error) {
	return a.repo.List()
}
