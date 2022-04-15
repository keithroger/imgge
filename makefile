TESTIMAGES=$(CURDIR)/test_images

clean:
	@go clean
	@if [ -d $(TESTIMAGES) ] ; then rm -r $(TESTIMAGES) ; fi

imagedir:
	@if [ ! -d $(TESTIMAGES) ] ; then mkdir -p $(TESTIMAGES) ; fi

test: imagedir
	go test ./...

cover: imagedir
	go test ./... -cover

lint:
	golangci-lint run --enable-all

fmt:
	gofmt -w *.go
