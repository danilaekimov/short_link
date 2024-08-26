package storage

import (
	"errors"
	"log"

	"github.com/danilaekimov/short_link/internal/util"
)

type webResourse struct {
	URL     string
	counter int32
}

type URLStorage struct {
	ShortURLs map[string]webResourse
}

func NewStorage() *URLStorage {
	return &URLStorage{
		ShortURLs: make(map[string]webResourse),
	}
}

func (st URLStorage) SaveURL(URL string) (shortURL string) {

	shortURL = util.GetMD5Hash(URL, 4)
	st.ShortURLs[shortURL] = webResourse{
		URL:     URL,
		counter: 0}
	log.Println(st.ShortURLs[shortURL])
	return

}

func (st URLStorage) GetByShortURL(requiredURL string) (shortURL string, err error) {
	webResourse, ok := st.ShortURLs[requiredURL]
	if ok {
		webResourse.counter += 1
		st.ShortURLs[requiredURL] = webResourse

		log.Println(st.ShortURLs[requiredURL])

		return webResourse.URL, nil
	} else {
		err = errors.New(requiredURL + " does not exist")
		return
	}
}
