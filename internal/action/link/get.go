package link

type Get struct {
	repo LinkRepository
}

func NewGet(repo LinkRepository) *Get {
	return &Get{repo: repo}
}

func (g *Get) Do(key string) (string, error) {
	return g.repo.Get(key)
}
