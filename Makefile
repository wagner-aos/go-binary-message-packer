.PHONY: all
FORCE: ;

test: test-no-cache
	#go test -tags testing ./... -v
	go test -v ./...

test-no-cache:
	go clean -testcache 	

test-convey:
	~/go/bin/goconvey