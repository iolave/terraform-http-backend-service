build:
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = amd64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = 386)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go
	
	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go

	$(eval $@GOOS = linux)
	$(eval $@GOARCH = arm64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = arm64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go

	$(eval $@GOOS = darwin)
	$(eval $@GOARCH = amd64)
	GOOS=$($@GOOS) GOARCH=$($@GOARCH) go build -o "dist/api-$($@GOOS)-$($@GOARCH)" cmd/api/main.go

serve:
	go run cmd/api/main.go
