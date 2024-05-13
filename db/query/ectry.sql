-- name: CreateEntry :one
INSERT INTO ectries (
    account_id,
    amount
   ) VALUES (
      $1, $2
     ) RETURNING *;


-- name: GetEntry :one
SELECT * FROM ectries WHERE id= $1 LIMIT 1;

-- name: ListEntryWithID :many
SELECT * FROM ectries ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListEntryWithAccountID :many
SELECT * FROM ectries WHERE account_id=$1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateEntry :one
UPDATE ectries SET amount = $2 WHERE id = $1 RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM ectries WHERE id = $1;