# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Main package name
MAIN_PKG=main.go

# Binary names
BINARY_NAME=employee_system

# make 명령어를 실행하면 all 타겟이 실행되어 test와 build가 차례대로 실행
all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PKG)

test:
	$(GOTEST) -v ./...

# clean은 빌드한 바이너리 파일을 삭제
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# run은 build를 실행하고 바로 프로그램을 실행
run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PKG)
	./$(BINARY_NAME)

