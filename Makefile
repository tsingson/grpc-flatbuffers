

generate_fbs:
	flatc --go --grpc bookmarks.fbs

generate_proto:
    protoc  -I=/Users/qinshen/go/src -I=/usr/local/include  -I=./ --gofast_out=plugins=grpc:bookmarkspb  ./*.proto

compile: compile_bookmarks_client compile_bookmarks_server

compile_bookmarks_client:
	cd bookmarks-client && go build -o ../client && cd ..

compile_bookmarks_server:
	cd bookmarks-server && go build -o ../server && cd ..

.PHONY: generate_fbs generate_proto compile compile_bookmarks_client compile_bookmarks_server