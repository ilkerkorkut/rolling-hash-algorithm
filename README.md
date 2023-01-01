# Rolling Hash Algorithm
[![CI](https://github.com/ilkerkorkut/rolling-hash-algorithm/workflows/CI/badge.svg?event=push)](https://github.com/ilkerkorkut/rolling-hash-algorithm/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/ilkerkorkut/rolling-hash-algorithm)](https://hub.docker.com/r/ilkerkorkut/rolling-hash-algorithm/)
[![Go Report Card](https://goreportcard.com/badge/github.com/ilkerkorkut/rolling-hash-algorithm)](https://goreportcard.com/report/github.com/ilkerkorkut/rolling-hash-algorithm)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ilkerkorkut_rolling-hash-algorithm&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ilkerkorkut_rolling-hash-algorithm)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=ilkerkorkut_rolling-hash-algorithm&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=ilkerkorkut_rolling-hash-algorithm)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=ilkerkorkut_rolling-hash-algorithm&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=ilkerkorkut_rolling-hash-algorithm)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=ilkerkorkut_rolling-hash-algorithm&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=ilkerkorkut_rolling-hash-algorithm)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ilkerkorkut_rolling-hash-algorithm&metric=coverage)](https://sonarcloud.io/summary/new_code?id=ilkerkorkut_rolling-hash-algorithm)
[![Release](https://img.shields.io/github/release/ilkerkorkut/rolling-hash-algorithm.svg)](https://github.com/ilkerkorkut/rolling-hash-algorithm/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/ilkerkorkut/rolling-hash-algorithm)](https://github.com/ilkerkorkut/rolling-hash-algorithm)
[![License](https://img.shields.io/badge/MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Build Locally
```shell
make build
```

## Usage
Create a signature file from original file:

```shell
rha signature -f original.txt -s signature.txt --chunk-size 2
```

Create delta from new file and signature file:

```shell
rha delta -s signature.txt -d delta.txt -n new.txt --chunk-size 2
```


## References
[https://en.wikipedia.org/wiki/Adler-32](https://en.wikipedia.org/wiki/Adler-32)
[https://en.wikipedia.org/wiki/Rolling_hash](https://en.wikipedia.org/wiki/Rolling_hash)

## Development
This project requires below tools while developing:
- [Golang 1.19](https://golang.org/doc/go1.19)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ pre-commit install -c build/ci/.pre-commit-config.yaml
```
