# ProxyFor

This repo contains a robust, fast, cross-platform proxy checker written in golang.
Currently, it supports checking on HTTP(s) / Socks5 proxies.

Unlike most proxy checkers, this tool doesn't care about the anonymity level and focused on working against your specific target and your desired response code. (e.g., 200 )

## Usage

```
usage: ProxyChecker [-h|--help] [-t|--threads <integer>] [-r|--response
                    <integer>] -p|--path "<value>" [-d|--domain "<value>"]
                    [-T|--timeout <integer>]

                    Checks for valid proxies and write valid ones in file

Arguments:

  -h  --help      Print help information
  -t  --threads   Number of threads. Default: 40
  -r  --response  expected HTTP response code. Default: 200
  -p  --path      path to proxy.txt (required)
  -d  --domain    Domain to check proxies against it. Default:
                  https://httpbin.org/ip
  -T  --timeout   timeout in seconds. Default: 10

```

### Example

```
ProxyFor -p proxylist.txt -d https://target.com
```

## Motivation

Wrote this to aid [CloudBrute]("https://github.com/0xsha/cloudbrute")
This tool is useful for finding proxies to bypass rate and region limitation. Using this tool, you can find proxies that work on your target. Some everyday use cases are scraping, brute force, bucket hunting, etc.

## ProxyFor in action

[![asciicast](https://asciinema.org/a/gK5bVyEE9CYBLP8fYoy3IavFF.svg)](https://asciinema.org/a/gK5bVyEE9CYBLP8fYoy3IavFF)

## How to contribute

- Add a module or fix something and then pull request.
- Share it with whomever you believe can use it.
- Do the extra work and share your findings with community &hearts;
