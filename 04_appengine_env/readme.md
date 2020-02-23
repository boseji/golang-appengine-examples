# Appengine Environment

In this section we would be looking closely into how the
[Appengine Runtime Environment](https://cloud.google.com/appengine/docs/standard/go112/runtime) works.

Remember that we are using the **Go 1.12 Appengine** environment.

That is specified in our [`app.yaml`](app.yaml) file:

```yaml
runtime: go112
```

Here we would look into specifically how to access the
**[default environment variables][1]** under **Go 1.12 Appengine**.

## tl;dr

To test the App :

```shell
go run *.go
```

You would possibly see a bunch of parameters that are empty.
This is because we are running in the local environment not
in Appengine.

So, we need to deploy to appengine to actually see all the parameters.

```shell
gcloud app deploy ./app.yaml --version 1
```

Note the `*.go` in the *Test command* since there are *2 Golang files*.

## Description

By default [**Go 1.12 Appengine**][2] runtime exposes some useful variables.
These can become really handy in many applications and our
coming examples.

For more details on these variables refer to:

<https://cloud.google.com/appengine/docs/standard/go112/runtime#environment_variables>

The program [`gaeinfo.go`](gaeinfo.go) provides a data-type and loading
function to achieve the read of these variables.

The same is printed on the webpage in the default endpoint.

Also note the change in our `indexHandler` function in the beginning:

```go
  if r.RequestURI != "/" {
    http.Error(w, "Incorrect request", http.StatusNotFound)
    return
  }
```

This helps to avoid the un-necessary all request responder.
We would later visit this for creating the **404** redirect.
At the moment it only prints a message for any
path other than `"/"` that the user tries to access.
Also it returns the status for this error as *404* instead
of the usual *200*.

Next we look at ***custom environment variables***.

These are enabled by the special `env_variable` definition inside
[`app.yaml`](app.yaml):

```yaml
runtime: go112

env_variables:
  MY_VAR: "My Custom Environment Variable"
```

We read this variable also and can provide action based on that.

```go
  myVARenv = os.Getenv("MY_VAR")
  if myVARenv == "" {
    log.Println("Environment Variable 'MY_VAR' not set")
  }
```

While running in local environment without setting the `MY_VAR` -
It would show a message as above.
To test out the value display you can do the following steps:

```shell
export MY_VAR="This is Custom Value"
go run *.go
```

Of course in the Appengine it should work as expected and should
show the value mentioned in the `app.yaml`.

 [1]:https://cloud.google.com/appengine/docs/standard/go112/runtime#environment_variables
 [2]:https://cloud.google.com/appengine/docs/standard/go112/runtime