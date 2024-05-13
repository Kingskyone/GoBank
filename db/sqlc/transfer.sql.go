// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfer (
    from_account_id,
    to_account_id,
    money
   ) VALUES (
      $1, $2, $3
     ) RETURNING id, from_account_id, to_account_id, money, created_at
`

type CreateTransferParams struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Money         float64 `json:"money"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Money)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfer WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, money, created_at FROM transfer WHERE id= $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRow(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}

const listTransferWithFromAccountID = `-- name: ListTransferWithFromAccountID :many
SELECT id, from_account_id, to_account_id, money, created_at FROM transfer WHERE from_account_id=$1 ORDER BY id LIMIT $2 OFFSET $3
`

type ListTransferWithFromAccountIDParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransferWithFromAccountID(ctx context.Context, arg ListTransferWithFromAccountIDParams) ([]Transfer, error) {
	rows, err := q.db.Query(ctx, listTransferWithFromAccountID, arg.FromAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Money,
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

const listTransferWithID = `-- name: ListTransferWithID :many
SELECT id, from_account_id, to_account_id, money, created_at FROM transfer ORDER BY id LIMIT $1 OFFSET $2
`

type ListTransferWithIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransferWithID(ctx context.Context, arg ListTransferWithIDParams) ([]Transfer, error) {
	rows, err := q.db.Query(ctx, listTransferWithID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Money,
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

const listTransferWithToAccountID = `-- name: ListTransferWithToAccountID :many
SELECT id, from_account_id, to_account_id, money, created_at FROM transfer WHERE to_account_id=$1 ORDER BY id LIMIT $2 OFFSET $3
`

type ListTransferWithToAccountIDParams struct {
	ToAccountID int64 `json:"to_account_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListTransferWithToAccountID(ctx context.Context, arg ListTransferWithToAccountIDParams) ([]Transfer, error) {
	rows, err := q.db.Query(ctx, listTransferWithToAccountID, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Money,
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

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfer SET money = $2 WHERE id = $1 RETURNING id, from_account_id, to_account_id, money, created_at
`

type UpdateTransferParams struct {
	ID    int64   `json:"id"`
	Money float64 `json:"money"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, updateTransfer, arg.ID, arg.Money)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}
