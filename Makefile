#
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

NOW=$(shell date)

DATE=$(shell date +"%Y%m%d")

# 初期化設定
INIT ?=yes

now:
	@echo $(NOW)
new-prac:
	mkdir -p $(DATE)/
ifeq ($(INIT), no)
		echo "Skipping touch initial file"
else
		touch $(DATE)/main.go
		echo "package main" > $(DATE)/main.go
endif
prac-push:
	git add .
	git commit -m "add $(DATE)"
	git push
goconfirm:
	@echo $(GOVERSION)
	@echo $(GOOS)
	@echo $(GOARCH)
lint:
	golint -set_exit_status $$(go list ./)
	go ver ./
