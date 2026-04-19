-- name: GetProducts :many
SELECT id, name, description, price, stock, image_url
FROM product
WHERE is_deleted = false;
