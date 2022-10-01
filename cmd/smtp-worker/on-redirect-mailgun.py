#!/usr/bin/env python3
import email.parser
import os
import sys
from smtplib import SMTP_SSL


def main(to_addr: str, from_addr: str):
    # not allowed use disposable mail (one-time mail)
    if os.getenv('SMTP_TO_DOMAIN_TYPE') == 'DISPOSABLE_MAIL':
        sys.exit(1)
    # paring email payload
    message = email.parser.Parser().parse(sys.stdin)
    # make smtp client
    client = SMTP_SSL('smtp.mailgun.org', port=465)
    # login
    client.login('postmaster@YOUR_DOMAIN_NAME', 'YOUR_PASSWORD')
    # send message
    client.send_message(message, from_addr, to_addr)
    # quit
    client.quit()


if __name__ == "__main__":
    main(sys.argv[1], sys.argv[2])
