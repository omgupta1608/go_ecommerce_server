# Go Ecommerce REST API server

Aftershoot Task - Om Gupta

## Running locally
- Clone the repository
- Run `docker compose up -d` - this will spin up 2 containers
    - One for the API server
    - One for the database
- Go to the server directory - `cd server/`
- Run `make db-migrate-up` to seed the database
- The API server is ready to use on `http://localhost:5000`

## Testing
- APIs are avaliable at `http://localhost:5000/api/v1/`
- List of all APIs and their usage is present in this postman collection as well as written below in the 'Available APIs section'

## Design Description
- The application contains 2 types of users or 2 tenants - CUSTOMER and ADMIN
- The server differentiates users with the help of the web token they provide
- Some APIs are available only to ADMINs or CUSTOMERs and some are available for both
## Features Implemented
- Customer/Admin Register
- Customer/Admin Login
- Add New Product - ADMIN
- Rate Product - USER
- Place Order - CUSTOMER
- Process Order - ADMIN
- Get Order Details - ADMIN and CUSTOMER
- Get Product Information - ADMIN and CUSTOMER
- Get Top 3 Customers - ADMIN and CUSTOMER


## APIs
- `tenant_type` can be one of the following types
    - `CUSTOMER`
    - `ADMIN`

<details>

<summary style="font-size:20px">Register API</summary>

`POST /auth/register`
#### Request Body

``` json
{
    "email":"admin@gmail.com",
    "name":"Admin User",
    "password":"admin123",
    "tenant_type":"ADMIN" | "CUSTOMER"
}
```

#### Response
``` json
{
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkFETUlOIiwidXNlcl9pZCI6IjZmMmZjZGI1LTdkZTUtNDFhYS04NTE1LWNlMjRlMWM2MWFjMiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20ifQ.tbYAXLeY9Wg5JgcaM422fuHfO_9rduiADSGmzokCwqk",
        "email": "admin@gmail.com",
        "user_id": "6f2fcdb5-7de5-41aa-8515-ce24e1c61ac2"
    },
    "message": "Registered"
}
```
</details>

<details>

<summary style="font-size:20px">Login API</summary>

`POST /auth/login`

#### Request Body

``` json
{
    "email": "admin@gmail.com",
    "password":"admin123"
}
```

#### Response
``` json
{
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkFETUlOIiwidXNlcl9pZCI6IjZmMmZjZGI1LTdkZTUtNDFhYS04NTE1LWNlMjRlMWM2MWFjMiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20ifQ.tbYAXLeY9Wg5JgcaM422fuHfO_9rduiADSGmzokCwqk",
        "email": "admin@gmail.com",
        "user_id": "6f2fcdb5-7de5-41aa-8515-ce24e1c61ac2"
    },
    "message": "Logged In"
}
```
</details>

<details>

<summary style="font-size:20px">Add Product API</summary>

`POST /product/new`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkFETUlOIiwidXNlcl9pZCI6ImUwNWNkNmU0LTg0MmYtNGU4Mi05YzI4LTc2YzJmZWNjMzc0NiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20ifQ.Yh4RhypBo-1sFBma57m5yYikPeCYVIiOtRRjIf6MRQ0"
}
```

#### Request Body

``` json
{
    "name": "One Plus 10 R",
    "price": 35000,
    "in_stock": 100
}
```

#### Response
``` json
{
    "data": {
        "is_product_in_stock": 100,
        "product_id": "ad45e09d-c805-4ca5-ac7f-f73aceea7e61",
        "product_name": "One Plus 10 R",
        "product_price": 35000
    },
    "message": "Product Added!"
}
```
</details>

<details>

<summary style="font-size:20px">Rate Product API</summary>

`POST /product/rate`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkNVU1RPTUVSIiwidXNlcl9pZCI6IjIzNTVjZjk0LTMyNWEtNDZhMy1iZWYzLWVjM2FiYThlMTc5NSIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhQGdtYWlsLmNvbSJ9.Q4kZlTFM0968njz_s1CONwmvEw4yY6bvXc2rBzT9kS4"
}
```

#### Request Body

``` json
{
    "product_id": "23eb7f70-e9ba-48f7-9fe5-d9df27c097f4",
    "rating": 4 (1 to 5)
}
```

#### Response
``` json
{
    "data": {},
    "message": "Ratings saved"
}
```
</details>

<details>

<summary style="font-size:20px">List Product API</summary>

