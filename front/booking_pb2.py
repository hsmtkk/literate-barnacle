# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: booking.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rbooking.proto\x12\x07\x62ooking\"t\n\nNewBooking\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x0f\n\x07room_id\x18\x02 \x01(\x03\x12\x14\n\x0creserved_num\x18\x03 \x01(\x03\x12\x17\n\x0fstart_date_time\x18\x04 \x01(\t\x12\x15\n\rend_date_time\x18\x05 \x01(\t\"}\n\x07\x42ooking\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0f\n\x07user_id\x18\x02 \x01(\x03\x12\x0f\n\x07room_id\x18\x03 \x01(\x03\x12\x14\n\x0creserved_num\x18\x04 \x01(\x03\x12\x17\n\x0fstart_date_time\x18\x05 \x01(\t\x12\x15\n\rend_date_time\x18\x06 \x01(\t\"\x12\n\x10\x44ummyListRequest\"7\n\x11\x44ummyListResponse\x12\"\n\x08\x62ookings\x18\x01 \x03(\x0b\x32\x10.booking.Booking\"5\n\rCreateReqeust\x12$\n\x07\x62ooking\x18\x01 \x01(\x0b\x32\x13.booking.NewBooking\"3\n\x0e\x43reateResponse\x12!\n\x07\x62ooking\x18\x01 \x01(\x0b\x32\x10.booking.Booking2\x93\x01\n\x0e\x42ookingService\x12\x44\n\tDummyList\x12\x19.booking.DummyListRequest\x1a\x1a.booking.DummyListResponse\"\x00\x12;\n\x06\x43reate\x12\x16.booking.CreateReqeust\x1a\x17.booking.CreateResponse\"\x00\x42\x33Z1github.com/hsmtkk/literate-barnacle/proto/bookingb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'booking_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z1github.com/hsmtkk/literate-barnacle/proto/booking'
  _NEWBOOKING._serialized_start=26
  _NEWBOOKING._serialized_end=142
  _BOOKING._serialized_start=144
  _BOOKING._serialized_end=269
  _DUMMYLISTREQUEST._serialized_start=271
  _DUMMYLISTREQUEST._serialized_end=289
  _DUMMYLISTRESPONSE._serialized_start=291
  _DUMMYLISTRESPONSE._serialized_end=346
  _CREATEREQEUST._serialized_start=348
  _CREATEREQEUST._serialized_end=401
  _CREATERESPONSE._serialized_start=403
  _CREATERESPONSE._serialized_end=454
  _BOOKINGSERVICE._serialized_start=457
  _BOOKINGSERVICE._serialized_end=604
# @@protoc_insertion_point(module_scope)
