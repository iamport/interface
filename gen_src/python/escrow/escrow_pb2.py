# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: escrow/escrow.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='escrow/escrow.proto',
  package='escrow',
  syntax='proto3',
  serialized_options=b'Z.github.com/iamport/interface/gen_src/go/escrow',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x13\x65scrow/escrow.proto\x12\x06\x65scrow\"A\n\x04Info\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0b\n\x03tel\x18\x02 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x03 \x01(\t\x12\x10\n\x08postcode\x18\x04 \x01(\t\"N\n\x05Logis\x12\x0f\n\x07\x63ompany\x18\x01 \x01(\t\x12\x0f\n\x07invoice\x18\x02 \x01(\t\x12\x0f\n\x07sent_at\x18\x03 \x01(\x05\x12\x12\n\napplied_at\x18\x04 \x01(\x05\"k\n\rEscrowRequest\x12\x1c\n\x06sender\x18\x01 \x01(\x0b\x32\x0c.escrow.Info\x12\x1e\n\x08receiver\x18\x02 \x01(\x0b\x32\x0c.escrow.Info\x12\x1c\n\x05logis\x18\x03 \x01(\x0b\x32\r.escrow.Logis\"P\n\x0e\x45scrowResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x1f\n\x08response\x18\x03 \x01(\x0b\x32\r.escrow.LogisB0Z.github.com/iamport/interface/gen_src/go/escrowb\x06proto3'
)




_INFO = _descriptor.Descriptor(
  name='Info',
  full_name='escrow.Info',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='escrow.Info.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tel', full_name='escrow.Info.tel', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='addr', full_name='escrow.Info.addr', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='postcode', full_name='escrow.Info.postcode', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=96,
)


_LOGIS = _descriptor.Descriptor(
  name='Logis',
  full_name='escrow.Logis',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='company', full_name='escrow.Logis.company', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='invoice', full_name='escrow.Logis.invoice', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sent_at', full_name='escrow.Logis.sent_at', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='applied_at', full_name='escrow.Logis.applied_at', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=98,
  serialized_end=176,
)


_ESCROWREQUEST = _descriptor.Descriptor(
  name='EscrowRequest',
  full_name='escrow.EscrowRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='sender', full_name='escrow.EscrowRequest.sender', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='receiver', full_name='escrow.EscrowRequest.receiver', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='logis', full_name='escrow.EscrowRequest.logis', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=178,
  serialized_end=285,
)


_ESCROWRESPONSE = _descriptor.Descriptor(
  name='EscrowResponse',
  full_name='escrow.EscrowResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='escrow.EscrowResponse.code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='escrow.EscrowResponse.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='response', full_name='escrow.EscrowResponse.response', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=287,
  serialized_end=367,
)

_ESCROWREQUEST.fields_by_name['sender'].message_type = _INFO
_ESCROWREQUEST.fields_by_name['receiver'].message_type = _INFO
_ESCROWREQUEST.fields_by_name['logis'].message_type = _LOGIS
_ESCROWRESPONSE.fields_by_name['response'].message_type = _LOGIS
DESCRIPTOR.message_types_by_name['Info'] = _INFO
DESCRIPTOR.message_types_by_name['Logis'] = _LOGIS
DESCRIPTOR.message_types_by_name['EscrowRequest'] = _ESCROWREQUEST
DESCRIPTOR.message_types_by_name['EscrowResponse'] = _ESCROWRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Info = _reflection.GeneratedProtocolMessageType('Info', (_message.Message,), {
  'DESCRIPTOR' : _INFO,
  '__module__' : 'escrow.escrow_pb2'
  # @@protoc_insertion_point(class_scope:escrow.Info)
  })
_sym_db.RegisterMessage(Info)

Logis = _reflection.GeneratedProtocolMessageType('Logis', (_message.Message,), {
  'DESCRIPTOR' : _LOGIS,
  '__module__' : 'escrow.escrow_pb2'
  # @@protoc_insertion_point(class_scope:escrow.Logis)
  })
_sym_db.RegisterMessage(Logis)

EscrowRequest = _reflection.GeneratedProtocolMessageType('EscrowRequest', (_message.Message,), {
  'DESCRIPTOR' : _ESCROWREQUEST,
  '__module__' : 'escrow.escrow_pb2'
  # @@protoc_insertion_point(class_scope:escrow.EscrowRequest)
  })
_sym_db.RegisterMessage(EscrowRequest)

EscrowResponse = _reflection.GeneratedProtocolMessageType('EscrowResponse', (_message.Message,), {
  'DESCRIPTOR' : _ESCROWRESPONSE,
  '__module__' : 'escrow.escrow_pb2'
  # @@protoc_insertion_point(class_scope:escrow.EscrowResponse)
  })
_sym_db.RegisterMessage(EscrowResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
