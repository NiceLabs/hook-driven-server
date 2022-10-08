# FTP Worker

The [FTP](https://www.rfc-editor.org/rfc/rfc959) worker

Redirect upload (write operation) and download (read operation) as local program calling

## Usage

<!-- markdownlint-disable -->

```plain
Usage of ftp-worker:
  -hostname string
    	FTP listen hostname
  -port int
    	FTP listen port
  -on-read string
    	Read operation local hook program
  -on-write string
    	Write operation local hook program
  -workdir string
    	Work directory
  -username string
    	Username (default "user")
  -password string
    	Password (default "pass")
  -tls-explicit
    	Explicit FTPS
  -tls-cert-file string
    	TLS Certificate file
  -tls-key-file string
    	TLS Key file
```

<!-- markdownlint-restore -->

### Read hook

Equivalent to:

```plain
read-hook-program <filename> | <file-content>
```

Environions:

| Field Name        | Type      |
| ----------------- | --------- |
| `FTP_ACTION`      | `READ`    |
| `FTP_PATH`        | file path |
| `FTP_READ_OFFSET` | integer   |

### Write hook

Equivalent to:

```plain
<file-content> | write-hook-program <filename>
```

Environments:

| Field Name   | Type      |
| ------------ | --------- |
| `FTP_ACTION` | `WRITE`   |
| `FTP_PATH`   | file path |

## Example

- [JunOS Configuration Archival](EXAMPLE-JUNOS.md)
