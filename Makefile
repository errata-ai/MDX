.PHONY: test zip sync

all: test

zip:
	zip -r MDX.zip MDX -x "*.DS_Store"

sync:
	cd testdata && vale sync && vale transform test.mdx > out.md && cd -

test: zip sync
	go run main.go
