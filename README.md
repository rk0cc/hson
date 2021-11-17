# Hash check JSON storage

[![HSON test](https://github.com/rk0cc/hson/actions/workflows/main_test.yml/badge.svg)](https://github.com/rk0cc/hson/actions/workflows/main_test.yml)
[![HSON build release](https://github.com/rk0cc/hson/actions/workflows/release.yml/badge.svg?branch=release)](https://github.com/rk0cc/hson/actions/workflows/release.yml)

HSON is JSON based data storage manager library with SHA-3 hashing as checksum for validating JSON context and compressed by GZip
as a side project of `will_pub`'s data storage solution

## File extension

HSON will generate a binary data file (and yes, it the same structre of JSON) with GZip compressed, the file must be named with this extensions:
* `.hson`
* `.hashjson`
It returned error if a file does not obey extension naming.

## API
|Name|Description|Return value|
|:-:|:-:|:--|
|`readHSON(path *C.char) *C.Char`|Read HSON file context from the path|Raw JSON string if read data successfully, or empty string if not|
|`writeHSON(context, path *C.Char) C.int`|Write HSON context to a file|<ul><li>`0` if write data successfully</li><li>`1` if failed during updating context (can not generate hash or non JSON data applied)</li><li>`2` if failed during writing to file</li></ul>|

## Cross platform ability

Read HSON with cross platform compability is yes and no since it depends the context has indented or not.
If the context is minified (which wrapping data in a single line), it higher chance will be recognized.
However, it not 100% worked if indented since different OS may use different new line in code.

## Note for macOS users

~~Please download go 1.17.x, clone this project and run `make build` by yourself. I have no Mac.~~

Good news! AMD64 based macOS is available now! (You still need to compile by yourself if using Apple Silicon)

### License
WTFPL
