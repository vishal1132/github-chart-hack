buildMac:
	@GOOS=darwin GOARCH=amd64 go build -o charthackunix .

buildLinux:
	@GOOS=LINUX go build -o charthacklinux .

buildWindows:
	@GOOS=windows GOARCH=386 go build -o charthackwin.exe .