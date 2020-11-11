# pipe

A command util package. Download here: [Pipe!](https://github.com/ToolPackage/pipe/releases/tag/v1.0).

## Commands

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
- use text as input:
  ```sh
  pipe in('text', 'hello world')=out
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

## About command paramter

For now, pipe supports three kinds of command parameter:
- number: including integer and float
- string: wrapped by single quote <b>'</b>
- bool: <b>true</b> or <b>false</b>

<b>NOTE:</b> Most terminal will recogonize space as a separator of two arguments and split them into a string array, so space in a string parameter will be removed. For examples: ```pipe in('text', 'hello world')=out``` will output ```helloworld``` finally.

<b>Solution:</b> Wrap the string literal with double quote <b>"</b>. E.g: ```"'hello world'"```.