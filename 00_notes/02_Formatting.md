# Formatting

make use of go fmt to auto-format all files

```go
go fmt ./...
```

go fmt prints the names all all the files it fixes

use test to make go fmt have an exit status code of 1 if ther eare files to be formatted

```
test -z $(go fmt ./...)
```

$() for command substitution. runs the command inside then replaces $() with output

test will check for empty string, 0 if it is, 1 if it isn't 

echo $? to print the exit code

```
test -z $(go fmt ./...)
echo $?
```

add into ci.yml

```
test -z $(go fmt ./...)
```
