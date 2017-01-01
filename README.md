# lirc

Read-only IRC logger / crawler.

* Pronounced "lurk."
* _Not_ [Linux Infrared Remote Control](https://www.google.com/search?q=lirc).

Not a regular [Go](https://golang.org/) user?

    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH

Now you're a regular Go user! Install:

    go get github.com/chbrown/lirc

Run:

    lirc -addr irc.freenode.net:6667 -nick InTheEaves go-nuts

Arguments:

* `-addr` IRC server hostname
* `-nick` Your desired nickname
* Remaining strings: channels to watch
  - `#` will be prepended to channel names that start without a `#`, to avoid having to quote these arguments, since most shells consider `#` a comment indicator.
    But if you want to join, for example, `##linux`, just use `\##linux` or `'##linux'`.


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


## License

Copyright © 2016 Christopher Brown. [MIT Licensed](https://chbrown.github.io/licenses/MIT/#2016).
