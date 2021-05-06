#.PHONY: all
#FORCE: ;

test: test-no-cache
	#go test -tags testing ./... -v
	go test -v ./...

test-no-cache:
	go clean -testcache 	

test-convey:
	~/go/bin/goconvey

version:
# gets tags across all branches, not just the current branch
	git describe --tags `git rev-list --tags --max-count=1` 

# new-version:
# 	read -p "Enter the version description: "  description
# 	echo ${description}
