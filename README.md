# Simple Toml Name Service
[![Build Status](https://travis-ci.org/pyama86/STNS.svg?branch=master)](https://travis-ci.org/pyama86/STNS)

simple toml name service is Linux `/etc/passwd`,`/etc/group`,`/etc/shadow` name resolution from toml format config

client library:https://github.com/pyama86/libnss_stns

## install
## redhat/centos
```
$ curl -fsSL https://repo.stns.jp/scripts/yum-repo.sh | sh
$ yum install stns
```
## debian/ubuntu
```
$ curl -fsSL https://repo.stns.jp/scripts/apt-repo.sh | sh
$ apt-get install stns
```

## config
* /etc/stns/stns.conf
```
port = 1104
[users.example]
id = 1001
group_id = 1001
directory = "/home/example"(default:/home/:user_name)
shell = "/bin/bash"(default:/bin/bash)
keys = ["ssh-rsa XXXXX…"]
[groups.example]
id = 1001
users = ["example"]
```
support format /etc/passwd,/etc/groups,/etc/shadow

## author
* pyama86
