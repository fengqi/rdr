GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'

all: linux-amd64 darwin-amd64 windows-amd64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o rdr-darwin

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o rdr-linux

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o rdr-windows.exe

clean:
	rm rdr-*