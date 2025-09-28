package models

// Config представляет конфигурацию приложения.
type Config struct {
	// Path путь к файлу с данными и конфигурацией.
	Path string
	// BrowserApp приложение для открытия ссылок в браузере.
	BrowserApp string
}
