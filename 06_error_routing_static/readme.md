# Static File Serving with Special 404 Page for Appengine

In the [previous example](../05_error_routing) we looked at
how for a single page we can show custom `404` page.

Now we would try to Host a website with static files,
and a same single page in one. And use the same `404` page
for both.

## TLDR;

To execute just give the command :

```shell
go run *.go
```

Upon visiting the default `http://localhost:8080` location
you would be greeted by :

`Working Golang App`

Well that we already know. Let's now try to visit
the static homepage `http://localhost:8080/home`.

> **This is a Home Page**

Ya a simple HTML page hosted here.

Here are a few valid URLs :

- <http://localhost:8080>
- <http://localhost:8080/home>
- <http://localhost:8080/home/1.txt>
- <http://localhost:8080/home/2.txt>
- <http://localhost:8080/home/css/index.css>

Now let's try some invalid URLs :

- <http://localhost:8080/123>
- <http://localhost:8080/hello>
- <http://localhost:8080/home/22>
- <http://localhost:8080/home/test>

All the above would show our custom [`404.html`](public/404.html) page.
You would be greeted by a message :

> **Looks Like you have Stumbled upon something not here**

This would also work in Appengine:

```shell
gcloud app deploy ./app.yaml --version 1
```

# Description

What we did here is divided into 3 parts :

 1. Creating a callable custom 404 function for displaying the
    custom [`404.html`](public/404.html) page.
 2. Wrap the normal [`http.FileServer`](https://godoc.org/net/http#FileServer)
    to bypass the default 404 page displayed by It.
 3. Wrap the callable custom 404 function to be called through
    [`http.Handler`](https://godoc.org/net/http#Handler) interface.

## Attribution

The solution is a derived work from the original code by
[Thomas Broyer](https://stackoverflow.com/users/116472/thomas-broyer).

His solution is present at :

<https://stackoverflow.com/questions/39318367/how-to-handle-404-error-on-gae-using-http-fileserver-in-golang>

More specifically:

<https://stackoverflow.com/a/39599490>

## Creating the Callable Custom 404 function

We were trying to read the file and write it back to `http.ResponseWriter`.

Here is the same code with some added points:

```go
func return404(w http.ResponseWriter, _ *http.Request) {
    content, _ := ioutil.ReadFile("public/404.html")
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "%s", content)
}
```

We have used the<br>
`w.Header().Set("Content-Type", "text/html")`<br>
to make sure that we have the correct content type set.

Another **important** thing is this call (`w.Header().Set(`)
needs to be before `w.WriteHeader(http.StatusNotFound)`.
The reason being that `WriteHeader` finalizes the status and
packages into the HTTP response.

## Wrapping `http.FileServer`

This is actually achieved by two parts:

1. Custom `http.ResponseWriter` Implementation
2. Wrapping Function to generate the special Handler

### Customizing the `http.ResponseWriter`

First let's look at the `http.ResponseWriter` and the reasons
we choose to customize it.

Here are the reasons for using custom version of `http.ResponseWriter` :

1. The `http.FileServer` uses the standard `http.ResponseWriter` object
   internally.
2. We don't have access to this internal `http.ResponseWriter`.
3. We need to prevent `http.FileServer` from displaying the default message.

Here is *What we intend to do* :

1. Have a modular approach to the modification such that it can extended later.
2. Make sure that a different handler is invoked when the `404` is encountered.
3. Prevent the default page from being written to the response.

Let's have a quick look at what is [`http.ResponseWriter`](https://godoc.org/net/http#ResponseWriter):

```go
type ResponseWriter interface {
    Header() Header

    Write([]byte) (int, error)

    WriteHeader(statusCode int)
}
```

We have removed all the comments parts for visibility reasons.

You can always look at the actual code:
<https://golang.org/src/net/http/server.go#L94>

So, `http.ResponseWriter` is an Interface that also implements
`io.Writer` interface.

Let's now look at what we have defined:

```go
type intercept404Writer struct {
    req             *http.Request       // Store the Current Request
    writer          http.ResponseWriter // Store the Original Response Writer
    notFoundHandler http.Handler        // Initiate Not Found Response and Ignore writes
    header          http.Header         // Used as the Virtual Header Storage
}
```

Well there are 4 members.

 1. `req` we store the incoming request here
 2. `writer` we store the original `http.ResponseWriter` passed down from
    `http.FileServer`
 3. `notFoundHandler` is the HTTP request handler for the custom `404` page
 4. `header` This would be a virtual store where all headers set in the current
     response are stored. This is later would be used to initialize
     the actual `http.ResponseWriter` Header.

Let us understand the process in which the *Headers* are assembled and
*How block the `http.FileServer`* :

 1. Page to be rendered is selected by the `http.FileSystem` Interface
 implemented in `http.Dir` which intern is a `string` type.
 This is part of the `http.FileServer` implementation.

 2. Headers are assembled using `http.ResponseWriter.Header()` call.
 This is the First time the `http.ResponseWriter` interface
 method `Header() Header` is called.

 3. Page is written using the `http.ResponseWriter` interface method
 `Write([]byte) (int, error)`. We are going to disable this in case
 of the `404` request if so in the *step 4* below. Since we would
 be handing over control to the `intercept404Writer.notFoundHandler`.

 4. Final Status code is updated.
 This is done using `http.ResponseWriter` interface method
 `WriteHeader(statusCode int)`. At this time all the Headers need be
 populated for the particular request.
 For this same reason we are maintaining a *virtual Header storage*
 using the `intercept404Writer.header`.
 The assembly of headers would take place in this.
 Since we don't have control on the other next / prior *Middleware* used,
 we would copy all the things stored *virtual Header storage* to
 the Real header inside `intercept404Writer.writer`. <br>
 Then, we would also redirect the processing
 to `intercept404Writer.notFoundHandler`. And we would also disable
 the `intercept404Writer.writer` in case of `404` due to same.

 5. Finally when the Response is ready to sent on *Wire* the final call to
 `Header() Header` is made. This time we return the real header inside
 `intercept404Writer.writer` instead of the *virtual Header storage*.

#### `Header() Header`

Let's Look at the first interface method `Header() Header`.

1. It needs to return the actual Header only in *Final Operation*. Meaning -
   When the Response is actually prepared to be sent out the wire.
2. In case It's called before any header manipulation must be done
   on top of `intercept404Writer.header` internal header.

```go
func (i *intercept404Writer) Header() http.Header {
    if i.header == nil && i.writer != nil {
        return i.writer.Header()
    }
    return i.header
}
```

We use the `intercept404Writer.header == nil` condition to check
if all the Header processing is complete.
And typically this is after the call to `WriteHeader` function.
Which we implement next.
The reason we check the `intercept404Writer.writer != nil` is
to avoid errors in case the actual writer is de-initialized.

#### `WriteHeader(statusCode int)`

We would look at next at the implementation of `WriteHeader(statusCode int)` function.

The following needs to done when the finalizing `WriteHeader` is called :

 1. All Headers stored in the `intercept404Writer.header` need to be
    transferred to the actual `intercept404Writer.writer.Header()`
 2. In case a `404` is encountered then call the respective handler
    needs to takeover as stored in `intercept404Writer.notFoundHandler`.
 3. If there is No `404` encountered then call the default `WriteHeader`.
 4. Finally to signal the completion of Header processing by
    `intercept404Writer.header = nil` since `WriteHeader` is only
    called once.

```go
func (i *intercept404Writer) WriteHeader(statusCode int) {
    for key, valuearr := range i.header {
        for _, val := range valuearr {
            i.writer.Header().Add(key, val)
        }
    }
    if statusCode == http.StatusNotFound {
        i.notFoundHandler.ServeHTTP(i.writer, i.req)
        i.writer = nil
    } else {
        i.writer.WriteHeader(statusCode)
    }
    i.header = nil
}
```

#### Write([]byte) (int, error)

Finally the `Write([]byte) (int, error)` used here needs to do the following:

 1. In case the `404` is not encountered do normal write operation
 using the internal `intercept404Writer.writer`.
 2. In other case of `404` return the dummy output convincing the
 caller that the write has taken place successfully.

```go
func (i *intercept404Writer) Write(b []byte) (int, error) {
    if i.writer != nil {
        return i.writer.Write(b)
    }
    return len(b), nil
}
```

### Wrapping Handler for Interceptor 404 functionality

Finally to use this we would need a way to change the default
`http.ResponseWriter`. We do that by creating a function to
return a wrapped Handler from the existing `http.FileServer`.

```go
func static404Handler(h, h404Handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        wr := &intercept404Writer{
            req:             r,
            writer:          w,
            notFoundHandler: h404Handler,
            header:          make(http.Header),
        }
        h.ServeHTTP(wr, r)
    })
}
```

This is kind of ***MiddleWare*** injection for the default `http.FileServer` like so..:

```go
    hfileServer := http.FileServer(http.Dir("public"))
    ...
    hfileServerWith404 := static404Handler(hfileServer, h404)
    ...
    http.Handle("/home/", http.StripPrefix("/home/", hfileServerWith404))
```

Now this can't be complete without the correct
`intercept404Writer.notFoundHandler`. So we implement that next.

## Create a Custom 404 Handler

This implementation would internally use the `return404` for
rendering the desired `404.html` page.

Note that this is another method of creating a ***MiddleWare*** for
`http.Handler` interface implementors.

```go
type handlerFor404 struct{}

// ServeHTTP to satisfy the http.Handler interface
func (h *handlerFor404) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    return404(w, r)
}

```

Very simple indeed. But this also does the same work as the
`static404Handler` function.

## Assembling it all together

```go
    hfileServer := http.FileServer(http.Dir("public"))
    h404 := &handlerFor404{}
    hfileServerWith404 := static404Handler(hfileServer, h404)
    http.Handle("/home/", http.StripPrefix("/home/", hfileServerWith404))
```

Here there is specific sequence of ***Middleware*** :

1. `http.Dir` Implements the `http.FileSystem` Interface
2. `static404Handler` wraps on top of `http.FileServer`
3. `http.StripPrefix` wraps on top of `static404Handler`

This is the desired hierarchy.
