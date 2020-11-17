build:
		go build

install: build
	cp ./.env ~/.config/tldr.env
	sudo cp ./tldr /usr/bin/

