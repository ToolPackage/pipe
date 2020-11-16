# pipe

[![PkgGoDev](https://pkg.go.dev/badge/github.com/ToolPackage/pipe)](https://pkg.go.dev/github.com/ToolPackage/pipe)
<!-- [![CircleCI](https://circleci.com/gh/goccy/go-yaml.svg?style=shield)](https://circleci.com/gh/goccy/go-yaml)
[![codecov](https://codecov.io/gh/goccy/go-yaml/branch/master/graph/badge.svg)](https://codecov.io/gh/goccy/go-yaml) -->
[![Go Report Card](https://goreportcard.com/badge/github.com/ToolPackage/pipe)](https://goreportcard.com/report/github.com/ToolPackage/pipe)

A command util integrated with multiple functions. Download here: [Pipe!](https://github.com/ToolPackage/pipe/releases/tag/v1.0).

## Usage

```
usages:
  in -> in(): read input from stdin
    file -> in.file(name: string): read input from file
    text -> in.text(value: string): use text value as input
  out -> out(): output pipe data to stdout
    file -> out.file(name: string): output pipe data to file
  base64
    encode -> base64.encode()
    decode -> base64.decode()
  gzip
    compress -> gzip.compress()
    decompress -> gzip.decompress()
  color
    json -> color.json(): colorize input in json syntax
  json
    pretty -> json.pretty(): pretty json input
    get -> json.get(path: string): get value from json input by path
  regexp
    test -> regexp.test(pattern: string): test input with regexp pattern
  http
    get -> http.get(url: string, headers?: dict, outputMode?: 'body' | 'raw'): create http get request
  url
    encode -> url.encode()
    decode -> url.decode()
  text
    cut -> text.cut(start: int, end?: int): extract substring
    replace -> text.replace(old: string, new: string): replace substring
    repeat -> text.repeat(n: int): repeat input n times
  html
    pretty -> html.pretty()
```

## Examples

### in & out

examples:
- read from stdin and output to stdout:
  ```sh
  echo asd | pipe in=out
  ```
- read from file and output to file:
  ```sh
  pipe in.file('./test.txt')=out.file('./test1.txt')
  ```
- use text as input:
  ```sh
  pipe in.text('hello world')=out
  ```
- use labeled parameters:
  ```sh
  echo asd | pipe in=out.file(name:'./test1.txt')
  ```

### base64

examples:
- encode:
  ```sh
  echo asd | pipe in=base64.encode=out
  ```
- decode:
  ```sh
  echo asd | pipe in=base64.decode=out
  ```

### gzip

examples:
- compress:
  ```sh
  echo asd | pipe in=gzip.compress=out.file('./test.gzip')
  ```
- decompress:
  ```sh
  pipe in.file('./test.gzip')=gzip.decompress=out
  ```

### json

examples:
- pretty:
  ```sh
  pipe in.file('./test.json')=json.pretty=out
  ```
- get: need one parameter to indicate json path, more details see [gjson](https://github.com/tidwall/gjson)
  ```sh
  echo "{"name":{"first":[-999]}}" | pipe in=json.get('name.fisrt.0')=out
  ```

### regexp

examples:
- test:
  ```sh
  pipe in('text', '192.168.1.1')=regexp.test(pattern: '^([0-9]{1,3})(\.([0-9]{1,3})){3}$')=out
  ```

## About command parameter

For now, pipe supports three kinds of command parameter:
- number: including integer and float
- string: wrapped by single quote <b>'</b>
- bool: <b>true</b> or <b>false</b>

<b>NOTE:</b> Most terminal will recogonize space as a separator of two arguments and split them into a string array, so space in a string parameter will be removed. For examples: ```pipe in('text', 'hello world')=out``` will output ```helloworld``` finally.

<b>Solution:</b> Wrap the string literal with double quote <b>"</b>. E.g: ```"'hello world'"```.