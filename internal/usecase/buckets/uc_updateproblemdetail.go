package buckets

import "context"

type UpdateProblemDetailRequest struct {
	Description string
	Type        string
	Title       string
	Status      int8
	Detail      string
	Instance    string
}

type UpdateProblemDetailResult struct {
}

func (uc *bucketsUsecase) UpdateProblemDetail(
	ctx context.Context,
	ID int32,
	request UpdateProblemDetailRequest) error {
	sql := `
		update problem_details 
		set 
		    description = $2, 
		    type = $3, 
		    title = $4, 
		    status = $5, 
		    detail = $6, 
		    instance = $7
		where id = $1`

	_, err := uc.storage.Pool.Exec(
		ctx,
		sql,
		ID,
		request.Description,
		request.Type,
		request.Title,
		request.Status,
		request.Detail,
		request.Instance)

	if err != nil {
		return err
	}

	return nil
}
