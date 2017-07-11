# AppEngine Hello World

This is First App Engine example which is a derivative of the previous
`Hello World` example.

The example consist of 2 files:

 * app.yaml - Configuration file needed for the App Engine to deploy

 * main.go - Application main file. Though naming it `main` might not
 be needed, it is essential that this file defines `package main` and
 contains the `func init()` as the entry point.

## TLDR;

In order to run the Google AppEngine Dev environment use the following
command in the program directory:
```shell
dev_appserver.py app.yaml
````

*Yes, this is a python file* mapped to help you host the AppEngine
Development server instance.

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

`Working Golang App`

Next, you can also access the **Google AppEngine
Development Server** at `http://localhost:8000`

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

It would be interesting to note that there is no `main()` function in the
go file.

This is because the **Standard AppEngine Golang** instance
uses your code like a package.

There is a separate process running your code with the web server.

The distinction between **flexible** and **standard** is made by the
`app.yaml` file section:

```yaml
runtime: go
api_version: go1
```

Additional in the `app.yaml` there is a handler section which tells
the **Standard AppEngine Golang** Runtime to redirect
all request to the `golang app`

```yaml
handlers:
- url: /.*
  script: _go_app
```

Rest of the code is as explained in the earlier example.
