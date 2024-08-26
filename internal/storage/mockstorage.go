package storage

var MockURLs = []string{
	"http://ya.ru/", "https://yandex.ru/", "https://mail.ru/"}

func GenMockData() *URLStorage {
	storage := NewStorage()
	for _, url := range MockURLs {
		storage.SaveURL(url)
	}
	return storage
}
