# App Engine Datastore Example 1

This example would show how to interface to the **[Google Cloud Datastore][3]** to an *App Engine `golang`` program*.

In this example we would be using the new **[`golang` Modules][1]** and upgrade to [App Engine Golang version 1.11][2].

## `golang` Modules

Let us first define the golang module for this project:

```shell
go mod init github.com/boseji/golang-appengine-examples/06_datastore
```

This command initializes the `golang` modules and creatates the file `go.mod`.

Next lets add the dependencies we would need for the project:

```shell
go get -u google.golang.org/appengine
go get -u google.golang.org/appengine/datastore
go get -u google.golang.org/appengine/log
```

After these command we would have a `go.sum` file that contains the hash of the version to be used. And the `go.mod` file would get updated also with the dependencies.

If you wish to see all the dependencies installed.

```shell
go list -m all
```

Subsequently if we use other `go get` commands the module would get automatically updated.



  [1]: https://blog.golang.org/using-go-modules
  [2]: https://cloud.google.com/appengine/docs/standard/go111/
  [3]: https://cloud.google.com/datastore/docs/concepts/overview