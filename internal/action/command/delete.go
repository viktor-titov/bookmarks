package command

type Delete struct {
	repo CommandRepository
}

func NewDelete(repo CommandRepository) *Delete {
	return &Delete{repo: repo}
}

func (d *Delete) Do(key string) error {
	return d.repo.Delete(key)
}
