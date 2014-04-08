package main

// Is this right? 

const usage =
`usage: winmux [-r [acme window id | acme window title]] [ cmd ]

winmux is a terminal multplexer for Plan9.

The simplest invocation is

	winmux

which will go hunting for a running instance of winmux and tell it to
connect to a running acme, find a reasonable window there and connect
itself to that window. This invocation falls back in several ways.

If there is not a running winmux (there can be only one winmux per
namespace), then this instance of winmux registers itself in the
current namespace, forks the default shell (typically $PLAN9/bin/rc)
and finds an acme window and connects to it.

If there is not an acme window that could be associated with winmux,
winmux will connect to acme and make one.

The next simplest invocation is

	winmux <cmd>

which will make a new acme window connected to cmd.
`


