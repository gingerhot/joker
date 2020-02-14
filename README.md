# Joker - Helps to check your dev environment

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

## Q&A

1. Why I name it Joker?

Not special but I just watched the movie [Joker](https://www.imdb.com/title/tt7286456/), and it's more fun than to name it as `dev-env-checker`.

2. Why not add a `match` directive as `expected` in outputs check?

This can be achieved by some shell tricks, for example:

```yaml
outputs:
  - name: Check Go version
    cmd: go version | grep -q 'go1.13.6' && echo 'matched'
    expected: matched
```

And hope this example can enlighten your tremendous creativity.

## Contributing

Bug reports and pull requests are welcome. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The software is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
