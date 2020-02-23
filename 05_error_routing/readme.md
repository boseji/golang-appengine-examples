# Routing the Invalid Request

There are many times the incoming HTTP request report
something or resource not available.
By default the `net/http` Router would show a `404` message.
However, its not good - when in an App there is only one page.
This is not the desired behavior.
We might need to restrict
users from accessing the incorrect URI.
Perhaps some apps need a custom `404` page
to be shown akin to the site's own style.

We look into this first using a error based
routing.

## tl;dr

Test the program by:

```shell
go run *.go
```

You should see the familiar message on the screen.

`Working Golang App`

The Server is accessible at `http://localhost:8080`

Well now try to visit any odd hyperlink such as :

- `http://localhost:8080/test`
- `http://localhost:8080/123`
- `http://localhost:8080/anything`

All the Above would generate a message :

> **Looks Like you have Stumbled upon something not here**

Thats our [`404` Page](404.html) being served.

The same works well on Appengine as well :

```shell
gcloud app deploy ./app.yaml --version 1
```

## Description

All we are doing here is when ever the `indexHandler` detects
any invalid URI's it serves the [`404` Page](404.html).

```go
    if r.RequestURI != "/" {
        // Display the 404 file
        content, _ := ioutil.ReadFile("404.html")
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "%s", content)
        return
    }
```

One might ask why not use `http.ServeFile` well we need to
also set the **Header** status to `404`.
Hence the `ioutil.ReadFile`.

Plus our HTML file isn't that big so a direct memory load would
not dent much of a performance. However repeated fetches might
hinder the operation.
