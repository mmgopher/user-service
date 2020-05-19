# HTTP server for managing a list of users

## Descrption

HTTP server using Gin-Gonic framework and implementing CRUD operations.
It operates on User entity. The whole project is dockerized

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
- `GET /v1/users/:user_id` - return User entity
- `DELETE /v1/users/:user_id` - delete User entity
- `PUT /v1/users/:user_id` - update User entity
- `POST /v1/users` - create User entity
- `GET /v1/users` - search Users using sorting, filtering and pagination

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