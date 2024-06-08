.PHONY: protoc
protoc:
	protoc --go_out=. proto/*.proto
