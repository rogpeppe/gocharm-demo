Gocharm talk
----------

This repository holds the files for a talk given by me (Roger Peppé)
at the Canonical Juju/MAAS sprint in Nuremberg on 2015/04/17.

The charms showing in the talk are inside this repository
and are compilable with gocharm.

To install the demo charm code, run the following commands

	go get launchpad.net/godeps
	go get -d github.com/rogpeppe/gocharm-demo/...
	cd $GOPATH/github.com/rogpeppe/gocharm-demo
	godeps -u *.tsv
	go install ./...

To build the charms, you'll first need to set the $JUJU_REPOSITORY
directory where they will be stored. It will be created if necessary.
For example:

	JUJU_REPOSITORY=$HOME/charms

Then you can build the charms; for example, to build the
last charm in the talk:

	cd $GOPATH/github.com/rogpeppe/gocharm-demo/mycharm-four
	gocharm

To view the presentation, run:
	
	go get golang.org/x/tools/cmd/present
	present

and point your web browser to the link printed by the present tool.

The notes inside the gocharm.slide contain additional information.
