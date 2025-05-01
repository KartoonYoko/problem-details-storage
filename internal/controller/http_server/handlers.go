package http_server

import (
	"encoding/json"
	"github.com/KartoonYoko/problem-details-storage/internal/usecase/buckets"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (c *Controller) handlePOSTBucket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.CreateBucketRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	result, err := c.usecase.CreateBucket(ctx, request)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "can not serialize response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func (c *Controller) handleGETBuckets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.GetBucketsListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	result, err := c.usecase.GetBucketsList(ctx, request)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "can not serialize response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) handleGETBucketProblemDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.GetBucketProblemDetailsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	result, err := c.usecase.GetBucketProblemDetails(ctx, request)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "can not serialize response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) handleGETBucketProblemDetailByType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.GetBucketProblemDetailByTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	result, err := c.usecase.GetBucketProblemDetailByType(ctx, request)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "can not serialize response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) handlePOSTBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.CreateProblemDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	result, err := c.usecase.CreateProblemDetail(ctx, request)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "can not serialize response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (c *Controller) handlePUTBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.UpdateProblemDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = c.usecase.UpdateProblemDetail(ctx, int32(id), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) handleDELETEBucketProblemDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request buckets.DeleteProblemDetailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "can not parse body", http.StatusBadRequest)
		return
	}

	err := c.usecase.DeleteProblemDetail(ctx, request)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
