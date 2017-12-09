# Install golang

## Install in linux

1. go to the https://golang.org/dl URL and download linux option.
2. open the console and go to the folder where you have downloaded golang.

`$ sudo tar -C /usr/local -xzvf goX.X.X.linux-xxx.tar.gz`

3. Add the folder bin to the PATH. Add the following exports to your ~/.bashrc

* `export GOROOT=/usr/local/go`
* `export GOPATH=${HOME}/work`
* `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`

4. Source the new environment

`$ source ~/.bashrc`

5. Testing it all

* `$ go env`
* `go version`
