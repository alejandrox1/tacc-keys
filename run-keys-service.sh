#!/usr/bin/env bash
set -e

/usr/sbin/sshd -D &

/app/keys-service

