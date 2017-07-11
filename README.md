# Golang AppEngine Examples

[![Join the chat at https://gitter.im/golang-appengine-examples/Lobby](https://badges.gitter.im/golang-appengine-examples/Lobby.svg)](https://gitter.im/golang-appengine-examples/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Examples expalation showing:
 How to write Golang code for Google App Engine web App hosting services

 Each of the examples are documented within them selves with a separate `README.md`
 Please go through to understand the objective of each of them.

## Downloading the Repo

There would be 2 ways to do this

```shell
go get github.com/boseji/golang-appengine-examples
```

This would keep all the examples in you designated `GOPATH`

Or you can just download the Zip File:
https://github.com/boseji/golang-appengine-examples/archive/master.zip
And extract it where ever comfortable.

## Installing Tool Suit for AppEngine

We would first need to setup the tools to be able to use the App hosting for this
we have two options:

  - Cloud SDK - https://cloud.google.com/sdk/
  - AppEngine SDK (Orignal but older) - https://cloud.google.com/appengine/downloads

> *Here we would look at **Cloud SDK** as the choice*
> The commands hence forth would also comply to this assumption

### Windows Installation of *Cloud SDK*

First we would need to satisfy some dependencies:
  
  - Python 2.7 https://www.python.org/downloads/
    Choose the approporate one (x86 or x64)
  - Golang https://golang.org/dl/
  - Path configuration for `%GOPATH%` where the code 
  can be downloaded and one can store their own code.
  It helps to keep this consisten with the docs - https://golang.org/doc/code.html#GOPATH
    
    - Make sure that `%GOPATH%\bin` is also part of the local `%PATH%` variable 
    if not add it using `PATH=%PATH%;%GOPATH%\bin` while running things. 
    
    - However its best to set it in the 
     `SystemProperties -> Environment Variables`
    
    - To find this in the latest **Windows 10** type in the `Settings` search bar `path` and 
    then click on the option `Edit the system environment variable` that would open the 
    traditional systems console.

Now we are ready to proceed with the next phase of installaing the **Cloud SDK**

  - Download the **Cloud SDK** from https://cloud.google.com/sdk/docs/quickstart-windows  
  
  - Click on the link to Download the *SDK installer* 
  https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe
  
    - The installer is intutive not much to think there,    
    There is one **important point** make sure 
    that **Cloud SDK installation directory is in the PATH**
    - Restart the PC just to make sure that the path as already taken effect
  
  - Open an **Administrative Console**
   
    - First open `cmd` using the Run or search
    - On the task bar shpwing the `Command Prompt` right click on the Icon and 
    Again right click on `Command Prompt` to open another menu, finally in this 
    menu click `Run as administrator`
    - This should ask once for permissions and then open the `Adminstrator: Command Prompt`
  
  - Now in the `Adminstrator Command Prompt` execute the following commands:
    
    ```cmd
    # This would install both for Python and Golang
    gcloud components install app-engine-python app-engine-go
    ```

    The reason Python is also needed is to avail the dev tools for local serving 
    to test out the website configuration.

    Once this command is complete then type the following
    
    ```cmd
    gcloud components update
    ```

    This would initiate the Update for all the tools and SDK itself. 
    Though we just installed it its wise to make sure all the tools are up to date.

    Next we need to log into the **Google Cloud services** for that you 
    must have a gmail account or another company account affiliated to **Google**.

This concludes the basic setup for windows.

### Linux Installation of *Cloud SDK*

We would be looking specifically the **Debian** based linux version.

A **DockerFile** is also included to setup an container environment for you.

First we install the dependencies:

```shell
sudo apt update
sudo apt install -y curl python wget git apt-transport-https

# Follow the instructions at https://golang.org/dl/ and 
# https://golang.org/doc/install#tarball 
# to install Golang
```

We would need the `https` support later for **Cloud SDK**.

Next we download the required file for installation:

You can obtain the latest one from 
https://cloud.google.com/sdk/docs/quickstart-linux

```shell
wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-160.0.0-linux-x86_64.tar.gz
tar -xzf google-cloud-sdk-160.0.0-linux-x86_64.tar.gz
```

> Note: Where ever you have extracted the SDK that 
would become your default installation directory for *Cloud SDK*

Next to begin the installation for tools and SDK Components:

```shell
sh ./google-cloud-sdk/install.sh

./google-cloud-sdk/bin/gcloud components install app-engine-go app-engine-python

./google-cloud-sdk/bin/gcloud components update

# Optionally you can remove the installed zip to save space
rm google-cloud-sdk-160.0.0-linux-x86_64.tar.gz
```

The reason Python is also needed is to avail the dev tools for local 
serving to test out the website configuration.

Also just to be sure update all the components of **Clodu SDK** in the latter command.

> Note: Make sure that the Google *Cloud SDK* has been added to `PATH` if 
not add the following at the end of the `.bashrc` file 
```shell
export APP_ENGINE_DIR='/<YourInstalldir>/google-cloud-sdk'
# The next line updates PATH for the Google Cloud SDK.
if [ -f '$APP_ENGINE_DIR/path.bash.inc' ]; then 
  source '$APP_ENGINE_DIR/path.bash.inc' 
fi
# The next line enables shell command completion for gcloud.
if [ -f '$APP_ENGINE_DIR/completion.bash.inc' ]; then
  source '$APP_ENGINE_DIR/completion.bash.inc'
fi
export PATH=$PATH:$APP_ENGINE_DIR/bin
```

**Restart the PC or install the new `.bashrc` to update system path**

### Installing Required Golang packages

There are 2 main packages that are of interest here:

 - X-Crypto https://godoc.org/golang.org/x/crypto
 - App-Engine
 https://godoc.org/google.golang.org/appengine

We can recursively install both
```shell
go get -v golang.org/x/crypto/./...
go get -v google.golang.org/appengine/./...
```

***Now your environment is ready to move to the next Steps***


## TODO

  - [ ] Template based example with static hosting (02)
  - [ ] Shortcut link generator using Database (03)
  - [ ] User Login using OAuth (04)
  - [x] Microservices starting example (05)
  - [ ] Key Value store and JWT authentication service linked to OAuth (06)

*More to be added on your requests*
