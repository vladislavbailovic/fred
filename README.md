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
- [Usage](#usage)


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


## Config file

The feeds will be fetched from the list in config file. At the moment, the config file is essentially just a newline-separated list of URLs. See [testdata/fredrc](testdata/fredrc) for an example.


## Usage

Aside from limiting output to `N` number of days, it is also possible to scope the displayed topics to a subset of categories using `-t`/`--topic` CLI flag. Multiple topics can be added, and the articles tagged with any of the provided topics will be listed.

Note: articles that don't have any categories in the feed will have topics assigned by parsing the article title.

Tip: the topics filter will be applied _after_ the topics have been processed into their final string form - meaning, a more advanced matching is possible by including the separator (`:`). This is handy for short topics (e.g. `:go:`).