# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v2/txs/txs.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from v2.basis import basis_pb2 as v2_dot_basis_dot_basis__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v2/txs/txs.proto',
  package='txs_v2',
  syntax='proto3',
  serialized_options=b'Z.github.com/iamport/interface/gen_src/go/v2/txs\252\002\006V2.Txs',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10v2/txs/txs.proto\x12\x06txs_v2\x1a\x1cgoogle/api/annotations.proto\x1a\x14v2/basis/basis.proto\"}\n\x14GetTxsRecordsRequest\x12\x0f\n\x07imp_uid\x18\x01 \x01(\t\x12\x14\n\x0cmerchant_uid\x18\x02 \x01(\t\x12\x0f\n\x07paid_by\x18\x03 \x01(\t\x12\x0e\n\x06status\x18\x04 \x01(\t\x12\x0c\n\x04page\x18\x05 \x01(\x05\x12\x0f\n\x07sorting\x18\x06 \x01(\t\"Y\n\x14GetTxsRecordsReponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\"\n\x08response\x18\x03 \x03(\x0b\x32\x10.basis_v2.UnitTx2}\n\nTxsService\x12o\n\x13GetTxsRecordService\x12\x1c.txs_v2.GetTxsRecordsRequest\x1a\x1c.txs_v2.GetTxsRecordsReponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/api/payments/v2/txsB9Z.github.com/iamport/interface/gen_src/go/v2/txs\xaa\x02\x06V2.Txsb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,v2_dot_basis_dot_basis__pb2.DESCRIPTOR,])




_GETTXSRECORDSREQUEST = _descriptor.Descriptor(
  name='GetTxsRecordsRequest',
  full_name='txs_v2.GetTxsRecordsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='imp_uid', full_name='txs_v2.GetTxsRecordsRequest.imp_uid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='merchant_uid', full_name='txs_v2.GetTxsRecordsRequest.merchant_uid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='paid_by', full_name='txs_v2.GetTxsRecordsRequest.paid_by', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='txs_v2.GetTxsRecordsRequest.status', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='page', full_name='txs_v2.GetTxsRecordsRequest.page', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sorting', full_name='txs_v2.GetTxsRecordsRequest.sorting', index=5,
      number=6, type=9, cpp_type=9, label=1,
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
  serialized_start=80,
  serialized_end=205,
)


_GETTXSRECORDSREPONSE = _descriptor.Descriptor(
  name='GetTxsRecordsReponse',
  full_name='txs_v2.GetTxsRecordsReponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='txs_v2.GetTxsRecordsReponse.code', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='txs_v2.GetTxsRecordsReponse.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='response', full_name='txs_v2.GetTxsRecordsReponse.response', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=207,
  serialized_end=296,
)

_GETTXSRECORDSREPONSE.fields_by_name['response'].message_type = v2_dot_basis_dot_basis__pb2._UNITTX
DESCRIPTOR.message_types_by_name['GetTxsRecordsRequest'] = _GETTXSRECORDSREQUEST
DESCRIPTOR.message_types_by_name['GetTxsRecordsReponse'] = _GETTXSRECORDSREPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetTxsRecordsRequest = _reflection.GeneratedProtocolMessageType('GetTxsRecordsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTXSRECORDSREQUEST,
  '__module__' : 'v2.txs.txs_pb2'
  # @@protoc_insertion_point(class_scope:txs_v2.GetTxsRecordsRequest)
  })
_sym_db.RegisterMessage(GetTxsRecordsRequest)

GetTxsRecordsReponse = _reflection.GeneratedProtocolMessageType('GetTxsRecordsReponse', (_message.Message,), {
  'DESCRIPTOR' : _GETTXSRECORDSREPONSE,
  '__module__' : 'v2.txs.txs_pb2'
  # @@protoc_insertion_point(class_scope:txs_v2.GetTxsRecordsReponse)
  })
_sym_db.RegisterMessage(GetTxsRecordsReponse)


DESCRIPTOR._options = None

_TXSSERVICE = _descriptor.ServiceDescriptor(
  name='TxsService',
  full_name='txs_v2.TxsService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=298,
  serialized_end=423,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetTxsRecordService',
    full_name='txs_v2.TxsService.GetTxsRecordService',
    index=0,
    containing_service=None,
    input_type=_GETTXSRECORDSREQUEST,
    output_type=_GETTXSRECORDSREPONSE,
    serialized_options=b'\202\323\344\223\002\026\022\024/api/payments/v2/txs',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TXSSERVICE)

DESCRIPTOR.services_by_name['TxsService'] = _TXSSERVICE

# @@protoc_insertion_point(module_scope)
