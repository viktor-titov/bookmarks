package models

// Config представляет конфигурацию приложения.
type Config struct {
	// PathBookmarks путь к файлу c ссылками.
	PathBookmarks string `json:"pathBookmarks"`
	// PathCommands путь к файлу c командами.
	PathCommands string `json:"pathCommands"`
	// BrowserApp приложение для открытия ссылок в браузере.
	BrowserApp string `json:"browserApp"`
}
