-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('persion_pb')


local INFO = protobuf.Descriptor();
local INFO_TAG_FIELD = protobuf.FieldDescriptor();
local LIST = protobuf.Descriptor();
local LIST_INFOS_FIELD = protobuf.FieldDescriptor();

INFO_TAG_FIELD.name = "Tag"
INFO_TAG_FIELD.full_name = ".models.Info.Tag"
INFO_TAG_FIELD.number = 1
INFO_TAG_FIELD.index = 0
INFO_TAG_FIELD.label = 1
INFO_TAG_FIELD.has_default_value = false
INFO_TAG_FIELD.default_value = 0
INFO_TAG_FIELD.type = 5
INFO_TAG_FIELD.cpp_type = 1

INFO.name = "Info"
INFO.full_name = ".models.Info"
INFO.nested_types = {}
INFO.enum_types = {}
INFO.fields = {INFO_TAG_FIELD}
INFO.is_extendable = false
INFO.extensions = {}
LIST_INFOS_FIELD.name = "Infos"
LIST_INFOS_FIELD.full_name = ".models.List.Infos"
LIST_INFOS_FIELD.number = 1
LIST_INFOS_FIELD.index = 0
LIST_INFOS_FIELD.label = 3
LIST_INFOS_FIELD.has_default_value = false
LIST_INFOS_FIELD.default_value = {}
LIST_INFOS_FIELD.message_type = INFO
LIST_INFOS_FIELD.type = 11
LIST_INFOS_FIELD.cpp_type = 10

LIST.name = "List"
LIST.full_name = ".models.List"
LIST.nested_types = {}
LIST.enum_types = {}
LIST.fields = {LIST_INFOS_FIELD}
LIST.is_extendable = false
LIST.extensions = {}

Info = protobuf.Message(INFO)
List = protobuf.Message(LIST)