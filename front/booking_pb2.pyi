from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Booking(_message.Message):
    __slots__ = ["end_date_time", "id", "reserved_num", "room_id", "start_date_time", "user_id"]
    END_DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    RESERVED_NUM_FIELD_NUMBER: _ClassVar[int]
    ROOM_ID_FIELD_NUMBER: _ClassVar[int]
    START_DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    end_date_time: int
    id: int
    reserved_num: int
    room_id: int
    start_date_time: int
    user_id: int
    def __init__(self, id: _Optional[int] = ..., user_id: _Optional[int] = ..., room_id: _Optional[int] = ..., reserved_num: _Optional[int] = ..., start_date_time: _Optional[int] = ..., end_date_time: _Optional[int] = ...) -> None: ...

class CreateReqeust(_message.Message):
    __slots__ = ["booking"]
    BOOKING_FIELD_NUMBER: _ClassVar[int]
    booking: NewBooking
    def __init__(self, booking: _Optional[_Union[NewBooking, _Mapping]] = ...) -> None: ...

class CreateResponse(_message.Message):
    __slots__ = ["booking"]
    BOOKING_FIELD_NUMBER: _ClassVar[int]
    booking: Booking
    def __init__(self, booking: _Optional[_Union[Booking, _Mapping]] = ...) -> None: ...

class DummyListRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class DummyListResponse(_message.Message):
    __slots__ = ["bookings"]
    BOOKINGS_FIELD_NUMBER: _ClassVar[int]
    bookings: _containers.RepeatedCompositeFieldContainer[Booking]
    def __init__(self, bookings: _Optional[_Iterable[_Union[Booking, _Mapping]]] = ...) -> None: ...

class NewBooking(_message.Message):
    __slots__ = ["end_date_time", "reserved_num", "room_id", "start_date_time", "user_id"]
    END_DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    RESERVED_NUM_FIELD_NUMBER: _ClassVar[int]
    ROOM_ID_FIELD_NUMBER: _ClassVar[int]
    START_DATE_TIME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    end_date_time: int
    reserved_num: int
    room_id: int
    start_date_time: int
    user_id: int
    def __init__(self, user_id: _Optional[int] = ..., room_id: _Optional[int] = ..., reserved_num: _Optional[int] = ..., start_date_time: _Optional[int] = ..., end_date_time: _Optional[int] = ...) -> None: ...
