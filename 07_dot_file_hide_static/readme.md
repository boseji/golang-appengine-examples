# Dot File restrict for Static File Serving on Appengine

Though the [earlier example](../06_error_routing_static) solves our
problem with `404`, there is more to the picture.

While serving static files we often have the *dot-files*.
These too get served out to public in the `http.FileServer`.
For example `.gitignore` or `.gcloudignore` are all *dot-files*

This is due to the use of [`http.Dir`](https://godoc.org/net/http#Dir)
implementing the internal `http.FileSystem` interface.

Let's look warning snippet from
[`http.Dir` implementation](https://golang.org/src/net/http/fs.go#L26):

```go
// Note that Dir will allow access to files and directories starting with a
// period, which could expose sensitive directories like a .git directory or
// sensitive files like .htpasswd. To exclude files with a leading period,
// remove the files/directories from the server or create a custom FileSystem
// implementation.
```

There we have it. That's the **Problem** we are going to solve here.

Though there is already some given solutions to this directly
from the [`http.FileServer` docs](https://godoc.org/net/http#FileServer).

It's called the
[*DotFileHiding Example*](https://godoc.org/net/http#example-FileServer--DotFileHiding).

We would be enhancing our know-how on :

- `.gcloudignore` file <https://cloud.google.com/sdk/gcloud/reference/topic/gcloudignore>
- Also looking at debugging in [*VSCode IDE*](https://code.visualstudio.com/)
- Finally we would be improving upon the
  [previous `404` handling](../06_error_routing_static/). Since now
  we need to Handle other errors as well.

## tl;dr

To Execute the project use:

```shell
go run *.go
```

Upon visiting the default `http://localhost:8080` location
you would be greeted by :

`Working Golang App`

Well that we already know. Let's now try to visit
the static homepage `http://localhost:8080/home`.

> **This is a Home Page**

Ya a simple HTML page hosted here.

Here are a few valid URLs :

- <http://localhost:8080>
- <http://localhost:8080/home>
- <http://localhost:8080/home/1.txt>
- <http://localhost:8080/home/2.txt>
- <http://localhost:8080/home/css/index.css>

For invalid URLs our custom [`404.html`](public/404.html) page.
You would be greeted by a message :

> **Looks Like you have Stumbled upon something not here**

Let's now try to find our *Hidden file* now:

- <http://localhost:8000/home/.test>

Again you would greeted by the same custom [`404.html`](public/404.html) page.

We have successfully blocked access to *dot-files*.

This would also work in Appengine:

```shell
gcloud app deploy ./app.yaml --version 1
```

## Description

There were 4 main objectives that we started with this development:

1. Being able to prevent serving of *dot-files*
2. Improve upon the `404` handling
3. Enhance the `.gcloudignore` file
4. Do a better hib at development environment

Let's look at each one in detail and how we achieve them.

### Prevent *dot-file* Serving

This was completely derived from the
[*DotFileHiding Example*](https://godoc.org/net/http#example-FileServer--DotFileHiding).

The only exception in case of the error returned when a *dot-file* was found.
We wanted to have a Not-found instead of Forbidden code being generated.

```go
func (fs dotFileHidingFileSystem) Open(name string) (http.File, error) {
    if containsDotFile(name) {
        return nil, os.ErrNotExist // Modification to Allow the 404 to work
    }
    ...
```

### Improve upon the `404` handling

In our test we found that some times the `WriteHeader` function was
not getting called. Hence the full copy of headers from
*virtual Header storage* was not happening.

Thus we moved the copy of the headers from *virtual Header storage*
to actual `intercept404Writer.writer`.

Here is the modification to the two functions:

```go
func (i *intercept404Writer) Header() http.Header {
    if /*i.header == nil &&*/ i.writer != nil {
        for key, valuearr := range i.header {
            for _, val := range valuearr {
                i.writer.Header().Add(key, val)
            }
        }
        return i.writer.Header()
    }
    return i.header
}

func (i *intercept404Writer) WriteHeader(statusCode int) {
    /*for key, valuearr := range i.header {
        for _, val := range valuearr {
            i.writer.Header().Add(key, val)
        }
    }*/
    if statusCode == http.StatusNotFound /* || statusCode == http.StatusForbidden*/ {
        i.notFoundHandler.ServeHTTP(i.writer, i.req)
        i.writer = nil
    } else {
        i.writer.WriteHeader(statusCode)
    }
    i.header = nil
}
```

### Enhance the `.gcloudignore` file

We found more information about the `.gcloudignore` here:

<https://cloud.google.com/sdk/gcloud/reference/topic/gcloudignore>

Hence now we have included a lot more into it:

```gitignore
.gcloudignore
*.md
*.json
.vscode/
node_modules/
.git
.gitignore
*~
```

This would help us avoid uploading many unnecessary files up Appengine.

### Better Development Environment

We generally code using the [***VSCode IDE***](https://code.visualstudio.com/)
for *Golang*.
Here with a simple modification ot the `launch.json` -
we were able to automatically debug the program.

Here is how our current `launch.json` looks like:

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {},
            "args": []
        }
    ]
}
```

The specific change we did was to set the`"program"` value as `"${workspaceFolder}"`.
Also we started the IDE using `code .` from the directory of the actual program.
Rather than the root of this repository.

Clicking on debug we could actually set breakpoints and observe the operations.
This helped us to dig out the problem that `WriteHeader` is not called always.

#### New Automation Tools

Additionally, we found two other repeated build tools - which initiate commands
when files are modified:

- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
This is a *Golang* based utility that can be used with *Golang projects*.
Actually they can be used with any type of build requirement -
where repeated execution of a command needs to be performed based on specific
files in the file-system getting updated. This is specifically targeted
for Golang programs using the `http.ListenAndServe` functionality.

- [reflex](https://github.com/cespare/reflex) Again a *Golang* based utility
  for more generic activity. More akin to [*nodemon*](https://nodemon.io/)
  from the ***Node JS*** world. We decided to give this one a try.

#### Using *reflex* tool

To install *reflex* we do :

```shell
go get -u github.com/cespare/reflex
```

It has `Go 1.11 or Greater`  dependency. Does not work with older versions.

To do the Build on watch this is the command for *reflex* :

```shell
reflex -d none -s -r \.go$ -- go run *.go
```

Where the following flags have the respective meaning:

- `-d none` This is to enable the output from target program. There are three
  to choose from `fancy`, `none` and `plain`
- `-s` Runs the program like a service and restarts it when ever change is
  detected in the files.
- `-r \.go$` This is the regex filter for file extension to observe.
- `-- go run *.go` is our actual command that would get re-run upon file change.

The whole *reflex* process can be exited using key combination of `Ctrl+C`.
Like for any other program it would receive SIGINT to terminate the child processes
meaning the command it is running as well as itself.
