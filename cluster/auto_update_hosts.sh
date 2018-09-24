#!/usr/bin/env bash
#
# auto_update_hosts.
#
# Pass as argument the system's host file. Use this value to update the hosts
# obtained from `get_host`.
#

hosts=$(get_hosts)
echo "$hosts" > "$1"

# Make a copy of the hosts file.
cp /etc/hosts /tmp/hosts

while sleep 5; do
    current_hosts=$(get_hosts)
    [ "$hosts" != "$current_hosts" ] && echo "$current_hosts" > "$1"
    hosts=$current_hosts

    # If there are no differences between the hostfile and the original
    # (/tmp/hosts), then try and update the hosts.
    cp /tmp/hosts /etc/hosts
    update_hosts >> /etc/hosts
done
