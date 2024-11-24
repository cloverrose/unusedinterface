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

### Or golangci-lint custom plugin

https://golangci-lint.run/plugins/module-plugins/

Here are reference settings

`.custom-gcl.yml`

```yaml
# https://golangci-lint.run/plugins/module-plugins/
version: v1.62.0
name: custom-golangci-lint
destination: bin
plugins:
  - module: 'github.com/cloverrose/unusedinterface'
    import: 'github.com/cloverrose/unusedinterface'
    version: v0.2.3
```

`.golangci.yml`

```yaml
linters-settings:
  custom:
    unusedinterface:
      type: "module"
      description: check unused interface
```

This repo is an example https://github.com/cloverrose/linter-playground
