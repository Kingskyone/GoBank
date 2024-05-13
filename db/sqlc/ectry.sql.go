// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ectry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO ectries (
    account_id,
    amount
   ) VALUES (
      $1, $2
     ) RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64   `json:"account_id"`
	Amount    float64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Ectry, error) {
	row := q.db.QueryRow(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Ectry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM ectries WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteEntry, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM ectries WHERE id= $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Ectry, error) {
	row := q.db.QueryRow(ctx, getEntry, id)
	var i Ectry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntryWithAccountID = `-- name: ListEntryWithAccountID :many
SELECT id, account_id, amount, created_at FROM ectries WHERE account_id=$1 ORDER BY id LIMIT $2 OFFSET $3
`

type ListEntryWithAccountIDParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntryWithAccountID(ctx context.Context, arg ListEntryWithAccountIDParams) ([]Ectry, error) {
	rows, err := q.db.Query(ctx, listEntryWithAccountID, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ectry{}
	for rows.Next() {
		var i Ectry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEntryWithID = `-- name: ListEntryWithID :many
SELECT id, account_id, amount, created_at FROM ectries ORDER BY id LIMIT $1 OFFSET $2
`

type ListEntryWithIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntryWithID(ctx context.Context, arg ListEntryWithIDParams) ([]Ectry, error) {
	rows, err := q.db.Query(ctx, listEntryWithID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ectry{}
	for rows.Next() {
		var i Ectry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntry = `-- name: UpdateEntry :one
UPDATE ectries SET amount = $2 WHERE id = $1 RETURNING id, account_id, amount, created_at
`

type UpdateEntryParams struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Ectry, error) {
	row := q.db.QueryRow(ctx, updateEntry, arg.ID, arg.Amount)
	var i Ectry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
