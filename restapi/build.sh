make gen-swag
cd grpcserver && go build && cd .. && cd cmd/qpdistributor && go build && cd ../..
