# noteworthy
Command line user interface for reading RSS feeds written in Golang

## Background
A personal experiment in Go development and programming for the cli.

## TODOs
- [ ] convert to cui format (move around, launch browser, some pretty styling)
- [ ] fetch feed from ze interwebz (instead of good old file)
- [ ] file with a list of feeds (urls basically that are then fetched, decoded, and listed)
- [ ] Add a heirarchy of Top(options, manage(add/remove feed), and feeds)/Feeds(based on input)/Items(fetched per feed)
- [ ] possibly convert to client/server style with a DB in the bg that keeps track of read/unread, and managed feeds (futuristic)
- [ ] Serve as a client towards other RSS servers (e.g. owncloud)
- [ ] Configuration file to determine the usage (statically configured feeds, connecting to RSS server, or both)
- [ ] Multiple layouts; vertical & horizontal, and configurable/hotkey controlled
- [ ] Expand to other feed formats; Atom, ???
