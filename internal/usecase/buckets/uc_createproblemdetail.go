package buckets

import "context"

type CreateProblemDetailRequest struct {
	Description string
	Type        string
	Title       string
	Status      int8
	Detail      string
	Instance    string
}

type CreateProblemDetailResult struct {
	ID int32
}

func (uc *bucketsUsecase) CreateProblemDetail(
	ctx context.Context,
	request CreateProblemDetailRequest) (
	*CreateProblemDetailResult,
	error) {
	sql := `
		insert into problem_details (description, type, title, status, detail, instance)
		values ($1, $2, $3, $4, $5, $6)
		returning id`

	row := uc.storage.Pool.QueryRow(
		ctx,
		sql,
		request.Description,
		request.Type,
		request.Title,
		request.Status,
		request.Detail,
		request.Instance)

	var result CreateProblemDetailResult
	err := row.Scan(&result.ID)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
