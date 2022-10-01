#!/bin/bash
set -xeuo pipefail
curl -fsSL "https://github.com/eusonlito/disposable-email-validator/raw/master/data/domains.txt" |
	gzip -9 >disposable.txt.gz
curl -fsSL "https://github.com/willwhite/freemail/raw/master/data/free.txt" |
	gzip -9 >free.txt.gz
curl -fsSL "https://github.com/JetBrains/swot/releases/download/latest/swot.txt" |
	grep -v '^-' |
	gzip -9 >swot.txt.gz
curl -fSSL "https://github.com/neu5ron/dynamic_dns_lists/raw/main/dynamic-dns.txt" |
	gzip -9 >ddns.txt.gz
