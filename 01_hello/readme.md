# AppEngine Hello World

This is First App Engine example which is a derivative of the previous
`Hello World` example.

The example consist of 2 files:

* `app.yaml` - Configuration file needed for the App Engine to deploy

* `main.go` - Application main file. Though naming it `main` might not
 be needed, it is essential that this file defines `package main` and
 contains the `func init()` as the entry point.

## Upgrade to Go 1.12

### Module Initialization

```shell
go mod init github.com/boseji/golang-appengine-examples/01_hello
go mod tidy
```

### [`app.yaml`](app.yaml) Modifications

```yaml
runtime: go112
```

The `api_version` has been deprecated all controls are from `runtime`.

No need for the `handlers` section as It's a Single Entry App:

```yaml
handlers:
- url: /.*
  script: auto
```

### Program Modifications

1. We need to explicitly register the HTTP Handlers.
This is now true for both *Flexible* and *Standard Environments* for **Go 1.12**.
This helps to keep the program simple and easy to test out.

2. We also need to Get the specific `PORT` as an *Environment variable*.

### Ignore File

Yes we now need to add an explicit ignore file called **`.gcloudignore`**

For all the things we don't want to get included into the *Cloud Deployment*.

## tl;dr

To test just run like a *normal Go program*

```shell
go rum main.go
```

Yes this is new with **Go 1.12 Appengine Upgrade**.

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

`Working Golang App`

> *So Many Changes with **Go 1.12** - Makes things Easy but also quite different*

## Configuration Command

```shell
gcloud init
```

This would help to configure the specific Google Cloud project and
location configuration.

> Note : Its best to select `us-central` if you are not sure which
location to select. Plus its cheaper for hosting

## Deployment Command

```shell
gcloud app deploy ./app.yaml --version 1
```

This would deploy the app and set the **Version** to `1`

Adding a version helps to keep track of progress and later
use specific instances of the Apps

## Description

The Go program can now use any Framework and Library as needed.
This would enable to use the `go mod` with vendor support.
Another new things with **Go 1.12 Appengine Upgrade**.

Rest of the code is as explained in the [earlier example](../00_basic/readme.md).

## Appengine Edition in Go 1.12 Upgrade

The distinction between **flexible** and **standard** is made by the
`app.yaml` file section:

```yaml
runtime: go112
```

This is normal for **Standard Environment** without any *Scaling*.

Additional in the `app.yaml` there is a `env` section and `scaling` which tells
the **Flexible AppEngine Golang** Runtime to run the app in specific.

```yaml
env: flex

# This sample incurs costs to run on the App Engine flexible environment.
# The settings below are to reduce costs during testing and are not appropriate
# for production use. For more information, see:
# https://cloud.google.com/appengine/docs/flexible/python/configuring-your-app-with-app-yaml
manual_scaling:
  instances: 1
resources:
  cpu: 1
  memory_gb: 0.5
  disk_size_gb: 10
```
