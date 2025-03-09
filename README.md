# GoFt

[![Go Reference](https://pkg.go.dev/badge/github.com/souhoc/goft.svg)](https://pkg.go.dev/github.com/souhoc/goft)

Goft is a [Go](https://golang.org/) package that provides low level 
bindings to the [42 API](https://api.intra.42.fr/apidoc).
Sturcts are generated automaticly from the apidoc.

**For help with this package or general Go discussion, please contact me on discord (`sou.hoc`).**

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/souhoc/goft
```

### Usage

Import the package into your project.

```go
import "github.com/souhoc/goft"
```

Construct a new fortytwo session which can be used to access the variety of 
42 API endpoints.
*Here you can get a user*

```go
conf := clientcredentials.Config{
	ClientID:     "Your Client ID",
	ClientSecret: "Your Client Secret",
	TokenURL:     goft.TokenUrl,
	AuthStyle:    oauth2.AuthStyleInParams,
}
client := conf.Client(context.Background())
ft := goft.New(client)

var user goft.User
_, err := ft.Get("/users/" + login, &user)
if err != nil {
	return err
}

```

See Documentation and Examples below for more detailed information.

## Documentation

**NOTICE**: This library and the 42 API are unfinished.
Because of that there may be major changes to library in the future.

*WIP*

## Examples

*WIP*

## Troubleshooting
For help with common problems please open an issue

## Contributing
Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be
discussed.  
- Try to match current naming conventions as closely as possible.  
- This package is intended to be a low level direct mapping of the Discord API, 
so please avoid adding enhancements outside of that scope without first 
discussing it.
- Create a Pull Request with your changes against the master branch.
