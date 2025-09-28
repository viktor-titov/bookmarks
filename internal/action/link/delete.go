package link

type Delete struct {
	repo LinkRepository
}

func NewDelete(repo LinkRepository) *Delete {
	return &Delete{repo: repo}
}

func (d *Delete) Do(key string) error {
	return d.repo.Delete(key)
}
