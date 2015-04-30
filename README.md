# c2v - Colorized Code Viewer

`c2v` is simple code viewer supports syntax highlighting, some encodings and multiple files.

## Description

***DEMO:***

![demo](url)

When the code reading, if it is not syntax highlighting, we feel very painful. Therefore we use the text editor instead of `cat` and `less` when you see quickly the source code. Such case, it is preferable to use the pygments. This tool improves the usability of `pygmentize` which is a Command Line Interface of pygments.

## Requirement

- python2 or later
- [Pygments](http://pygments.org)
	- `pip install Pygments`

## Features

- Supports 300 or more colors syntax highlighting
- Supports many encodings (e.g., iso-2022-jp, utf-8, ucs-2, euc-jp, cp932)
- Supports multiple files
- Fast thanks to the *parallel processing*
- Written in Golang

## Usage

Basically,

```bash
$ c2v file1 file2 ...
```

To specify the style (color scheme):

```bash
$ c2v -s molokai file
```

If you specify an invalid style, the "default" style is outfitted. Whether style that you want to specify is valid, you can list up with the following command.

```bash
$ c2v -l               # list all styles
style1
style2
...
$ c2v -l some_style    # check if available
$ echo $?
0
```

For more information of the usage, see the following help `c2v --help`.

## Installation

1. Download from [here]().
2. Install to the directory in your `$PATH`.

How to install for developers is:

```bash
$ go get github.com/b4b4r07/c2v
```


## Todo

- Supports `stdin`

## Author

BABAROT (a.k.a. [b4b4r07](https://b4b4r07.com))
 
## License

MIT