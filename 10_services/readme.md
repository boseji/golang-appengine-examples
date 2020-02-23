# App Engine Services Example 1

This is an example shows how to implement the Micro-services in
`golang` on App Engine.

In this we would be implementing a **UUID Generator** Micro-service.

This example has also been upgraded to **Go 1.12 Appengine** requirements.

## tl;dr

To run this example use the command:

```shell
dev_appserver.py dispatch.yaml default/app.yaml uuid/app.yaml
```

This would make the servers available in 2 separate ports.

Since presently the **Google AppEngine Development Server** does not
support combined micro-services.

You can access them at `http://localhost:8080` and
`http://localhost:8081`

You can also run them independently in parallel :

```shell
cd default
export PORT=8080
go run main.go &
cd ../uuid
export PORT=8081
go run main.go
```

To deploy to the cloud there are two command files:

- For Windows `deploy.cmd`
- For Linux `deploy.sh`

Both would deploy with Version 1 which can be updated
in case multiple versions need to be supported of the
same API / micro-service.

**Note:** to Terminate the background Server Processes during
development use:

```shell
sudo pkill -15 go
sudo pkill -15 main
```

## Description

The project is divided into two separate sub-projects :

- **default** This would be normal frontend of base service
- **uuid** This is our special UUID Generator micro-service

In order to decide how to route the request there is a specific
`dispatch.yaml` file.

This file details how the requests are redirected:

```yaml
dispatch:
  # Default service serves simple hostname request.
  - url: "simple-sample.appspot.com/"
    service: default

  # Send all Service[UUID] traffic to the specific endpoint.
  - url: "*/uuid/*"
    service: uuid
```

- The First section details the dispatch or root or main access to the
   **default** service. Don't mind the `simple-sample.appspot.com/` this
   is only a way to express the path but not the exact name of the
   AppEngine project.

- The second section is special form of RegEx expression that selects
   any URI containing the `uuid` in the beginning of the host name.
   This means if we access `http://uuid.testproject.appspot.com` that
   take me to this service - assuming `testproject` is the name of our
   Appengine project. In the special case of `https` it becomes
   `https://uuid_dot_testproject.appspot.com`. This is kind of
   convention by **Google AppEngine Golang** Standard instances.

For More insight into `dispatch.yaml` go to:
<https://cloud.google.com/appengine/docs/standard/go112/reference/dispatch-yaml>
