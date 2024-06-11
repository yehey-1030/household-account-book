package constants

import "net/http"

var (
	AllMethod = map[string]bool{
		http.MethodGet:     true,
		http.MethodHead:    true,
		http.MethodPost:    true,
		http.MethodPut:     true,
		http.MethodPatch:   true,
		http.MethodDelete:  true,
		http.MethodOptions: true,
	}
	GetMethod  = map[string]bool{http.MethodGet: true}
	PostMethod = map[string]bool{http.MethodPost: true}
	HeadMethod = map[string]bool{http.MethodHead: true}
)
