# htmlr = HTML resolver
Is a light weight cmd tool to resolve html templates into one files by using `script` tag with `type="text/x-template"` and `{% include "template.html" %}`.

## Install:
run with:
```
go install https://github.com/artvel/htmlr/htmlr
```
## Build with:
```
go build htmlr/htmlr.go
```
## Usage:
```sh
htmlr template.html -o output.html
```