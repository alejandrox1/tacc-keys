#!/usr/bin/env bash

check_hostname_is_worker() {
    hostname=$1
    output=$(echo $hostname | grep "cluster_worker")
    test ! -z $output
}

get_hostname() {
    IP=$1
    nslookup $IP | grep name | awk '{print $NF}' | sed 's/\.$//'
}

get_tcp_established_connections() {
    if command -v ss > /dev/null 2>&1; then
        IPS=$(ss -t | grep ESTAB | awk '{print $NF}' | cut -d: -f1)
    elif command -v netstat > /dev/null 2>&1; then
        IPS=$(netstat -tn | grep ESTABLISHED | awk '{print $5}' | cut -d: -f1)
    else
        printf "Neither ss nor netstat could be found\n" >&2
        exit 1
    fi

    for ip in $IPS; do
        hostname=$(get_hostname $ip);
        if check_hostname_is_worker $hostname; then
            printf "%s %s\n" $ip $hostname
        fi
    done
}

get_tcp_established_connections
