package buckets

import (
	"context"
	"fmt"
)

type GetBucketProblemDetailsRequest struct {
	BucketID int32 `json:"bucketId"`
	Offset   int32 `json:"offset"`
	Limit    int32 `json:"limit"`
}

type GetBucketProblemDetailsResult struct {
	items []GetBucketProblemDetailsResultItem `json:"items"`
}

type GetBucketProblemDetailsResultItem struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Status      int8   `json:"status"`
	Detail      string `json:"detail"`
	Instance    string `json:"instance"`
}

func (uc *Usecase) GetBucketProblemDetails(
	ctx context.Context,
	request GetBucketProblemDetailsRequest) (
	response *GetBucketProblemDetailsResult,
	err error) {
	sql := `
		select 
		    pd.id, 
		    pd.description, 
		    pd.type, 
		    pd.title, 
		    pd.status, 
		    pd.detail, 
		    pd.instance
		from bucket_problem_details bpd
		left join problem_details pd on pd.id = bpd.problem_detail_id
		where bpd.id = $1
		order by pd.id
		offset $2 
		limit $3`

	rows, err := uc.storage.Pool.Query(ctx, sql, request.BucketID, request.Offset, request.Limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []GetBucketProblemDetailsResultItem
	for rows.Next() {
		var item GetBucketProblemDetailsResultItem
		err := rows.Scan(
			&item.ID,
			&item.Description,
			&item.Type,
			&item.Title,
			&item.Status,
			&item.Detail,
			&item.Instance,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return &GetBucketProblemDetailsResult{items: items}, nil
}
