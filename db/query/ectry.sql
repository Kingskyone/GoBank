-- name: CreateEctry :one
INSERT INTO ectries (
    account_id,
    amount
   ) VALUES (
      $1, $2
     ) RETURNING *;


-- name: GetEctry :one
SELECT * FROM ectries WHERE id= $1 LIMIT 1;

-- name: ListEctryWithID :many
SELECT * FROM ectries ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListEctryWithAccountID :many
SELECT * FROM ectries WHERE account_id=$1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateEctry :one
UPDATE ectries SET amount = $2 WHERE id = $1 RETURNING *;

-- name: DeleteEctry :exec
DELETE FROM ectries WHERE id = $1;