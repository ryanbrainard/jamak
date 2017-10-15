# jamak (자막)

Parses subtitles for import into other applications.

## Installation

First install [Go](https://golang.org/doc/install) and then install `jamak`:

    $ go get github.com/ryanbrainard/jamak/cmd/jamak

## Usage

`jamak` is a small, sharp UNIX-like tool. As such, it only reads from stdin and writes to stdout, so it can be used in a pipeline. Use a command like `cat` or `pbpaste` to input a file or the clipboard, and then use redirection (i.e. `>`) or `pbpaste` to output back to a file or the clipboard.

This application also includes a [web interface](#web-interface) with a simplified set of options.

#### File Based

For example, if there is a file named `input.srt` that looks like this:

```
1
00:00:02,710 --> 00:00:04,410
예전에 영화에서 봤는데

2
00:00:06,923 --> 00:00:08,453
이렇게 앉아있으니까
```

It can be processed and written to `output.json` like this:

```sh
$ cat input.srt | jamak --formatter readlang > output.json
```

The resulting file will be in tab-separated format:

```tsv
{
  "audioMap": [
    {
      "t": 2.71,
      "w": 0
    },
    {
      "t": 6.923,
      "w": 3
    }
  ]
}```

The output can then be imported into your favorite flashcard app.

#### Clipboard Based

Alternatively, if you'd rather work with just data from/to the clipboard, use `pbpaste` and `pbcopy` on MacOS:

```sh
$ pbpaste | jamak | pbcopy
```

# Options

The parser and formatter can be set for different inputs and outputs.

## Parsers

Options for `--parser` flag:

 - `srt`: (default) SRT format.

## Formatters

Options for `--formatter` flag:

 - `srt`: (default) SRT format
 - `readlang`: Readlang audiomap format
 
# Web Interface

This application also has a web interface with a simplified set of options, which can easily be run locally or deployed to Heroku.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/ryanbrainard/jamak
$ cd $GOPATH/src/github.com/ryanbrainard/jamak
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

# Disclaimer

For personal use only. Do not use this tool for publishing copyrighted content. Respect copyright holders' rights of their content. Logos shown above are copyrighted by their respective owners.
