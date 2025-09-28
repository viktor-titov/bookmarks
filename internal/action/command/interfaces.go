package command

type CommandRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	List() (map[string]string, error)
}

type ConfigRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
	List() map[string]string
}
