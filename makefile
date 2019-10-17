build:
	/usr/bin/protoc -I. --go_out=plugins=micro:. proto/walkie.proto