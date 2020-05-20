# HTTP server for managing a list of users

## Descrption

HTTP server using Gin-Gonic framework and implementing CRUD operations.
It operates on User entity. The whole project is dockerized.
It also contains implemetation of seek pagination with sorting and filtering

### User entity
Attributes:
- name
- surname
- gender
- age
- address
- created_at

> name and surname are the user's unique identifier. There can be only one user with given name and surname

### Endpoints
- `GET /v1/users/:user_id` - return User entity in JSON format
- `DELETE /v1/users/:user_id` - delete User entity
- `PUT /v1/users/:user_id` - update User entity. Request in JSON format
- `POST /v1/users` - create User entity. Request in JSON format
- `GET /v1/users` - search Users using sorting, filtering and seek pagination

Exmples"
- `GET /v1/users` - return up to 30 users sort by `id` ascending
- `GET /v1/users?limit=100&sort=name:desc` - return up to 100 users sort by `name` descending
- `GET /v1/users?gender=male&limit=100&sort=age:asc` - return up to 100 users sort by `age` ascending with gender `male`
- `GET /v1/users?name=sonny&gender=male&limit=100&sort=age:desc` - return up to 100 users sort by `age` descending with gender `male` and name like `%sonny%` case insensitive
- `GET /v1/users?name=sonny&gender=male&limit=100&sort=age:desc&min_age=30` - return up to 100 users sort by `age` descending with gender `male` and name like `%sonny%` case insensitive with `age` >= 30
- `GET /v1/users?name=sonny&gender=male&limit=100&sort=age:asc&max_age=30` - return up to 100 users sort by `age` ascending with gender `male` and name like `%sonny%` case insensitive with `age` <= 30
- `GET /v1/users?name=sonny&gender=male&limit=100&sort=age:asc&min_age=30&max_age=45` - return up to 100 users sort by `age` ascending with gender `male` and name like `%sonny%` case insensitive with `age` >= 30 and `age` <= 45
- `GET /v1/users?limit=100&sort=created_at:desc` - return up to 100 users sort by `created_at` descending

> There is also Pagination object in response to know how to query next or previous page

## Testing

### Unit tests

To start unit tests please call command:

`make go_test_unit`

The unit tests do not have any dependencies

### Integration tests

To start integration tests please call command:

`make docker_build_image && make application_test`

The command does:
- build application docker image
- start postgres DB in docker container
- start application itself in docker container
- execute goose migration to preapre DB schema
- insert test data into postgres
- execute integration tests in docker container

## Run Service Locally

To run service locally please call command:

`application_run`

Start you browser and enter url:

`http://localhost:8080/v1/users`

## Errors reporting

In case of errors there will be returned custom error object in JSON format with custom error code and message.

For example in case of validation error there is retruend HTTP code 400 with detailed error code and message
Examples:
- {"code":2140004,"message":"`limit` can not be negative"}
- {"code":2040001,"message":"`name` can't be empty"}
- {"code":2040001,"message":"`afterID` can not be negative"}

## Project structure

- **app**  - aplication code
    - **api** - definition of response and request objects
    - **config** - application configuration object
    - **controller** - controller layer
    - **dao** - repository layer
    - **db** - db helpers
    - **httperrors** - definition of custom http error object and predefined application errors
    - **middleware** - gin-gonic middleware
    - **model** - database models
    - **service** -service layer 
- **build** - docker and docker-compose files to build, run and test application
- **test** - integration tests    