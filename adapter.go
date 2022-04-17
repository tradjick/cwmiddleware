package cwmiddlware

import "net/http"

type Adapter func(h http.Handler) http.Handler

func AdaptArray(h http.Handler, a []Adapter) http.Handler {
	for _, ae := range reverseArray(a) {
		h = ae(h)
	}
	return h
}

func Adapt(h http.Handler, as ...Adapter) http.Handler {
	return AdaptArray(h, as)
}
