package command

type Get struct {
	repo CommandRepository
}

func NewGet(repo CommandRepository) *Get {
	return &Get{repo: repo}
}

func (g *Get) Do(key string) (string, error) {
	return g.repo.Get(key)
}
