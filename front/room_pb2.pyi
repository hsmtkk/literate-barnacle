from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateReqeust(_message.Message):
    __slots__ = ["room"]
    ROOM_FIELD_NUMBER: _ClassVar[int]
    room: NewRoom
    def __init__(self, room: _Optional[_Union[NewRoom, _Mapping]] = ...) -> None: ...

class CreateResponse(_message.Message):
    __slots__ = ["room"]
    ROOM_FIELD_NUMBER: _ClassVar[int]
    room: Room
    def __init__(self, room: _Optional[_Union[Room, _Mapping]] = ...) -> None: ...

class DummyListRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class DummyListResponse(_message.Message):
    __slots__ = ["rooms"]
    ROOMS_FIELD_NUMBER: _ClassVar[int]
    rooms: _containers.RepeatedCompositeFieldContainer[Room]
    def __init__(self, rooms: _Optional[_Iterable[_Union[Room, _Mapping]]] = ...) -> None: ...

class NewRoom(_message.Message):
    __slots__ = ["capacity", "name"]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    capacity: int
    name: str
    def __init__(self, name: _Optional[str] = ..., capacity: _Optional[int] = ...) -> None: ...

class Room(_message.Message):
    __slots__ = ["capacity", "id", "name"]
    CAPACITY_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    capacity: int
    id: int
    name: str
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., capacity: _Optional[int] = ...) -> None: ...
