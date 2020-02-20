# œfi

Get train/bus departures for public transport (öffentlicher Personalnahverkehr) in Germany.

[![asciicast](https://asciinema.org/a/Q6reoxtFK3vvUlCyeYVJvCUdl.svg)](https://asciinema.org/a/Q6reoxtFK3vvUlCyeYVJvCUdl)

## Installation

- Install [go](https://golang.org/doc/install)
- Run `go get github.com/nfode/oefi`

## Command completion

### Bash

Bash is currently not implemented 

### ZSH

To use completion run:
```
source <(oefi completion zsh)
```
This works until you end the session.

If you want to load completions for each sessions run:

```
echo "source <(oefi completion zsh)" >> .zshrc
```

or add `source <(oefi completion zsh)` manually to your `.zshrc`.

