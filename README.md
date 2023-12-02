# `do-func-golang-example`

## Install

```bash
git clone https://github.com/egorsmkv/do-func-golang-example

doctl serverless install

doctl serverless deploy do-functions-golang-example --verbose-build --remote-build
```

## Development

```bash
golangci-lint run packages/sample/hello
```
