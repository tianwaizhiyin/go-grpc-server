cd pbfiles && protoc --go_out=plugins=grpc:../services Prod.proto
protoc --go_out=plugins=grpc:../services Orders.proto
protoc --go_out=plugins=grpc:../services Models.proto
cd ..