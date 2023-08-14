-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetFailedOrderProducts :many
SELECT * from order_products
WHERE placed = false;

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
  order_id, product_id, quantity, placed
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: CreateProductRating :one
INSERT INTO ratings (
  user_id, product_id, rating
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateProductInStockUnits :one
UPDATE products SET in_stock = $1 WHERE id = $2 RETURNING *;

-- name: UpdateOrderStatus :one
UPDATE orders SET status = $1 WHERE id = $2 RETURNING *;

-- name: UpdateOrderProductStatus :one
UPDATE order_products SET placed = $1 WHERE id = $2 RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (
  user_id, status
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListProductsWithRatings :many
select p.id, p.name, CAST(AVG(r.rating) AS DECIMAL(10,2)) as avg_rating, count(*) as rating_count from products as p inner join ratings as r on r.product_id = p.id group by p.id;

-- name: GetOrderDetails :many
select o.id as order_id, o.status, p.id as product_id, p.name, op.quantity, p.price, op.placed, u.name as user_name, u.email from orders as o 
join (order_products as op join products as p on op.product_id = p.id)
on op.order_id = o.id
join users as u on o.user_id = u.id
where o.id = $1;

-- name: GetTop3Customers :many
select u.id, u.name as user_name, count(*) as orders_placed from orders as o 
join users as u on o.user_id = u.id 
group by u.id 
order by orders_placed desc
limit 3;