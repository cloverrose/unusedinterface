# unusedinterface

`unusedinterface` finds unused interface in their package.

## Install

```sh
$ go install github.com/cloverrose/unusedinterface/cmd/unusedinterface@latest
```

### Or Build from source

```shell
$ go build -o bin/ ./cmd/...
```

### Or Install via aqua

https://aquaproj.github.io/


## Usage

```sh
$ go vet -vettool=`which unusedinterface` ./...
```
