# A Simple Go API

This app is just mean to be an API collection of different interfaces.

- [ ] Will have a basic RESTful API structure
- [x] Inclues JWT Authentication
- [ ] Searching
- [x] Caching

## What can this API do

Essentially, There will be a catalog of products that anyone can view. They'll also be able to search the products catalog. A client will then be able to create a user profile and from there will then be able to create/edit/delete a product.

## Setup

### Redis

You can run the command to spin up a new docker instance of Reids. This will add a default Redis instance with no password (should be setup through the env var).

```bash
docker run --name go-api-redis -p 6379:6379 -d redis
```

### Elasticsearch

You can create a new elasticsearch instance on Docker with this command

```bash
docker run --name go-api-es01 --net elastic -p 9200:9200 -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:latest
```

You can run this command if you want to add the Kibana dashboard. Essentially, Kibana is used to monitor your Elasticsearch. You can view more here: [Elasticsearch Kibana](https://www.elastic.co/kibana).

```bash
docker run --name go-api-kib01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:latest
```

## Tests

You can run the tests via the go command

```bash
go test -v ./...
```

## Routes

- `GET /` index
- `GET /healthcheck` healthcheck
- `GET /api/products` list of all products
- `GET /api/products/search` search products
- `CREATE /api/products` Creates a new product
- `GET /api/products/:product_id` Gets a detailed description of a product
- `PUT /api/products/:product_id` Updates the product
- `DELETE /api/products/:product_id` Deletes the product
- `CREATE /api/users` creates a new user
- `GET /api/users/:user_id` View user information
- `PUT /api/users/:user_id` Updates the user
- `CREATE /api/sessions` Logs in the user

## TODOs

This is a very basic and messy api system. I more wanted to showcase the overall theme. There are some nice-to-haves that I want to build

- [x] Need to figure out some sort of middleware for the api. Quite a bit of things are repeated.
    - It's the the best, I need to figure out if there's a better way to implement a middleware function.
- [ ] Need to better sanitize these queries. Right now they can be infected with SQL injection.
- [ ] Want to introduce a centralized config file. This is where the env vars can be initialized and missing, it will not start. This will prevent some false-positives.
- [x] GitHub Actions tests are currently failing because there is currently no setup for the database to utilize the app. Add config settings to build/use/connect psql.
