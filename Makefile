APP=discoupload

run: build
	./$(APP)

build:
	go build

debs:
	go get ./...

update-debs:
	go get -u ./...

clean:
	go clean

.PHONY: run build debs update-debs clean
