package buckets

import (
	"context"
	"fmt"
)

type GetBucketsListRequest struct {
	Limit  int32
	Offset int32
}

type GetBucketsListResult struct {
	items []GetBucketsListResultItem
}

type GetBucketsListResultItem struct {
	ID          int32
	Name        string
	Description string
}

func (uc *bucketsUsecase) GetBucketsList(ctx context.Context, request GetBucketsListRequest) (*GetBucketsListResult, error) {
	sql := `
		select id, name, description
		from buckets
		order by name desc
		offset $1 
		limit $2;`

	rows, err := uc.storage.Pool.Query(ctx, sql, request.Offset, request.Limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []GetBucketsListResultItem
	for rows.Next() {
		var item GetBucketsListResultItem
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return &GetBucketsListResult{items: items}, nil
}