`GET /product`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZjAwMDkyOWItZmFhYS00NDE2LWFlMjQtZTZlMzhkOGI4ODMzIiwiaXNfYWN0aXZlIjp0cnVlLCJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSJ9.fsx3zxCoD4sMYBQzc7QX6O1kCFyC0JNtZd5ZxANCMHs"
}
```

#### Response
``` json
{
    "data": {
        "products": [
            {
                "ID": "23eb7f70-e9ba-48f7-9fe5-d9df27c097f4",
                "Name": "One Plus 10 R",
                "AvgRating": "4.00",
                "RatingCount": 1
            }
        ]
    },
    "message": "All our products"
}
```
</details>

<details>

<summary style="font-size:20px">Place Order API</summary>

`POST /order/new`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkNVU1RPTUVSIiwidXNlcl9pZCI6IjIzNTVjZjk0LTMyNWEtNDZhMy1iZWYzLWVjM2FiYThlMTc5NSIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhQGdtYWlsLmNvbSJ9.Q4kZlTFM0968njz_s1CONwmvEw4yY6bvXc2rBzT9kS4"
}
```
#### Request Body

``` json
{
    "products": [
        {
            "product_id": "23eb7f70-e9ba-48f7-9fe5-d9df27c097f4",
            "quantity": 3
        }
    ]
}
```
#### Response
``` json
{
    "data": {
        "order_total": 105000,
        "products": [
            {
                "product_id": "23eb7f70-e9ba-48f7-9fe5-d9df27c097f4",
                "placed": true
            }
        ]
    },
    "message": "Order Placed. Thanks"
}
```
</details>


<details>

<summary style="font-size:20px">Order Details API</summary>

`GET /order/:id`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkNVU1RPTUVSIiwidXNlcl9pZCI6IjIzNTVjZjk0LTMyNWEtNDZhMy1iZWYzLWVjM2FiYThlMTc5NSIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhQGdtYWlsLmNvbSJ9.Q4kZlTFM0968njz_s1CONwmvEw4yY6bvXc2rBzT9kS4"
}
```

#### Response
``` json
{
    "data": {
        "order_total": 105000,
        "products": [
            {
                "id": "23eb7f70-e9ba-48f7-9fe5-d9df27c097f4",
                "name": "One Plus 10 R",
                "quantity": 3,
                "price": 35000,
                "placed": true
            }
        ],
        "user_email": "adam@gmail.com",
        "user_name": "Adam"
    },
    "message": "Order Details"
}
```
</details>


<details>

<summary style="font-size:20px">TOP 3 Customers API</summary>

`GET /user/top-3`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkNVU1RPTUVSIiwidXNlcl9pZCI6IjIzNTVjZjk0LTMyNWEtNDZhMy1iZWYzLWVjM2FiYThlMTc5NSIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhQGdtYWlsLmNvbSJ9.Q4kZlTFM0968njz_s1CONwmvEw4yY6bvXc2rBzT9kS4"
}
```

#### Response
``` json
{
    "data": {
        "customers": [
            {
                "ID": "2355cf94-325a-46a3-bef3-ec3aba8e1795",
                "UserName": "John Doe",
                "OrdersPlaced": 11
            },
            {
                "ID": "3c1fb8df-106d-4513-878b-fc9df07a2295",
                "UserName": "Alice",
                "OrdersPlaced": 7
            },
            {
                "ID": "7715cf94-325a-46a3-bef3-ec3aba85gh17",
                "UserName": "Bob",
                "OrdersPlaced": 3
            }
        ]
    },
    "message": "Our top 3 customers"
}
```
</details>



<details>

<summary style="font-size:20px">Process Order API</summary>

`POST /order/process`
#### Request Headers
``` json
{
    "Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfdHlwZSI6IkFETUlOIiwidXNlcl9pZCI6IjZmMmZjZGI1LTdkZTUtNDFhYS04NTE1LWNlMjRlMWM2MWFjMiIsImlzX2FjdGl2ZSI6dHJ1ZSwiZW1haWwiOiJhZG1pbkBnbWFpbC5jb20ifQ.tbYAXLeY9Wg5JgcaM422fuHfO_9rduiADSGmzokCwqk"
}
```
#### Request Body

``` json
{
    "order_id":"3c1fb8df-106d-4513-878b-fc9df07a2295",
    "status": "COMPLETED" | "INITIATED" | "CANCELLED" | "FAILED"
}
```

#### Response
``` json
{
    "data": {},
    "message": "Order Updated"
}
```
</details>

