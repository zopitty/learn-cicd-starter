# Linting

Linting checks

```
go install honnef.co/go/tools/cmd/staticcheck@latest
```

Running the check

```
staticcheck ./...
```

If 'command not found' 

```
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```

test if it's working with an unused function

```
func unused() {
    // this function does nothing
    // and is called nowhere
}
```
