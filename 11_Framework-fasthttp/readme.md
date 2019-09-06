# Using `fasthttp` Framework on Appengine

We are going to try using **[`fasthttp`][1]**
as a frameworks example.

This is much different from `net/http` in ways it handle request.
Also [`fasthttp` claims to be the fastest][2] when it comes to
raw performance. The secret is in using `bytes.Buffer` type.
This allows nearly zero reallocation.

## TLDR;

To test just run like a *normal Go program*

```shell
go run *.go
```

or using `reflex` tool:

```shell
reflex -d none -s -r \.go$ -- go run *.go
```

Yes this is new with **Go 1.12 Appengine Upgrade**.

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

`Working Golang App`

> *So Many Changes with **Go 1.12** - Makes things Easy but also quite different*
>
> Now we are on the ***Fastest!!*** web framework `fasthttp`
> Much more lean Than `net/http`
> With nearly zero re-allocation in memory

And of course we can deploy this to *Appengine* :

```shell
gcloud app deploy ./app.yaml --version 1
```

## Description

The Go program can now use any Framework and Library as needed.
This would enable to use the `go mod` with vendor support.
Another new things with **Go 1.12 Appengine Upgrade**.

Here we have added the `fasthttp` as dependency:

```go.mod
module github.com/boseji/golang-appengine-examples/11_Framework-fasthttp

go 1.12

require (
    github.com/klauspost/compress v1.8.2 // indirect
    github.com/klauspost/cpuid v1.2.1 // indirect
    github.com/valyala/fasthttp v1.4.0
)
```

After the module initialization one needs to do the `tidy` to
make sure all is well.

```shell
go mod init github.com/boseji/golang-appengine-examples/11_Framework-fasthttp
go get -u github.com/valyala/fasthttp
go mod tidy
```

 [1]:https://github.com/valyala/fasthttp
 [2]:https://github.com/valyala/fasthttp#http-server-performance-comparison-with-nethttp
