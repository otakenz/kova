```sh
Testing

Write tests for each layer independently:

core/ domain logic pure unit tests.

app/ service tests with mocked ports interfaces.

infra/ integration tests for real DB or logger.

api/v1/ handler tests with HTTP test server and mocked services.


High Priority
1. internal/core
2. internal/app

Medium Priority
1. internal/ports
2. internal/infra

Low Priority
1. api/v1 and middleware

```
