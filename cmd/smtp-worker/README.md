# SMTP worker

The SMTP worker

Redirect SMTP request to local hook program.

## Usage

<!-- markdownlint-disable -->

```plain
Usage of smtp-worker:
  -addr string
    	SMTP listen address
  -on-request string
    	SMTP request handler
  -tls-cert-file string
    	TLS Certificate file
  -tls-key-file string
    	TLS Key file
  -username string
    	Username (default "user")
  -password string
    	Password (default "pass")
```

<!-- markdownlint-restore -->

### On Request

Equivalent to:

```plain
<body> | hook-program <to-address> <from-address>
```

Environments:

| Field Name            | Type                |
| --------------------- | ------------------- |
| `SMTP_FROM`           | string              |
| `SMTP_FROM_USERNAME`  | string              |
| `SMTP_FROM_DOMAIN`    | string (fqdn)       |
| `SMTP_TO`             | string              |
| `SMTP_TO_USERNAME`    | string              |
| `SMTP_TO_DOMAIN`      | string (fqdn)       |
| `SMTP_TO_DOMAIN_TYPE` | string              |
| `SMTP_HOSTNAME`       | string              |
| `SMTP_LOCAL_ADDR`     | string (ip address) |
| `SMTP_REMOTE_ADDR`    | string (ip address) |
| `SMTP_UTF8`           | boolean             |
| `SMTP_REQUIRE_TLS`    | boolean             |
| `SMTP_BODY_TYPE`      | string              |
| `SMTP_BODY_SIZE`      | interge             |

`SMTP_TO_DOMAIN_TYPE`:

- `DISPOSABLE_MAIL`, disposable email provides
- `FREE_MAIL`, free email provides
- `SWOT_MAIL`, education email domains
- `DDNS`, dynamic dns provides

`SMTP_BODY_TYPE`:

- `7BIT`, 7bit, e.g: UTF-7, Base64 or Quoted-Printable Encoding
- `8BITMIME`, 8bit MIME
- `BINARYMIME`, Binary MIME
