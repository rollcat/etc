# prompter

Opinionated `$PS1` generator for ~~your~~ my shell.

## How it looks like

The layout and features are *not* configurable. When there's nothing
else to report, the prompt looks like this:

```
: user@host ~
$
```

When there's something interesting to report about the environment
(such as the current git branch), it looks like that:

```
: user@host [git=master] ~/src/project
$
```

## Features

- Colors. Non-root users are highlighted in green, root is red;
  hostname in a local session is green, remote sessions are red.
- `git`: current git branch.
- `nix`: whether we are in a Nix shell.
- `venv`: current Python virtualenv (if activated).
- `docker`: current Docker context (if different from default).

## Compatibility

Tested with:

- [Zsh](https://www.zsh.org),
  [GNU Bash](https://www.gnu.org/software/bash/),
  MirBSD [mksh](http://www.mirbsd.org),
  OpenBSD [ksh](http://man.openbsd.org/ksh).
- macOS, Linux, OpenBSD.
- 64-bit Intel and ARM architectures.

## Installation

Download the release appropriate for your OS and architecture, and put
the executable somewhere in your `$PATH`.

Set your `$PS1` as follows:

### Bash, mksh, ksh

```sh
PS1='$(command prompter)'
```

### Zsh

```zsh
precmd() {
    PS1="$(command prompter)"
}
```

### Plan 9 rc(1)

Sorry.

## FAQ

### I like the prompt but your code sucks

Use this instead:

```sh
#!/bin/sh
set -eu
exec printf ': %s@%s %s\n$ ' "${USER:-$(id -nu)}" "${HOSTNAME:-$(hostname)}" "${PWD:-$(pwd)}"
```
