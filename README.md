# A Simple Go API

This app is just mean to be an API collection of different interfaces.

- [ ] Will have a basic RESTful API structure
- [x] Inclues JWT Authentication
- [ ] Searching
- [ ] Will have caching

## What can this API do

Essentially, There will be a catalog of products that anyone can view. They'll also be able to search the products catalog. A client will then be able to create a user profile and from there will then be able to create/edit/delete a product.

## Routes

- `GET /`: index
- `GET /healthcheck`: healthcheck
- `GET /api/products`: list of all products
- `CREATE /api/products`: Creates a new product
- `GET /api/products/:product_id`: Gets a detailed description of a product
- `PUT /api/products/:product_id`: Updates the product
- `DELETE /api/products/:product_id`: Deletes the product
- `CREATE /api/users`: creates a new user
- `GET /api/users/:user_id`: View user information
- `PUT /api/users/:user_id`: Updates the user
- `CREATE /api/sessions`: Logs in the user

## TODOs

This is a very basic and messy api system. I more wanted to showcase the overall theme. There are some nice-to-haves that I want to build

- [ ] Need to figure out some sort of middleware for the api. Quite a bit of things are repeated.
- [ ] Need to better sanitize these queries. Right now they can be infected with SQL injection.
- [ ] Want to introduce a centralized config file. This is where the env vars can be initialized and missing, it will not start. This will prevent some false-positives.
