-- name: GetProducts :many
SELECT id, name, description, price, stock, image_url
FROM product
WHERE is_deleted = false;

-- name: GetProductByID :one
SELECT id, name, description, price, stock, image_url
FROM product
WHERE id = $1 AND is_deleted = false;

-- name: CreateProduct :one
INSERT INTO product (name, description, price, stock, image_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, price, stock, image_url;
