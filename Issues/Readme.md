## Listing Issues faced & resolution:

### Issue-01. [**$GOPATH/go.mod exists but should not**](https://stackoverflow.com/questions/59144120/gopath-go-mod-exists-but-should-not-in-aws-elastic-beanstalk/62062562#62062562): 

   * When GOPATH is set go get installs the requested module to the path provided in GOPATH but if you use .mod file, it uses the working directory.
   * If you have both GOPATH set and have the .mod file. You can either `unset $GOPATH` or `delete go.mod`
   
