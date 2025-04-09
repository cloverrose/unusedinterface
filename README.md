# unusedinterface

`unusedinterface` unusedinterface find interface that are not used in their package.

## Install

```sh
$ go install github.com/cloverrose/unusedinterface/cmd/unusedinterface@latest
```

### Or Build from source

```shell
$ make build/unusedinterface
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
version: v1.64.8
name: custom-golangci-lint
destination: bin
plugins:
  - module: 'github.com/cloverrose/unusedinterface'
    import: 'github.com/cloverrose/unusedinterface'
    version: v0.2.4
```

`.golangci.yml`

```yaml
linters-settings:
  custom:
    unusedinterface:
      type: "module"
      description: unusedinterface find interface that are not used in their package
```

This repo is an example https://github.com/cloverrose/linter-playground
