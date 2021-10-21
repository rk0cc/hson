# Hash check JSON storage

[![HSON test](https://github.com/rk0cc/hson/actions/workflows/go.yml/badge.svg)](https://github.com/rk0cc/hson/actions/workflows/go.yml)

HSON is JSON based data storage manager library with SHA-3 hashing as checksum for validating JSON context and compressed by GZip
as a side project of `will_pub`'s data storage solution

## API
|Name|Description|Return value|
|:-:|:-:|:--|
|`readHSON(path *C.char) *C.Char`|Read HSON file context from the path|Raw JSON string if read data successfully, or empty string if not|
|`writeHSON(context, path *C.Char) int`|Write HSON context to a file|<ul><li>`0` if write data successfully</li><li>`1` if failed during updating context (can not generate hash or non JSON data applied)</li><li>`2` if failed during writing to file</li></ul>|

## Note for macOS users

Please download go 1.17.x, clone this project and run `make build` by yourself. I have no Mac.

### License
WTFPL
