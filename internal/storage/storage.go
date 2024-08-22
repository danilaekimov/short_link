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

type URLsStorage struct {
	shortURLs map[string]webResourse
}

func NewStorage() *URLsStorage {
	return &URLsStorage{
		shortURLs: make(map[string]webResourse),
	}
}

func (st URLsStorage) SaveURL(URL string) (shortURL string) {

	shortURL = util.GetMD5Hash(URL, 4)
	st.shortURLs[shortURL] = webResourse{
		URL:     URL,
		counter: 0}
	log.Println(st.shortURLs[shortURL])
	return

}

func (st URLsStorage) GetByShortURL(requiredURL string) (shortURL string, err error) {
	webResourse, ok := st.shortURLs[requiredURL]
	if ok {
		webResourse.counter += 1
		st.shortURLs[requiredURL] = webResourse

		log.Println(st.shortURLs[requiredURL])

		return webResourse.URL, nil
	} else {
		err = errors.New(requiredURL + " does not exist")
		return
	}
}
