package genshared

const NoEditHead = `DO NOT EDIT. Generated by protoc-gen-dart-ext/tools.`
const NoEditCommentHead = "// " + NoEditHead + "\n"

const FormatProtoHead = NoEditCommentHead + `syntax = "proto3";

package pgde.format;

option go_package = "github.com/empirefox/protoc-gen-dart-ext/pkg/format";
`

const UnitsProtoHead = NoEditCommentHead + `syntax = "proto3";

package pgde.units;

option go_package = "github.com/empirefox/protoc-gen-dart-ext/pkg/units";
`

const DartHead = NoEditCommentHead
