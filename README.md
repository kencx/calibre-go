# calibre-go

Proof of concept API library for Calibre and `calibredb` with Go and
[sqlboiler](https://github.com/volatiletech/sqlboiler).

## Regenerate Models

1. Install `sqlboiler` and `sqlboiler-sqlite3` driver:

```bash
$ go install github.com/volatiletech/sqlboiler/v4@latest
$ go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
```

2. Generate the `models` package:

```bash
# ensure ~/go/bin is in PATH
$ export PATH=$PATH:$HOME/go/bin

$ sqlboiler --add-global-variants -c sqlboiler.toml sqlite3
```
