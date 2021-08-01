package handlers

import (
	"github.com/Miraddo/SimpleShortlink/pkg/shorter"
	"go.uber.org/zap"
	"net/http"
)

type HTTPHandler struct {
	Shorter *shorter.ShorterFunc
	Logger  *zap.SugaredLogger
}

func (hh *HTTPHandler) MainUrlFunc(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")

	rt, err := hh.Shorter.MainUrl(key)

	if key == "" || err != nil || rt == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte("Key Is Not Correct"))
		if err != nil {
			return
		}
		hh.Logger.Errorf("Key is not Correct!")
		return
	}

	//writer.Header().Set("Access-Control-Allow-Origin", "*")

	_, err = writer.Write([]byte(rt))
	if err != nil {
		hh.Logger.Errorf("Could not disply string with Write method in writer!")
		return
	}
	//http.go.Redirect(writer, request, rt, 301)
	//log.Println(rt)
}

func (hh *HTTPHandler) ShortUrlFunc(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Query().Get("url")
	result, err := hh.Shorter.ShortUrl(url)

	if AvailableUrl(url) != true || url == "" || result == "" || err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte("Url Is Not Correct"))
		if err != nil {
			return
		}
		hh.Logger.Errorf("Url is not Correct!")
		return
	}

	rt := "127.0.0.1:8080/url?key=" + result

	_, err = writer.Write([]byte(rt))

	if err != nil {
		hh.Logger.Errorf("Could not disply string with Write method in writer!")
		return
	}
}
