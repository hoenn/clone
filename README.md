# `clone`

Clone is a small tool that wraps `git` for cloning repositories into a `(host)/owner/repo` folder structure. For example:

```
github.com/
    hoenn/
        clone/
        go-hn/
        ynab-metrics/
    prometheus/
        node_exporter/
        prometheus/
gitlab.com/
    hoennn/
        dotfiles/
```

This is handy for saving some key strokes if this is how you prefer to organize separate repositories to quickly navigate later.

Initially this was a bash script but I rewrote it in go to make it easier to hack on as I need to.

```
$ ./clone --help
Clone remote git repositories into a host/owner/repo file structure relative to where this command is run

Usage:
  clone 'http(s) or git@ URL' [flags]

Flags:
  -d, --dry-run        show information but do not actually clone repository.
  -p, --progress       show progress of git to stdout (default true).
  -h, --help           help for clone.
  -i, --include-host   include host in cloned folder structure (default false).
```
