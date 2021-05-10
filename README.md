# TLDR

[TL;DR pages](https://github.com/tldr-pages/tldr/) in Go.

## How to install
With Go, run:
> go get github.com/Pauloo27/tldr

And then create the config at `~/.config/tldr.env` with the following content:
```env
# bs, da, de, es, fr, hbs, hi, id, it, ja, ko, ml, nl, no, pl, pt_BR, pt_PT, ru, sv, ta, th, tr, zh, zh_TW
TLDR_LANG=""
# bat, cat, less, vim, mdcat
TLDR_VIEWER="/usr/bin/less"
```

_You also need to add the Go binary folder to your path, if you haven't done it yet._
_If you don't know where the Go binary folder is, run `go env GOPATH`._

_You can set a language, leave it empty to use English._

_I recommend using mdcat as the viewer_

## Compiling
Clone the project:
> $ git clone https://github.com/Pauloo27/tldr.git && cd tldr

Install it:
> $ make install

## Viewers
Recommended viewer: [catmd](https://www.archlinux.org/packages/community/x86_64/mdcat/).

Also: `cat`, `bat`, `less`.

## Info
The tldr pages repository will be cloned to `~/.cache/tldr`. 
Once cloned, run `tldr --update` to update the pages.

## License

<img src="https://i.imgur.com/AuQQfiB.png" alt="GPL Logo" height="100px" />

This project is licensed under [GNU General Public License v2.0](./LICENSE).

This program is free software; you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

