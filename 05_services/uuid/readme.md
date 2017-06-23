# App Engine Microservices Example 1

This part is the Microservice called `uuid`

This generates a new UUID whenever the route is visited.

> Note: The `app.yaml` specifically mentions the service name `service: uuid`

> Note: Even though this is a microservice the `main.go` file still connects to the `package main`. Its only by redirection that it has a separate route.