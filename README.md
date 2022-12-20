# fred

`fred` is a console RSS reader, made for my specific purposes of consuming feeds in Vim, using (franken-)VimWiki flavored markdown.

Example usage string:

```console
$ fred -d 2 -r
```

This will fetch feeds from the [config file](#config-file), output last 2 days' worth of articles in a temp file and open it in Vim. From there, I can easily yank and paste pre-formatted links with tags/descriptions into my daily bookmarks diary.


## Table of Contents

- [Quick Start](#quick-start)
	- [Building](#building)
	- [Running](#running)
	- [Testing](#testing)
	- [Config file](#config-file)


## Quick Start


### Building

```console
$ go build -o ./ fred/cmd/...
```


### Running

```console
$ go run fred/cmd/fred
```


### Testing

```console
$ go test ./...
```


### Config file

The feeds will be fetched from the list in config file. At the moment, the config file is essentially just a newline-separated list of URLs. See [testdata/fredrc](blob/main/testdata/fredrc) for an example.