# FTP Worker

The [FTP](https://www.rfc-editor.org/rfc/rfc959) worker

Redirect upload (write operation) and download (read operation) as local program calling

## Usage

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
