Zaim CLI Client
===

A read-only [Zaim](https://zaim.net/) command line client, written in Golang.

# Usage

```
$ zaim --help
NAME:
   zaim - Zaim CLI Client (readonly)

USAGE:
   zaim [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     auth      Authenticate
     verify    Verify
     money     Money
     category  Category
     genre     Genre
     account   Account
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

`zaim money` sub command has options.

```
$ zaim money --help
NAME:
   zaim money - Money

USAGE:
   zaim money [command options] [arguments...]

OPTIONS:
   --mode value             mode
   --place value            place
   --start_date value       start_date
   --end_date value         end_date
   --category_id value      category_id
   --genre_id value         genre_id
   --from_account_id value  from_account_id
   --to_account_id value    to_account_id
   --account_id value       account_id
```

# Installation

1. `git clone` this repository.

2. `make`

3. Copy `bin/zaim` to your bin directory.


# Configuration

1. Copy `zaim.config.json.sample` to `~/.zaim.config.json`.

2. Edit values of `consumer_key` and `consumer_secret`.