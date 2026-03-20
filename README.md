Fastbound Downloader
====================
A docker container that will authenticate to your [Fastbound](https://www.fastbound.com/) account and perform the mandatory daily bound book download as a PDF
in order to maintain compliance with [ATF record keeping requirements](https://www.fastbound.com/atf-2016-1-compliant/.

Settings File
-------------
This file must be mode `0400` and the path must be `/config/settings.json` unless overridden by a CLI arg.

Currently all of these values are required:
```json
{
  "fastbound": {
    "account-number": "123ABC1234",
    "api-key": "123ABC1234",
    "audit-user": "pgibbons@initech.com"
  },
  "paths": {
      "bound-books": "/books/"
  },
  "is-cron": false,
  "disable-metrics": false,
  "metrics-port": "9090",
  "scanning-interval": 1440
}
```

**Optional Variables:**

1. `is-cron` (Default: false) will disable the cycle logic. The downloads will execute once and exit. This is useful if you want to run this as a cron in K8s or elsewhere. This also disables metrics.
2. `disable-metrics` (Default: false) will disable the Prometheus `/metrics` endpoint on the container.
3. `metrics-port` (Default: 9090) lets you override the default port.
4. `scanning-interval` (Default: 1440) how often, in minutes fbdownloader should check for new files to download

Functionality
-------------
This tool loops on a 24-hour cycle from the time the container starts. Each interval will result in a download of the specified Fastbound account's
A&D book to the specified path. This should be a volume mount of some kind as ephemeral data defeats the purpose of this process.

Dependencies
------------
These are the direct dependencies fetched with `go get` inside [go.mod](go.mod)
