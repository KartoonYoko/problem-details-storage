package buckets

import "context"

type GetBucketProblemDetailByTypeRequest struct {
	Type string
}

type GetBucketProblemDetailByTypeResult struct {
	ID          int32
	Description string
	Type        string
	Title       string
	Status      int8
	Detail      string
	Instance    string
}

func (uc *bucketsUsecase) GetBucketProblemDetailByType(
	ctx context.Context,
	request GetBucketProblemDetailByTypeRequest) (
	*GetBucketProblemDetailByTypeResult,
	error) {
	sql := `
		select 
		    pd.id, 
		    pd.description, 
		    pd.type, 
		    pd.title, 
		    pd.status, 
		    pd.detail, 
		    pd.instance
		from problem_details pd
		where pd.type = $1`

	row := uc.storage.Pool.QueryRow(ctx, sql, request.Type)

	var result GetBucketProblemDetailByTypeResult
	err := row.Scan(
		&result.ID,
		&result.Description,
		&result.Type,
		&result.Title,
		&result.Status,
		&result.Detail,
		&result.Instance)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
