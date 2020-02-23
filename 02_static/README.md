# Appengine Hosting static files

This example shows how static files are hosted along with the golang program.
The project consists of 2 parts:

* Assets - Containing images , CSS, Javascript files etc.
* templates - HTML rendering files

We would also have the similar `app.yaml` file like the pervious
**[`01_hello`](../01_hello/readme.md)** example.
However this time with a slight twist.

Additional files:
  
* `assets\css\index.css` - This is the CSS file used by the Index Page template
* `templates\index.gohtml` - This is the Template file that would be rendered
     for Index page

Similar Modification for **Go 1.12 Upgrade** as like the earlier example.

Only additional modification would the inclusion of `handlers`.

## tl;dr

Test it out using:

```shell
go run main.go
```

The Server is accessible at `http://localhost:8080`

If you visit this in the browser it should display

```shell
Hello There!
Golang Webapp is working
```

In the center of the page.

Of course you can deploy the application the same way as earlier.

```shell
gcloud app deploy ./app.yaml --version 1
```

### Note for Development

In order to facilitate easier development the static `assets` directory
is also served when the `PORT` is not set.

```go
http.Handle("/assets/",
      http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
```

## Description

This example uses the **Standard AppEngine Golang** instance in the `app.yaml`

### Static Routing

The important change in the `app.yaml` file :

```yaml
runtime: go112

handlers:
- url: /assets
  static_dir: assets

- url: /.*
  script: auto
```

The line with `url: /assets` specifies a **Static Path** that would be served.
And the content to be served is a **Static Directory** given by
`static_dir: assets`.

This is what gets used to source the `index.css` file in the line:

```html
<link rel="stylesheet" href="assets/css/index.css">
```

Of the main page.

Similarly one can add Javascript and images as needed into the assets
directory and they would automatically be available for use in the webpages.

#### Template

You might have already noticed that we are using a strange name
`index.gohtml` for the template. This is just a *Extension* name
convention to help identify **template files** from normal *html* files.

The format of the template used here looks like any normal HTML page.
However that be quickly enhanced to provide additional features using
the **golang Templating Engine**

In the current program we are using the `html/template` package to be
able to render HTML pages. Similar package is available for `text`
where arbitrary template rendering is needed.

Let's have a deeper look at how templates are working here.

#### Templates Usage

First we need to tell the template engine to load all the templates.

```go
var tpl *template.Template

func init() {
  var err error
  tpl, err = template.ParseGlob("templates/*.gohtml")
    ....
```

Here you might notice that we are using a global `var tpl *template.Template`

The reason for using that is , we wish to be able to access this in any of our
handler instances created with

```go
func indexHandler(w http.ResponseWriter, r *http.Request) {
```

Additionally this is a pointer since we can have multiple templates and the
loaded templates is basically an **array of templates**

The parsing / loading of the templates is done using the
`template.ParseGlob("templates/*.gohtml")` line.
This processes all the templates under the *template directory* with
extension of `.gohtml`

The extension `.gohtml` is just a preferential convention to indicate
that the file is a template. One might directly use `.html` also,
but this is a better for separation-of-concerns.

Next is actually rendering the template that has been loaded.
By rendering we mean that template is compiled, added with data
and finally sent to the client.

This three step process is delivered by:

```go
err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
```

Here we are rendering the `index.gohtml` page template.
The last parameter is `nil`, this is actually the place where data is sent
to be processed into the rendered page.
