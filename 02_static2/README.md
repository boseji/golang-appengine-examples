# Appengine Angular Example

This example is similar to the `02_static` but it also adds angular for client side 
processing.

Javascript library for **Angular** is served from the `assets` directory.

The template has been updated accordingly.
To support the Angular framework.

## TLDR;

In order to run the Google AppEngine Dev environment use the following
command in the program directory:

```shell
dev_appserver.py app.yaml
````

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

```
Hello There! 
Golang Webapp is working
```

In the center of the page.

Ofcourse you can deploy the application the same way as earlier.

```shell
gcloud app deploy ./app.yaml --version 1
```

#### Note: About Template / Angular parsing

In order to avoid confusion between template engine and Angular the folloing helps to encapsulate the parameters:

```html
        <ul>
            <li ng-repeat="m in messages track by $index">{{"{{"}}m{{"}}"}}</li>
        </ul>
```

This way the processing is only done on the client side.
Hence the angular and template engine can operate simultaneously.