# Instructions

This project allows to test Glide as Goland package management (or *vendoring*). The glide-test project calls [go-sample-package](https://github.com/sh3lld00m/go-sample-package), that returns a simple static string stating a branch name.

First, clone this project and initialize Glide. What Glide will do is detect the dependencies in use in the project and create an initial ```glide.yaml``` file with the list of detected dependencies:

```bash
cd $GOPATH/src
git clone https://github.com/sh3lld00m/glide-test.git
cd glide-test/glidetest
glide init --non-interactive
```

## Master branch

There should be now a file named ```glide.yaml```in the root of the project directory. Edit the file and add ```version: master```at the same indentation level as package, as shown below:

```yaml
package: glide-test/glidetest
import:
  - package: github.com/sh3lld00m/go-sample-package
    version: master
    subpackages:
    - myutil
```

To instruct glide to fetch the dependencies, run:

```bash
glide up
```

```up``` is short for ```update```, and it will fetch whatever is specified in ```glide.yaml``` walking the dependency tree to make sure any dependencies of the dependencies are fetched and set to the proper version. There should be now a file called ```glide.lock``` created along with the download of the third party packages into the ```vendor```directory. This file contains the entire dependency tree pinned to specific commit ids.

Install now the dependencies running:

```bash
glide install
```

Because there's a ```glide.lock``` file, it will retrieve if missing from the ```vendor``` folder the dependency and sets it to the exact version specified in the ```glide.lock``` file.

Now continue with the standard Golang flow and run

```bash
go build
./glidetest
Version:master
```

Our third party package is being pulled from the HEAD of the master branch.

## Dev branch

The go-sample-package also has a second ```devel``` branch, which is just a slightly altered version of the ```master``` branch, meaning that it's returning the string ```devel``` instead of ```master```. Let's tell Glide that we want to use the ```devel``` branch by modifying ```glide.yaml``` accordingly:

```yaml
package: glide-test/glidetest
import:
  - package: github.com/sh3lld00m/go-sample-package
    version: devel
    subpackages:
    - myutil
```

Then, have Glide update and install the branch version of the samplepackage package, and then build and run again:

```bash
glide up
glide install
go build
./glidetest
Version: devel
