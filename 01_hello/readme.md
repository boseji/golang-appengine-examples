# Appengine Hello World

This is basic App Engine example written in golang.

The example consist of only 2 files:

 * app.yaml - Configuration file needed for the App Engine to deploy

 * main.go - Application main file. Though naming it `main` might not be needed, it is essential that this file defines `package main` and contains the `func init()` as the entry point.

## Configuration Command

```shell
gcloud init
```

This would help to configure the specific Google Cloud project and location configuration.

> Note : Its best to select `us-central` if you are not sure which location to select. Plus its cheaper for hosting

## Deployment Command

```shell
gcloud app deploy ./app.yaml --version 1
```

This would deploy the app and set the **Version** to `1`

Adding a version helps to keep track of progress and later use specific instances of the Apps