gen_protos:
	protoc --go_out . --go-grpc_out . -I ./../junk-dealer/protos ./../junk-dealer/protos/*.proto
