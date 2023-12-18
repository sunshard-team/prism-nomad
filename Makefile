.PHONY:

version ?= v0.3.0

build-prism:
	go build -a -o build/prism -ldflags="-X 'prism/cmd.version=$(version)'"

build-cli:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -o build/prism.linux-386 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.linux-386 prism && tar czf prism.linux-386.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o build/prism.linux-amd64 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.linux-amd64 prism && tar czf prism.linux-amd64.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -o build/prism.linux-armv7 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.linux-armv7 prism && tar czf prism.linux-armv7.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o build/prism.linux-arm64 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.linux-arm64 prism && tar czf prism.linux-arm64.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o build/prism.darwin-amd64 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.darwin-amd64 prism && tar czf prism.darwin-amd64.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o build/prism.darwin-arm64 -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.darwin-arm64 prism && tar czf prism.darwin-arm64.tar.gz prism && rm prism && cd ..

	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -o build/prism.windows-386.exe -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.windows-386.exe prism.exe && zip prism.windows-386.zip prism.exe && rm prism.exe && cd ..

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o build/prism.windows-amd64.exe -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.windows-amd64.exe prism.exe && zip prism.windows-amd64.zip prism.exe && rm prism.exe && cd ..

	CGO_ENABLED=0 GOOS=windows GOARCH=arm GOARM=7 go build -a -o build/prism.windows-armv7.exe -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.windows-armv7.exe prism.exe && zip prism.windows-armv7.zip prism.exe && rm prism.exe && cd ..

	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -a -o build/prism.windows-arm64.exe -ldflags="-X 'prism/cmd.version=$(version)'"
	cd build && mv prism.windows-arm64.exe prism.exe && zip prism.windows-arm64.zip prism.exe && rm prism.exe && cd ..

	cd build && shasum -a 256 * > sha256sum.txt
	cat build/sha256sum.txt