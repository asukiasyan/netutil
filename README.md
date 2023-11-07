# netutil

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/asukiasyan/netutil?include_prereleases)
![GitHub last commit](https://img.shields.io/github/last-commit/asukiasyan/netutil)
![GitHub issues](https://img.shields.io/github/issues-raw/asukiasyan/netutil)
![GitHub pull requests](https://img.shields.io/github/issues-pr/asukiasyan/netutil)

### Network utility to lookup various records for given host

# Features

* Lookup NameServer for given host
* Lookup the IP addresses for given host
* Lookup CNAME records for given host
* Lookup MX records for given host


# Installation
```
git clone https://github.com/asukiasyan/netutil.git
cd netutil
go build .

```

# Usage


```
netutil command [url...]


// Lookup NameServer for given host
netutil ns <URL>

//Lookup the IP addresses for given host
netutil ip <URL>

//Lookup CNAME records for given host
netutil cname <URL>

//Lookup MX records for given host
netutil mx <URL>

```

# Known issues

This repository is a work in progress, I am trying to have on main somewhat usefull version, but it could have bugs.

* [x] refactor code to accept only valid args

* [ ] Convert case statement to individual functions

* [ ] Add tests for each case
