# MyGo_Learning
My hands on Trial of Go Tutorials

### Environment Variables: 
 * **`GOROOT`** : Go installation Directory
 *  `export PATH=$PATH:/$GOROOT/bin`
 * **`GOPATH`**: Defines Go-workspace. Kind of pushes towards a monolythic workspace to use, so that necessary libraries/packages/modules will be downloaded to GOROOT directory & can be used in the code developed under GOPATH
 * Normally `$GOROOT/lib` directory is empty. So we will have to download all required initial libraries by **`go get github.com/nsf/gocode`**
* You can find values of all Go Env Variables by `$ go env`

![image](https://user-images.githubusercontent.com/24938159/98212317-6a470b00-1f69-11eb-8ec1-0818897e2873.png)

## Examples
### [00-Hello.go](./00-hello.go)

### [1.1-hello-quote.go](./1.1.hello-quote.go)
 * When the module is not found by go by itself. Follow below steps
 * `$ go mod init <code name>` : It will create a new mod
 * `$ go run <code>.go` : This will download the  module & version. Then will execute the code.

![image](https://user-images.githubusercontent.com/24938159/98212594-d3c71980-1f69-11eb-9e61-20fedaa2b0ff.png)

```sh

C:\Users\HPANIGR\go\go-workspace-test> go mod init 1.1.go-quote.go
go: creating new go.mod: module 1.1.go-quote.go

C:\Users\HPANIGR\go\go-workspace-test>go run 1.1.go-quote.go

go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Don't communicate by sharing memory, share memory by communicating.

```
