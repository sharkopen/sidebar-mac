# gosidebar
gosidebar is a small utility on a Mac machine that allows you to add any folder to the sidebar
modify sidebar for macos

## How to build?
go build -ldflags="-s -w" -o gosidebar

## How to use?
Usage:
gosidebar [add|rm|list] /path/to/folder

1. list bar

    ./gosidebar list


2. add sidebar

    ./gosidebar add /path/to/folder


3. remove sidebar

    ./gosidebar rm /path/to/folder

License
-------

This is free software under the terms of the MIT license (check the
[COPYING file](/COPYING) included in this package).
