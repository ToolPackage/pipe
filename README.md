# pipe

A command util package. Download here: [Pipe!](https://github.com/ToolPackage/pipe/releases/tag/v1.0).

## commands

### in & out

examples:
- read from stdin and output to stdout:
  ```sh
  echo asd | pipe in=out
  ```
- read from file and output to file:
  ```sh
  pipe in('file', './test.txt')=out('file', './test1.txt')
  ```
- use label parameters:
  ```sh
  echo asd | pipe in=out(type:'file', filename:'./test1.txt')
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
  echo asd | pipe in=gzip.compress=out('file', './test.gzip')
  ```
- decompress:
  ```sh
  pipe in(type: 'file', './test.gzip')=gzip.decompress=out
  ```

### json

examples:
- pretty:
  ```sh
  pipe in('file', './test.json')=json.pretty=out
  ```
- get: need one parameter to indicate json path, more details see [gjson](https://github.com/tidwall/gjson)
  ```sh
  echo "{"name":{"first":[-999]}}" | pipe in=json.get('name.fisrt.0')=out
  ```
  