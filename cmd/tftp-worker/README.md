# TFTP Worker

The [TFTP](https://www.rfc-editor.org/rfc/rfc1350) service

redirect upload (write operation) and download (read operation) as local program calling

## Usage

<!-- markdownlint-disable -->

```plain
Usage of tftp-worker:
  -addr string
    	TFTP listen address (default ":tftp")
  -on-read string
    	Read operation local hook program
  -on-write string
    	Write operation local hook program
  -workdir string
    	Work directory
```

<!-- markdownlint-restore -->

### Read hook

Equivalent to:

```plain
read-hook-program <filename> | <file-content>
```

Environments:

| Field Name       | Type      |
| ---------------- | --------- |
| `TFTP_ACTION`    | `READ`    |
| `TFTP_READ_FILE` | file path |

### Write hook

Equivalent to:

```plain
<file-content> | write-hook-program <filename>
```

Environments:

| Field Name        | Type      |
| ----------------- | --------- |
| `TFTP_ACTION`     | `WRITE`   |
| `TFTP_WRITE_FILE` | file path |
