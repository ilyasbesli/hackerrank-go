fmt:
	go fmt ./...

build:

    for f in algorithm/**/main.go; do \
        go build -o build/$$f.exe $$f; \
        echo compiled: $$f; \
    done
