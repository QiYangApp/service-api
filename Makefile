BUILD_ENV := CGO_ENABLED=0
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

TARGET_EXEC := clone
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: all clean setup build-linux build-osx build-windows lib air dev generate

all: clean setup build-linux build-osx build-windows lib air
clean:
	rm -rf build
setup:
	mkdir -p build/linux
	mkdir -p build/osx
	mkdir -p build/windows
	mkdir -p build/lib
build-linux: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o build/linux/${TARGET_EXEC} -buildvcs=false
build-osx: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o build/osx/${TARGET_EXEC} -buildvcs=false
build-windows: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o build/windows/${TARGET_EXEC}.exe -buildvcs=false
#lib:
#	GOARCH=amd64 go build ${LDFLAGS} -buildmode=c-shared -buildvcs=false -o build/windows/${TARGET_EXEC}.so main.go

dev: generate
	go run ./src/main.go
air: generate
	air build -buildvcs=false
generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate .\src\models\ent\schema


gofmt:
	echo "正在使用gofmt格式化文件..."
	gofmt -s -w ${GOFILES}
	echo "格式化完成"
lint:
	golangci-lint run --enable-all
govet:
		echo "正在进行静态检测..."
		go vet $(VETPACKAGES)
