# go-stubs

Easily create `.go` files from stub files in your projects.

### Usage

```bash
go get github.com/nwby/go-stubs
```

Create a stub file:
```go
package stubs

type {{.Model}} struct {
	
}
```

Define a `StubDetails` struct and then that stub using `CreateStub`:
```
stub := StubDetails{
    Name:        "./stubs/model.go.stub",
    FileName:    "model.go",
    Destination: "./",
    Values: map[string]string{
        "Model": "User",
    },
}

stub.CreateStub()
```
