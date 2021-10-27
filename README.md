# `clone`

Clone is a small tool that wraps `git` for cloning repositories into a `(host)/owner/repo` folder structure.

This is handy for saving some key strokes if this is how you prefer to organize separate repositories to quickly navigate later.

Initially this was a bash script but I rewrote it in go to make it easier to hack on as I need to.

```
$ ./clone --help
Clone remote git repositories into a host/owner/repo file structure relative to where this command is run

Usage:
  clone 'http(s) or git@ URL' [flags]

Flags:
  -d, --dry-run        show information but do not actually clone repository
  -h, --help           help for clone
  -i, --include-host   include host in cloned folder structure (default true)
```

### ssh authentication not working

Due to a bug in `go-git` you'll likely need to run the following:
```
ssh-add ~/.ssh/id_rsa
```
