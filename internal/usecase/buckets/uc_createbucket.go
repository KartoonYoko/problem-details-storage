package buckets

import "context"

type CreateBucketRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateBucketResult struct {
	ID int32 `json:"id"`
}

func (uc *Usecase) CreateBucket(
	ctx context.Context,
	request CreateBucketRequest) (result *CreateBucketResult, err error) {
	sql := `
		insert into buckets (name, description) 
		values ($1, $2)
		returning id`

	var id int32
	err = uc.storage.Pool.QueryRow(ctx, sql, request.Name, request.Description).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &CreateBucketResult{ID: id}, nil
}
