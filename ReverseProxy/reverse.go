package ReverseProxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/Lornzo/ReverseProxyServer/configs"
)

type reverse struct{}

func New() (r *reverse) {
	r = &reverse{}
	return
}

func (thisObj *reverse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		domain  string = r.Host
		service configs.Service
		err     error
	)
	if service, err = configs.GetService(domain); err == nil {
		director := func(req *http.Request) {
			req.URL.Scheme = service.Schema
			req.URL.Host = service.Host
		}
		proxy := &httputil.ReverseProxy{
			Director:     director,
			ErrorHandler: thisObj.errHandler,
		}
		proxy.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (thisObj *reverse) errHandler(w http.ResponseWriter, r *http.Request, err error) {}
