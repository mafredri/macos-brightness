# macos-brightness

Your friendly neighbordhood macOS brightness cli. Sets the brightness of the builtin display(e.g. on MacBooks).

Tested on macOS Mojave (10.14), should work on High Sierra (10.13) as well.

## Installing

```
go get -u github.com/mafredri/macos-brightness/cmd/brightness
```

## Usage

### Get current brightness

```
$ brightness
75.66666603088379
```

### Set brightness percentage

```
$ brightness -b 90.5
```
