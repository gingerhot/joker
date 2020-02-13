# Joker - A commandline tool helps to check your dev environment

Joker is a commandline tool used to check your dev environment. If you need a lot of tools and services to serve your development, you'll miss some operation before coding and testing. You may spend some time to debug an odd error and then you find that it's just because you forget starting a service or exporting a PATH. That's aweful!

Joker provides several kinds of checking as:

* ports (is open or not?)
* env variables (is matched or not with the given expected value?)
* file paths (exists or not?)
* commands (is available or not?)
* command outputs (does command output match with the given expected value?)

## Installation

```sehll
go get github.com/gingerhot/joker
```

Note: macOS and Linux supported, not Windows at present.

## Usage

You can check `example.yaml` to get how to config to check your dev env. It's straight forward and easy to play.

When you have such a yaml-format config file, then run:

```shell
joker example.yaml
```

if you named your config file as exact `joker.yaml`, just run `joker` and it will work perfectly.

Joker is handy, enjoy!
