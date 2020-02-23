# Appengine Google Cloud Store

This is the storage service provided by GCP - 
***[Google Cloud Store](https://cloud.google.com/storage/)***.

This is the default location for storing files and bigger data-objects.
It can also be used to store the uploaded files from your web app.

Please note this is not to be confused with 
**[Cloud Datastore / Firestore](https://cloud.google.com/datastore/)**. 
That is a key-value pair document / NoSQL database engine.

**Google Cloud Store** or simply *the Store* is a file storage for
long term persistent data and possibly larger data-sets.

**Google Cloud** provides ***free 5GB of cumulative capacity*** for
*[free-tier](https://cloud.google.com/free/)*.
In most cases this is more than enough to do.

For even larger storage requirements one can use 
**[Google Drive API](https://developers.google.com/drive/)**.

### Module Configuration Step

In addition to the normal steps we need to add
the specific package for *GCP Store* :

```shell
go mod init github.com/boseji/golang-appengine-examples/04_cloudstore
go mod tidy
go get -u cloud.google.com/go/storage
```

## TLDR;

To test the App :

```shell
go mod download # This would download the require modules
export GCP_PROJECT_ID="<Your GCP Project ID>"
go run *.go
```

Make sure to check and then fill your GCP project ID
in place `<Your GCP Project ID>`. Its typically the 
name on top of your *[GCP Console](https://console.cloud.google.com)*.

Of course you can deploy the application the same way as earlier.

```shell
gcloud app deploy ./app.yaml --version 1
```

https://cloud.google.com/appengine/docs/standard/go112/runtime#environment_variables