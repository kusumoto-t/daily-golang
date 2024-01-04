#
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

NOW=$(shell date)

DATE=$(shell date +"%Y%m%d")

now:
	@echo $(NOW)
new-prac:
	mkdir -p $(DATE)/
	touch $(DATE)/main.go
	echo "package main" > $(DATE)/main.go
	cd $(DATE)
goconfirm:
	@echo $(GOVERSION)
	@echo $(GOOS)
	@echo $(GOARCH)
lint:
	golint -set_exit_status $$(go list ./)
	go ver ./