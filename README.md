# macos-brightness

Your friendly neighbordhood macOS brightness CLI. Sets the brightness of the builtin display (e.g. on MacBooks).

Tested on macOS Mojave (10.14), should work on High Sierra (10.13) as well.

## Installing

```console
$ go get -u github.com/mafredri/macos-brightness/cmd/brightness
```

## Usage

### Get current brightness

```console
$ brightness
75.66666603088379
```

### Set brightness percentage

```console
$ brightness -b 90.5
```

## See also

- [macos-darkmode](https://github.com/mafredri/macos-darkmode): Simple CLI for controlling macOS Dark Mode via Apples private API
