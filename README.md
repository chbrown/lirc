# lirc

Read-only IRC logger / crawler.

* Pronounced "lurk"
* _Not_ [Linux Infrared Remote Control](https://www.google.com/search?q=lirc)

Not a regular [Go](https://golang.org/) user?

    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH

Now you're a regular Go user! Install:

    go get github.com/chbrown/lirc/cmd/lirc

Run:

    lirc -addr irc.freenode.net:6667 -nick InTheEaves go-nuts

Arguments:

* `-help` Show all flags along with their short descriptions and defaults.

* `-addr` IRC server hostname + optional port
  - If missing, the default port is 6697 if `-tls` is specified, or 6667 if not
* `-tls` Connect to `addr` with SSL/TLS
* `-nick` Your desired nickname
* `-user` Your desired username, which some IRC servers require before they'll acknowledge your JOIN requests
* `-real` The realname to accompany the username
* `-out` Output format(s) to write; can be multiple, comma-separated (since they all write to `stdout`, multiple outputs aren't terribly useful yet, but I'm planning to write an InfluxDB output)
  + `raw` Raw IRC protocol with timestamps and direction indicators
  + `text` Colorful text
  + `json` Filtered Protobuf+JSON
* Remaining strings: channels to watch
  - `#` will be prepended to channel names that don't start with a `#` or `&`, to avoid having to quote these arguments, since most shells interpret those as special characters.
  - If you want to join a room that starts with two hashes, like `##linux`, then you have to quote it, e.g. `\##linux` or `'##linux'`.


## References

There are a couple of nice Go IRC clients:

- **GoIRC** [`github.com/fluffle/goirc`](https://github.com/fluffle/goirc)
  + ★ 381 / ⑂ 63
  + Loads of features, opinionated use case.
  + This is pretty close to a real client you could use at the command line.
  + In fact, the [`client.go`](https://github.com/fluffle/goirc/blob/master/client.go) example implements interaction via STDIN.
  + [Documentation](https://godoc.org/github.com/fluffle/goirc/client)
- **irc** [`github.com/sorcix/irc`](https://github.com/sorcix/irc)
  + ★ 171 / ⑂ 13
  + Minimal. Pretty much just splits IRC lines into their component parts.
  + [Documentation](https://godoc.org/github.com/sorcix/irc)

This package, `lirc`, depends on the latter (`github.com/sorcix/irc`).


### IRC Protocol

* [RFC 1459](https://tools.ietf.org/html/rfc1459) (1993): "Internet Relay Chat Protocol"
  - The original specification, which all the others merely "update".
* [RFC 2810](https://tools.ietf.org/html/rfc2810) (2000): "Internet Relay Chat: Architecture"
  - Describes the high-level relationship between servers and clients on an IRC network (which may include more than one server).
* [RFC 2811](https://tools.ietf.org/html/rfc2811) (2000): "Internet Relay Chat: Channel Management"
  - Non-technical channel policies, presumably for channel operators / moderators (no implementation details).
* [RFC 2812](https://tools.ietf.org/html/rfc2812) (2000): "Internet Relay Chat: Client Protocol"
  - This has the most to do with the original, and with the overall current specification.
  - Probably the canonical place for all technical details concerning implementing or using IRC at the moment.
* [RFC 2813](https://tools.ietf.org/html/rfc2813) (2000): "Internet Relay Chat: Server Protocol"
  - Primarily details for developers implementing an IRC server, or IRC administrators.
* [RFC 7194](https://tools.ietf.org/html/rfc7194) (2014): "Default Port for Internet Relay Chat (IRC) via TLS/SSL"
  - Formalizes the common usage of TCP port `6697` as the standard port for incoming IRC connections over TLS/SSL.
  - Reminder of the common usage of TCP port `6667` as the standard port for incoming IRC connections over plain text.
    + This port was used in several examples in the preceding RFCs, but never explicitly called out as the convention.
  - IANA specifies port TCP/UDP port `194` as the standard port for plain text connections and `994` for encrypted connections, but these both require root access on the server, which is not always available, or desired.


## License

Copyright © 2016 Christopher Brown. [MIT Licensed](https://chbrown.github.io/licenses/MIT/#2016).
