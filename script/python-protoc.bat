call python -m grpc_tools.protoc -Iproto/booking --python_out=front/ --pyi_out=front/ --grpc_python_out=front/ proto/booking/booking.proto
call python -m grpc_tools.protoc -Iproto/room --python_out=front/ --pyi_out=front/ --grpc_python_out=front/ proto/room/room.proto
call python -m grpc_tools.protoc -Iproto/user --python_out=front/ --pyi_out=front/ --grpc_python_out=front/ proto/user/user.proto
