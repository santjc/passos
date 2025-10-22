-- name: GetProductByID :one
SELECT * FROM product WHERE id = $1;