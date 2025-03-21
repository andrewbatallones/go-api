# A Simple Go API

This app is just mean to be an API collection of different interfaces.

- [ ] Will have a basic RESTful API structure
- [ ] Inclues JWT Authentication
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