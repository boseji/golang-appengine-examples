# App Engine Services Example 1

This is an example shows how to implement the Microservices in
`golang` on App Engine.

In this we would be implementing a **UUID Generator** Microservice.

## TLDR;

To run this example use the command:

```shell
dev_appserver.py dispatch.yaml default/app.yaml uuid/app.yaml
```

This would make the servers available in 2 separate ports.

Since presently the **Google AppEngine Development Server** does not
support combined microservices.

You can access them at `http://localhost:8080` and
`http://localhost:8081`

To deploy to the cloud there are two command files:

 - For Windows `deploy.cmd`
 - For Linux `deloy.sh`

Both would deploy with Version 1 which can be updated
in case multiple versions need to be supported of the
same API / microservice.

## Description

The project is divided into two separate sub-projects :

  - **default** This would be normal frontend of base service
  - **uuid** This is our special UUID Generator microservice

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

  - The First section details the dipatch or root or main access to the
   **default** service. Don't mind the `simple-sample.appspot.com/` this
   is only a way to express the path but not the exact name of the
   AppEngine project.

   - The second section is special form of RegEx expression that selects
   any URI containing the `uuid` in the beginning of the host name.
   This means if I access `http://uuid.testproject.appspot.com` that
   take me to this service. In the special case of `https` it becomes
   `https://uuid_dot_testproject.appspot.com`. This is kind of
   convention by **Google AppEngine Golang** Standard instances.
