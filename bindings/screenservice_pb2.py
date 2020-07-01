# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: screenservice.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='screenservice.proto',
  package='screenservice',
  syntax='proto3',
  serialized_pb=_b('\n\x13screenservice.proto\x12\rscreenservice\"-\n\x04Line\x12\x11\n\tline_type\x18\x01 \x01(\t\x12\x12\n\nline_value\x18\x02 \x01(\t\"D\n\x07Timeout\x12\x0e\n\x06length\x18\x01 \x01(\x05\x12\x13\n\x0bshowtimeout\x18\x02 \x01(\x05\x12\x14\n\x0creturnscreen\x18\x03 \x01(\t\"4\n\x07\x43ommand\x12\x13\n\x0b\x63ommandtype\x18\x01 \x01(\t\x12\x14\n\x0c\x63ommandvalue\x18\x02 \x01(\t\"e\n\x05Touch\x12\t\n\x01x\x18\x01 \x01(\x05\x12\t\n\x01y\x18\x02 \x01(\x05\x12\r\n\x05width\x18\x03 \x01(\x05\x12\x0e\n\x06height\x18\x04 \x01(\x05\x12\'\n\x07\x63ommand\x18\x05 \x01(\x0b\x32\x16.screenservice.Command\"\x93\x02\n\rScreenRequest\x12\"\n\x05line1\x18\x01 \x01(\x0b\x32\x13.screenservice.Line\x12\"\n\x05line2\x18\x02 \x01(\x0b\x32\x13.screenservice.Line\x12\"\n\x05line3\x18\x03 \x01(\x0b\x32\x13.screenservice.Line\x12\"\n\x05line4\x18\x04 \x01(\x0b\x32\x13.screenservice.Line\x12\"\n\x05line5\x18\x05 \x01(\x0b\x32\x13.screenservice.Line\x12\'\n\x07timeout\x18\x06 \x01(\x0b\x32\x16.screenservice.Timeout\x12%\n\x07touches\x18\x07 \x03(\x0b\x32\x14.screenservice.Touch\"1\n\x0eScreenResponse\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\"O\n\x0bScreenImage\x12\x12\n\nimage_data\x18\x01 \x01(\x0c\x12\x0e\n\x06height\x18\x02 \x01(\x05\x12\r\n\x05width\x18\x03 \x01(\x05\x12\r\n\x05\x66rame\x18\x04 \x01(\x03\x32\xa5\x01\n\x0cScreenServer\x12K\n\nSendScreen\x12\x1c.screenservice.ScreenRequest\x1a\x1d.screenservice.ScreenResponse\"\x00\x12H\n\tSendImage\x12\x1a.screenservice.ScreenImage\x1a\x1d.screenservice.ScreenResponse\"\x00\x62\x06proto3')
)




_LINE = _descriptor.Descriptor(
  name='Line',
  full_name='screenservice.Line',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='line_type', full_name='screenservice.Line.line_type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='line_value', full_name='screenservice.Line.line_value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=38,
  serialized_end=83,
)


_TIMEOUT = _descriptor.Descriptor(
  name='Timeout',
  full_name='screenservice.Timeout',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='length', full_name='screenservice.Timeout.length', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='showtimeout', full_name='screenservice.Timeout.showtimeout', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='returnscreen', full_name='screenservice.Timeout.returnscreen', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=85,
  serialized_end=153,
)


_COMMAND = _descriptor.Descriptor(
  name='Command',
  full_name='screenservice.Command',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='commandtype', full_name='screenservice.Command.commandtype', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='commandvalue', full_name='screenservice.Command.commandvalue', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=155,
  serialized_end=207,
)


_TOUCH = _descriptor.Descriptor(
  name='Touch',
  full_name='screenservice.Touch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='x', full_name='screenservice.Touch.x', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='y', full_name='screenservice.Touch.y', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='width', full_name='screenservice.Touch.width', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='height', full_name='screenservice.Touch.height', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='screenservice.Touch.command', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=209,
  serialized_end=310,
)


_SCREENREQUEST = _descriptor.Descriptor(
  name='ScreenRequest',
  full_name='screenservice.ScreenRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='line1', full_name='screenservice.ScreenRequest.line1', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='line2', full_name='screenservice.ScreenRequest.line2', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='line3', full_name='screenservice.ScreenRequest.line3', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='line4', full_name='screenservice.ScreenRequest.line4', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='line5', full_name='screenservice.ScreenRequest.line5', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timeout', full_name='screenservice.ScreenRequest.timeout', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='touches', full_name='screenservice.ScreenRequest.touches', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=313,
  serialized_end=588,
)


