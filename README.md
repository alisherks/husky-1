# Husky

[![GoDoc Widget](https://godoc.org/github.com/go-courier/husky?status.svg)](https://godoc.org/github.com/go-courier/husky)
[![codecov](https://codecov.io/gh/go-courier/husky/branch/master/graph/badge.svg)](https://codecov.io/gh/go-courier/husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-courier/husky)](https://goreportcard.com/report/github.com/go-courier/husky)

Husky.js like, but pure in golang

## Usage

install

```
go get -u github.com/go-courier/husky/cmd/husky
```

## Configuration `.husky.yaml`

```yaml
hooks:
  # hook scripts
  pre-commit:
    - golangci-lint run
    - husky lint-staged
  commit-msg:
    - husky lint-commit
  
# list staged files do some pre-process and git add
lint-staged:
  "*.go":
    - gofmt -l -w

# commit msg rule only support conventionalcommits
lint-commit:
  # could check if this exists
  # regexp
  email: "^(.+@gmail.com|.+@qq.com)$"
```

Commit msg rule follow <https://www.conventionalcommits.org/en/v1.0.0/>

```
type(scope?): header

body?

footer?
```
