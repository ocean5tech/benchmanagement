# 设置变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=benchmanagement
all: clear clean run 
build: clear clean bc 
clean:
	$(GOCLEAN)
	del /F $(BINARY_NAME).exe 
clear:
	cls && echo -------------START BUILD-------------
bc:
	$(GOBUILD) -o ./bin/$(BINARY_NAME).exe -v 
run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME).exe -v 
	./bin/$(BINARY_NAME).exe
