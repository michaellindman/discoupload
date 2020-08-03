# Discoupload

Application for uploading files to the [discourse](https://www.discourse.org) forum software using the [discourse API](https://docs.discourse.org).

## Requirements

To build a golang development environment and tools are needed with a proper `GOPATH`.

For instructions please see:

* [Golang Install](https://golang.org/doc/install)
* [SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

## Building

From the root of the source tree run:

```sh
make debs
make build
```

## Usage

Before running a `config.yml` file containing the api settings is needed, a template of which can be found below:

```yaml
api:
  url: <forum url>
  key: <api key>
  username: <api username>
```

and can be ran with:

```sh
./discoupload -file /path/to/upload/file
```

### As a go module

The upload function has been exported so it can be used by other go applications. With a correctly configured golang toolchain run:

```sh
go get github.com/michaellindman/discoupload
```

#### Example

```go

var (
    key      = "<api-key>"
    username = "<api-username>"
    url      = "<forum url>"
)

upload, err := Upload(key, username, url, "path/to/upload/file")
if err != nil {
    log.Println(err)
}
fmt.Printf("Uploaded %v (%v): %v\n", upload["original_filename"], upload["human_filesize"], upload["url"])
```

## License

```text
MIT License Copyright (c) 2020 Michael Lindman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished
to do so, subject to the following conditions:

The above copyright notice and this permission notice (including the next
paragraph) shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF
OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
