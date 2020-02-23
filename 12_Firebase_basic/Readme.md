# Connecting to Firebase

We are going to look at integrating [Firebase][1] in our
**Appengine Project**.
Here we are only going to check if we are able to successfully
authenticate and create an instance of **`Firebase`**.

There are 2 important steps that we need to undertake:

1. Assign our **GCP project** to have **`Firebase`** integration
2. Get the **`Firebase`** credentials for testing out our project
   locally

## tl;dr

To test just run like a *normal Go program*:

```shell
go mod vendor
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/firebase/key.json
go run main.go
```

> **Note:** The `GOOGLE_APPLICATION_CREDENTIALS` is needed only for
> local app authentication to **`Firebase`**.
>
> Also the `go mod vendor` will add all dependencies in the `vendor`
> directory.

When the application is deployed to **Appengine** it would
authenticate using the *[default authentication scheme][2]*.
We use the **Service Account** keys downloaded from **`Firebase`**
project settings console.

And of course we can deploy this to *Appengine* :

```shell
gcloud app deploy ./app.yaml --version 1
```

## Description

In this project we use use the `go mod` with vendor support.
This is in accordance to our **Go 1.12 Appengine Upgrade**
discussed earlier.

One *interesting* item that you may observe with
the **[`go.mod`](go.mod)** file.
We have used a non existing *URI*.

```txt
module example.com/firebase_basic

go 1.12

require (
    cloud.google.com/go v0.53.0 // indirect
...
```

This is just symbolic naming.
Since this package need not be published or documented.
Hence we can use a dummy URI.

Another thing to note is the `vendoring`. We are using
the vendor directory for all dependencies.

We use the following command to initialize the `vendor`
directory.

```shell
go mod vendor
```

This command would *automagically* download and install dependencies
needed here.

### Firebase Initialize

In order to get **`Firebase`** integrated into the **GCP Project**
visit the URI:

<https://console.firebase.google.com/>

Then press the **Add Project** Button.

Next in the *Selection Menu* select you **GCP Project**.
Finally agree to using the [***Blaze***][3] Pay-as-you-go plan.
Its easier to account for and a recommended option.

That's it your **Firebase** connection to **GCP Project** is Done !

## Get Firebase Service Account

In order to access the **`Firebase`** we need to get **Access Token**.
This will be only needed to test out our code locally.

1. Visit the respective **GCP Project** in **[`Firebase Console`][4]**.

2. Click on the *gear icon* or **Settings**.

3. Next select **Project setting**.

4. Go to the **Service accounts** tab.

5. Click on **Generate new private key**.
   This would download **`.json`** file of the private key.
   > **Note**: that sometimes the *same service account has multiple keys*
   > This can be verified by visiting the [IAM page][5] and selecting
   > a **GCP Project** in **Google Cloud Console**.

Copy this `JSON` key into a fixed folder in the system.
This folder where `JSON` key is stored should be outside the project for safety reasons.

The **path** to the location where `JSON` key is stored needs to exported.

```shell
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/firebase/key.json
```

With this its possible to run this example locally.

 [1]:https://firebase.googleblog.com/2017/08/introducing-firebase-admin-sdk-for-go.html
 [2]:https://firebase.google.com/docs/admin/setup?authuser=0#initialize-sdk
 [3]:https://firebase.google.com/pricing
 [4]:https://console.firebase.google.com/
 [5]:https://console.cloud.google.com/iam-admin/serviceaccounts/
