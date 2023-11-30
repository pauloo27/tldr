# TLDR

Unofficial [TL;DR pages](https://github.com/tldr-pages/tldr/) in Go.

_not fully compatible with the specification_

## How to install

### With Go, run:

```bash
go install github.com/pauloo27/tldr/cmd/tldr@latest
```

_You also need to add the Go binary folder to your path, if you haven't done it yet._
_If you don't know where the Go binary folder is, run `go env GOPATH`._

### Manually compiling

Clone the project:
```bash
git clone https://github.com/Pauloo27/tldr.git && cd tldr
```

Install it:
```sudo
sudo make install
```

## Config

Create a config at `~/.config/tldr/config.toml`, example:

```toml
viewer="less"
language="en"
```

_I recommend using bat or mdcat as the viewer_

## Info
The tldr pages repository will be cloned to `~/.cache/tldr-repo`. 
Once cloned, run `tldr -u` to update the pages.

## License

This project is licensed under the [MIT license](./LICENSE).

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
