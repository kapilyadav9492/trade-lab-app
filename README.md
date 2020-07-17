# Trade-Lab Application

Trade-Lab application project contains three API's :
1. <b>Health Care API </b> Checking server health. i.e. server running or not.
2. <b>Currency API</b> passing currency symbol and getting data from in-memory (Struct).
3. <b>Get ALL Currency API </b> Return all currency exist in the in-momeory data.


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Whats is WebSocket ?
WebSocket is a computer communications protocol, providing full-duplex communication channels over a single TCP connection.
The WebSocket protocol was standardized by the IETF as RFC 6455 in 2011, and the WebSocket API in Web IDL is being standardized by the W3C. WebSocket is distinct from HTTP.

#####Reference Links

[Golang Websocket](https://godoc.org/golang.org/x/net/websocket)
<br/>[Gorilla Websocket](https://godoc.org/github.com/gorilla/websocket)
<br/>[Nhooyr Websocket](https://godoc.org/nhooyr.io/websocket)



### Installation go(lang)

<b>Code Assignment </b> project is <b> Go language </b> based.
<br/>Install Go with [homebrew](https://brew.sh/):

```Shell
sudo brew install go
```

with [apt](https://packages.qa.debian.org/a/apt.html)-get:

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)


<br/>Install packages from github to my gopath/
```Shell
go get -u github.com/gorilla/mux
```

## Usage

TODO: Write usage instructions
####Login instruction
<b>Step One</b> — Get The users name and Password from postman or client
<br/><b>Step Two</b> — Generate access token with help of secret key. For best practice secret key stores into bash_profile files.

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b code-assignment`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin code-assignment`
5. Submit a pull request.

## go-testing
Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form
        
    $func TestXxx(*testing.T)

The test directory contains tests of the Go tool chain and runtime.
It includes black box tests, regression tests, and error output tests.

A simple test function looks like this:

    func TestAbs(t *testing.T) {
        got := Abs(-1)
        if got != 1 {
            t.Errorf("Abs(-1) = %d; want 1", got)
        }
    }

To run just these tests, execute:

    $ go test -run NameOfTest
    $ go test -run Test_Login
    
   Some Commands for test files -
    
    go test -run ''      # Run all tests.
    go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
    go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
    go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
   

Standard library tests should be written as regular Go tests in the appropriate package.

The tool chain and runtime also have regular Go tests in their packages.
The main reasons to add a new test to this directory are:

* it is most naturally expressed using the test runner; or
* it is also applicable to `gccgo` and other Go tool chains.