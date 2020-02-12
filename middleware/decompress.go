package middleware

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func TransparentGunzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// It only makes sense to support request body decoding for POST, PUT and PATCH methods
		if r.Method != http.MethodPost && r.Method != http.MethodPut && r.Method != http.MethodPatch {
			// Pass through to the next handler
			next.ServeHTTP(w, r)
			return
		}

		buf, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if http.DetectContentType(buf) == "application/x-gzip" {
			gr, err := gzip.NewReader(bytes.NewBuffer(buf))
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			r.Body = gr
		} else {
			r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		}

		next.ServeHTTP(w, r)
	})
}
