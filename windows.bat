protoc --proto_path=./internal --proto_path=./third_party --go_out=paths=source_relative:./internal ./internal/conf/*.proto 
protoc --proto_path=.  --proto_path=./third_party  --go_out=. --go-grpc_out=. --go-http_out=. --go-errors_out=. ./api/emqx/v1/*.proto
wire ./cmd/emqx_exhooks/