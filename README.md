# gocalc

## Installation

```shell
go install github.com/mabego/gocalc@latest
```

## Usage

### REPL

```
$ gocalc
calc> (2.5 - 1.35) * 2.0
2.3
calc> -sin((-1+2.5)*pi)
1
calc> 180*atan2(log(e), log10(10))/pi
45
calc> help

Commands:

help    - Show this help
clear   - Clear the screen
exit    - Exit the calculator

calc> exit
```

### Command mode

```
$ gocalc "(2.5 - 1.35) * 2.0"
2.3
$ gocalc "-sin((-1+2.5)*pi)"
1
$ gocalc "180*atan2(log(e), log10(10))/pi"
45
```

## Acknowledgements

This project is a fork of [mnogu/go-calculator](https://github.com/mnogu/go-calculator).

## Updates

* Command mode
* Prompt autocomplete suggestions
* Refactor for CLI use only