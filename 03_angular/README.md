# Appengine Angular Example

This example is similar to the `02_static` but it also adds angular for 
client side processing.

Javascript library for **Angular** is served from the `assets` directory.

The template has been updated accordingly.
To support the **[Angular framework](https://angular.io/)**.

Similar Modification for **Go 1.12 Upgrade** as like the earlier example.

## TLDR;

To test out the **Angular App** locally :

```
go run main.go
```

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

```
Hello There! 
Golang Webapp is working
```

In the center of the page.
And a Form For Message Input that would be displayed.

Of course you can deploy the application the same way as earlier.

```shell
gcloud app deploy ./app.yaml --version 1
```
# Description

In this application example, we look at how we can use 
**[Angular JS](https://angular.io/)** along with
**Go's `html` templates** on Appengine.

## About Template / Angular parsing

In order to avoid confusion between template engine and Angular the 
following helps to encapsulate the parameters:

```html
        <ul>
            <li ng-repeat="m in messages track by $index">{{"{{"}}m{{"}}"}}</li>
        </ul>
```

This way the processing is only done on the client side.
Hence the angular and template engine can operate simultaneously.

Well we have colored the messages with a different CSS specifier under 
the `above-the-fold` container.

## No more Auto static sharing / hosting

We do not use the **automatic static sharing** like the 
[earlier example](../02_static/README.md)

Though this is optional we just wanted to show a new technique
of static file sharing/hosting.

Here is the modified [`app.yaml`](app.yaml) :
```yaml
runtime: go112

handlers:
# - url: /assets
#   static_dir: assets

- url: /.*
  script: auto
```

We have commented out the static sharing and now it still works.
This is new in the **Go 1.12 Appengine Upgrade**.

The *Go program* does the *static file hosting* :

```go
	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
```

Thanks to ***[Google Cloud](https://cloud.google.com)*** for 
making coding apps in **Golang** easier.

