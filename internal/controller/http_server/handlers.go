package http_server

import (
	"net/http"
)

func handlePOSTBucket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleGETBuckets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleGETBucketProblemDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleGETBucketProblemDetailByType(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handlePOSTBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handlePUTBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleDELETEBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
