protoc --go_out=. --go_opt=paths=source_relative ^
    --go-grpc_out=. --go-grpc_opt=paths=source_relative ^
    proto/booking/booking.proto proto/room/room.proto proto/user/user.proto
    