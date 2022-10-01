# JunOS Configuration Archival

Based [Backup Configurations to an Archive Site][backup-configuration]

[backup-configuration]: https://www.juniper.net/documentation/en_US/junos/topics/task/configuration/junos-software-system-management-router-configuration-archiving.html

## Local side

```bash
ftp-worker \
  -on-write ./on-write-junos-archive.py \
  -workdir path/to/git-repo
```

## JunOS side

Based `transfer-on-commit` command automate backup JunOS configuration to remote-side

```plain
set system archival configuration transfer-on-commit
set system archival configuration archive-sites ftp://<username>@<host>:<port>/<path> password <password>
```
