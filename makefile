search:
	go build -o tt_search
	./tt_search -keyword="mask" -find_times=5

build:
	go build -o tt_search
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o linux_tt_search
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mac_tt_search
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o windows_tt_search
	