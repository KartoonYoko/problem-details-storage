package buckets

import "context"

type DeleteProblemDetailRequest struct {
	ID int32 `json:"id"`
}

func (uc *Usecase) DeleteProblemDetail(
	ctx context.Context,
	request DeleteProblemDetailRequest) error {
	sql := `delete from problem_details where id = $1`
	_, err := uc.storage.Pool.Exec(ctx, sql, request.ID)

	if err != nil {
		return err
	}

	return nil
}
