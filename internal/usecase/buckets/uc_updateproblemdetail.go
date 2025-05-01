package buckets

import "context"

type UpdateProblemDetailRequest struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Status      int8   `json:"status"`
	Detail      string `json:"detail"`
	Instance    string `json:"instance"`
}

func (uc *Usecase) UpdateProblemDetail(
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
