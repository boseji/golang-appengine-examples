# Hello World Example

This is basic example demonstrating a simple webserver
hosted on port `8080` localhost.

## Upgraded for Go 1.12 Compatibility

Use the following Commands for Module :

```shell
go mod init github.com/boseji/golang-appengine-examples/00_basic
go mod tidy
```

This was needed to be done only once to create the specific module dependency file.

## tl;dr

To run this code go to this directory

```shell
go run main.go
```

And in the browser access
`http://localhost:8080`

It should show the text:

`Working Golang App`

## Detail

### init()

In this example we use the special `init()` function

```go
func init() {
  http.HandleFunc("/", indexHandler)
  http.Handle("/favicon.ico", http.NotFoundHandler())
}
```

This function execute before the `main()` in every type
of `go` file or package.

Inside the `init()` We are adding web two end points :

- First `/` is the root end point added to the `DefaultServeMux`
  in the `net\http` package using the `HandleFunc` method.

- Second `/favicon.ico` which is generally the small icon shown on
  corner of the browser tab. Though this is not explicitly processed
  some times the browsers request it. So we associate the default
  `NotFoundHandler()` to generate a `404` Not found response.

With **Go 1.12 Appengine Upgrade** this special **`init()`** will
**no longer be needed** and hence omitted in the next examples.

#### indexHandler()

Next, we have the root handler
function

`func indexHandler(w http.ResponseWriter, r *http.Request)`

This function is called whenever the client on the browser visits our
website. This is the root end point for the server.

Inside the `indexHandler` we write out the string `Working Golang App`
that would be displayed whenever the root end point is hit.

This is achieved by `fmt.Fprint(w, "Working Golang App")` statement.

Here in this statement we are using the `fmt` package `Fprint` or
File Print function that would accept the `io.writer` interface.

Fortunately the `http.ResponseWriter` does implements the
`io.writer` interface.

#### main()

Finally the `main()` function, where we are using the `ListenAndServe`
method to host the web server.

Previously we have added endpoints to `DefaultServeMux` in
the `init()` function. So the second argument of `ListenAndServe` can
be `nil`. Internally this function would access the `DefaultServeMux`
endpoints whenever we pass `nil`

Next we can notice the use `log.Fatal`, this helps to log any error
that might occur while starting or during operation of the http server.

Since the `ListenAndServe` is a blocking function, it continues to run
infinitely serving the pages. Till we send the `SIGTERM` or `SIGKILL`
to the process or simply press `Ctrl+C`
