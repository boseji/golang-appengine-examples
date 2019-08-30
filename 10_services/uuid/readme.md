# App Engine Micro-services Example 1

This part is the Micro-service called `uuid`

This generates a new UUID whenever the route is visited.

> Note: The `app.yaml` specifically mentions the service name `service: uuid`

Also:

> Note: Even though this is a micro-service the `main.go` file still connects
> to the `package main`. Its only by redirection that it has a separate route.
