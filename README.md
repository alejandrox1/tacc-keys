# Key Service


To build and run the container, which has a sshd service running along with the
`tacc-keys` executable do
```
make run
```

If you want to ssh inside the container you can do so by running the following
command:
```
make ssh
```


For some simple benchmarks comparing the two ways to login:
```
perf stat -r 10 -d make ssh-norm-cmd
```

and 
```
perf stat -r 10 -d make ssh-cmd
```


Note, to kill the container running the ssh server you can try (if it is the
only container running)
```
docker rm -f $(ps -q)
```

## Contents

* [agave-config](agave-config) contains a keys-client apt for use by sshd as an
`authorizedKeysCommand` when the keys service is locked behind authentication
through agave.

* [auhtorized keys command](authorizedkeycommand) contains what the actual
program to be used by sshd when using `AuthorizedKeysCommand`.

* [generate-validate-keys](generate-validate-keys) contains a program that
shows how to create public and private keys in Go.

* [keys client](keys-client) is a full client for the keys service. It has the
ability to create ssh keys, delete public keys from the keys service, and list
all the public keys registered to a certain user.

* [ssh-server](ssh-server) builds an ssh server inside a container.
