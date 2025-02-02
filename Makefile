
test:
	go test ./... -run=.*/more-empty-lines-in-b-than-a-00
	#go test -v ./...

run:
	go run ./godiff/diff.go

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/godiff \
		-gcflags -m=2 \
		./godiff/ 