_SCREENRESPONSE = _descriptor.Descriptor(
  name='ScreenResponse',
  full_name='screenservice.ScreenResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='screenservice.ScreenResponse.status', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='screenservice.ScreenResponse.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=590,
  serialized_end=639,
)


_SCREENIMAGE = _descriptor.Descriptor(
  name='ScreenImage',
  full_name='screenservice.ScreenImage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image_data', full_name='screenservice.ScreenImage.image_data', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='height', full_name='screenservice.ScreenImage.height', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='width', full_name='screenservice.ScreenImage.width', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='frame', full_name='screenservice.ScreenImage.frame', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=641,
  serialized_end=720,
)

_TOUCH.fields_by_name['command'].message_type = _COMMAND
_SCREENREQUEST.fields_by_name['line1'].message_type = _LINE
_SCREENREQUEST.fields_by_name['line2'].message_type = _LINE
_SCREENREQUEST.fields_by_name['line3'].message_type = _LINE
_SCREENREQUEST.fields_by_name['line4'].message_type = _LINE
_SCREENREQUEST.fields_by_name['line5'].message_type = _LINE
_SCREENREQUEST.fields_by_name['timeout'].message_type = _TIMEOUT
_SCREENREQUEST.fields_by_name['touches'].message_type = _TOUCH
DESCRIPTOR.message_types_by_name['Line'] = _LINE
DESCRIPTOR.message_types_by_name['Timeout'] = _TIMEOUT
DESCRIPTOR.message_types_by_name['Command'] = _COMMAND
DESCRIPTOR.message_types_by_name['Touch'] = _TOUCH
DESCRIPTOR.message_types_by_name['ScreenRequest'] = _SCREENREQUEST
DESCRIPTOR.message_types_by_name['ScreenResponse'] = _SCREENRESPONSE
DESCRIPTOR.message_types_by_name['ScreenImage'] = _SCREENIMAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Line = _reflection.GeneratedProtocolMessageType('Line', (_message.Message,), dict(
  DESCRIPTOR = _LINE,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.Line)
  ))
_sym_db.RegisterMessage(Line)

Timeout = _reflection.GeneratedProtocolMessageType('Timeout', (_message.Message,), dict(
  DESCRIPTOR = _TIMEOUT,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.Timeout)
  ))
_sym_db.RegisterMessage(Timeout)

Command = _reflection.GeneratedProtocolMessageType('Command', (_message.Message,), dict(
  DESCRIPTOR = _COMMAND,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.Command)
  ))
_sym_db.RegisterMessage(Command)

Touch = _reflection.GeneratedProtocolMessageType('Touch', (_message.Message,), dict(
  DESCRIPTOR = _TOUCH,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.Touch)
  ))
_sym_db.RegisterMessage(Touch)

ScreenRequest = _reflection.GeneratedProtocolMessageType('ScreenRequest', (_message.Message,), dict(
  DESCRIPTOR = _SCREENREQUEST,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.ScreenRequest)
  ))
_sym_db.RegisterMessage(ScreenRequest)

ScreenResponse = _reflection.GeneratedProtocolMessageType('ScreenResponse', (_message.Message,), dict(
  DESCRIPTOR = _SCREENRESPONSE,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.ScreenResponse)
  ))
_sym_db.RegisterMessage(ScreenResponse)

ScreenImage = _reflection.GeneratedProtocolMessageType('ScreenImage', (_message.Message,), dict(
  DESCRIPTOR = _SCREENIMAGE,
  __module__ = 'screenservice_pb2'
  # @@protoc_insertion_point(class_scope:screenservice.ScreenImage)
  ))
_sym_db.RegisterMessage(ScreenImage)



_SCREENSERVER = _descriptor.ServiceDescriptor(
  name='ScreenServer',
  full_name='screenservice.ScreenServer',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=723,
  serialized_end=888,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendScreen',
    full_name='screenservice.ScreenServer.SendScreen',
    index=0,
    containing_service=None,
    input_type=_SCREENREQUEST,
    output_type=_SCREENRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SendImage',
    full_name='screenservice.ScreenServer.SendImage',
    index=1,
    containing_service=None,
    input_type=_SCREENIMAGE,
    output_type=_SCREENRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SCREENSERVER)

DESCRIPTOR.services_by_name['ScreenServer'] = _SCREENSERVER

# @@protoc_insertion_point(module_scope)
