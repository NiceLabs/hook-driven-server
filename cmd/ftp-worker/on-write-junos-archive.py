#!/usr/bin/env python3
import gzip
import os
import re
import sys
from subprocess import call

# pattern: <router-name>_YYYYMMDD_HHMMSS_juniper.conf<.n>.gz
re_filename = re.compile(r'^(\S+)?_\d{8}_\d{6}_juniper\.conf(?:\.\d+)?\.gz$')


def main(filename: str):
    matched = re_filename.match(os.path.basename(filename))
    if not matched:
        sys.exit(1)
    saved_filename = '%s.conf' % (matched.group(1) or 'unnamed')
    decompressed = gzip.decompress(sys.stdin.buffer.read()).decode()
    if not has_changed(saved_filename, decompressed):
        return
    with open(os.path.join(saved_filename), 'w') as fp:
        fp.write(decompressed)
    call(['git', 'add', saved_filename])
    call(['git', 'commit', '-m', filename])
    call(['git', 'push'])


def has_changed(filename: str, decompressed: str):
    if not os.path.exists(filename):
        return True
    with open(filename, 'r') as fp:
        original = fp.read()
    return original.splitlines()[1:] != decompressed.splitlines()[1:]


if __name__ == '__main__':
    main(sys.argv[1])
