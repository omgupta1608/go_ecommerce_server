-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateProduct :one
INSERT INTO products (
  name, price, in_stock
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateOrderProduct :one
INSERT INTO order_products (
  order_id, product_id, quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateProductInStockUnits :exec
UPDATE products SET in_stock = $1 RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (
  user_id, status
) VALUES (
  $1, $2
)
RETURNING *;