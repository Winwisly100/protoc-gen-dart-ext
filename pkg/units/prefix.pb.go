// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/units/prefix.proto

package units // import "github.com/empirefox/protoc-gen-dart-ext/pkg/units"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/empirefox/protoc-gen-dart-ext/pkg/l10n"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// https://physics.nist.gov/cuu/Units/prefixes.html
// https://physics.nist.gov/cuu/Units/binary.html
type PrefixV1 int32

const (
	PrefixV1_noPrefix PrefixV1 = 0
	// SI
	// Y, 10^24
	PrefixV1_yotta PrefixV1 = 1
	// Z, 10^21
	PrefixV1_zetta PrefixV1 = 2
	// E, 10^18
	PrefixV1_exa PrefixV1 = 3
	// P, 10^15
	PrefixV1_peta PrefixV1 = 4
	// T, 10^12
	PrefixV1_tera PrefixV1 = 5
	// G, 10^9
	PrefixV1_giga PrefixV1 = 6
	// M, 10^6
	PrefixV1_mega PrefixV1 = 7
	// k, 10^3
	PrefixV1_kilo PrefixV1 = 8
	// h, 10^2
	PrefixV1_hecto PrefixV1 = 9
	// da, 10^1
	PrefixV1_deka PrefixV1 = 10
	// d, 10^-1
	PrefixV1_deci PrefixV1 = 11
	// c, 10^-2
	PrefixV1_centi PrefixV1 = 12
	// m, 10^-3
	PrefixV1_milli PrefixV1 = 13
	// µ, 10^-6
	PrefixV1_micro PrefixV1 = 14
	// n, 10^-9
	PrefixV1_nano PrefixV1 = 15
	// p, 10^-12
	PrefixV1_pico PrefixV1 = 16
	// f, 10^-15
	PrefixV1_femto PrefixV1 = 17
	// a, 10^-18
	PrefixV1_atto PrefixV1 = 18
	// z, 10^-21
	PrefixV1_zepto PrefixV1 = 19
	// y, 10^-24
	PrefixV1_yocto PrefixV1 = 20
	// binary
	// Ki, 2^10
	PrefixV1_kibi PrefixV1 = 21
	// Mi, 2^20
	PrefixV1_mebi PrefixV1 = 22
	// Gi, 2^30
	PrefixV1_gibi PrefixV1 = 23
	// Ti, 2^40
	PrefixV1_tebi PrefixV1 = 24
	// Pi, 2^50
	PrefixV1_pebi PrefixV1 = 25
	// Ei, 2^60
	PrefixV1_exbi PrefixV1 = 26
)

var PrefixV1_name = map[int32]string{
	0:  "noPrefix",
	1:  "yotta",
	2:  "zetta",
	3:  "exa",
	4:  "peta",
	5:  "tera",
	6:  "giga",
	7:  "mega",
	8:  "kilo",
	9:  "hecto",
	10: "deka",
	11: "deci",
	12: "centi",
	13: "milli",
	14: "micro",
	15: "nano",
	16: "pico",
	17: "femto",
	18: "atto",
	19: "zepto",
	20: "yocto",
	21: "kibi",
	22: "mebi",
	23: "gibi",
	24: "tebi",
	25: "pebi",
	26: "exbi",
}
var PrefixV1_value = map[string]int32{
	"noPrefix": 0,
	"yotta":    1,
	"zetta":    2,
	"exa":      3,
	"peta":     4,
	"tera":     5,
	"giga":     6,
	"mega":     7,
	"kilo":     8,
	"hecto":    9,
	"deka":     10,
	"deci":     11,
	"centi":    12,
	"milli":    13,
	"micro":    14,
	"nano":     15,
	"pico":     16,
	"femto":    17,
	"atto":     18,
	"zepto":    19,
	"yocto":    20,
	"kibi":     21,
	"mebi":     22,
	"gibi":     23,
	"tebi":     24,
	"pebi":     25,
	"exbi":     26,
}

func (x PrefixV1) String() string {
	return proto.EnumName(PrefixV1_name, int32(x))
}
func (PrefixV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_prefix_73146b1729385ef3, []int{0}
}

func init() {
	proto.RegisterEnum("units.PrefixV1", PrefixV1_name, PrefixV1_value)
}

func init() { proto.RegisterFile("protos/units/prefix.proto", fileDescriptor_prefix_73146b1729385ef3) }

var fileDescriptor_prefix_73146b1729385ef3 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x2c, 0xd1, 0x4b, 0x52, 0xf3, 0x30,
	0x0c, 0x07, 0xf0, 0xef, 0xeb, 0x33, 0x35, 0x2f, 0x11, 0x68, 0xa1, 0x3d, 0x02, 0x33, 0x6d, 0x28,
	0x70, 0x02, 0x4e, 0xc0, 0x8a, 0x05, 0x3b, 0xd7, 0x55, 0x53, 0x4d, 0x12, 0xdb, 0x93, 0xba, 0x33,
	0x81, 0x3b, 0xb1, 0xe1, 0x04, 0x1c, 0x86, 0x83, 0x60, 0x59, 0x6c, 0x3a, 0xbf, 0xfe, 0xa5, 0xb1,
	0x25, 0x47, 0xcd, 0x7d, 0xeb, 0x82, 0x3b, 0x14, 0x47, 0x4b, 0xe1, 0x50, 0xf8, 0x16, 0x77, 0xd4,
	0xad, 0x52, 0x96, 0x0f, 0x53, 0xb6, 0x98, 0xfd, 0x75, 0xd4, 0xeb, 0x7b, 0x9b, 0x7e, 0xa4, 0x7c,
	0xf7, 0xd9, 0x53, 0xd9, 0x4b, 0xea, 0x7f, 0x5d, 0xe7, 0x53, 0x95, 0x59, 0x27, 0xff, 0xe0, 0xdf,
	0x62, 0xfc, 0xf5, 0xf3, 0xdd, 0xef, 0x65, 0xff, 0xf3, 0x89, 0x1a, 0xbe, 0xbb, 0x10, 0x34, 0x24,
	0x7e, 0x20, 0xb3, 0x97, 0x8f, 0x55, 0x1f, 0x3b, 0x0d, 0xfd, 0x3c, 0x53, 0x03, 0x8f, 0x31, 0x1a,
	0xb0, 0x02, 0xb6, 0x1a, 0x86, 0xac, 0x92, 0x4a, 0x0d, 0x23, 0x56, 0x83, 0x51, 0x63, 0x56, 0x45,
	0xb5, 0x83, 0x8c, 0x4f, 0xd9, 0xa3, 0x09, 0x0e, 0x26, 0x1c, 0x6e, 0xb1, 0xd2, 0xa0, 0x44, 0x86,
	0xe0, 0x84, 0xcb, 0x06, 0x6d, 0x20, 0x38, 0x65, 0x36, 0x54, 0xd7, 0x04, 0x67, 0x42, 0xd3, 0x3a,
	0x38, 0xe7, 0x56, 0xab, 0xad, 0x83, 0x8b, 0x74, 0x37, 0x19, 0x07, 0xc0, 0xe5, 0x1d, 0x36, 0xf1,
	0xcc, 0x4b, 0x0e, 0x75, 0x88, 0xca, 0x65, 0x5c, 0x1f, 0x79, 0x25, 0x4b, 0xf0, 0x9d, 0xd7, 0x32,
	0xc8, 0x86, 0x60, 0x2a, 0xc3, 0x45, 0xcd, 0x64, 0xe0, 0xa8, 0x1b, 0x59, 0x22, 0xea, 0x56, 0x16,
	0x8b, 0x9a, 0xb3, 0xb0, 0x8b, 0x5a, 0x3c, 0x3f, 0xbd, 0x3d, 0x94, 0x14, 0xf6, 0xc7, 0xcd, 0xca,
	0xb8, 0xa6, 0xc0, 0xc6, 0x53, 0x7c, 0x2e, 0xd7, 0x15, 0xe9, 0x35, 0xcd, 0xb2, 0x44, 0xbb, 0xdc,
	0xea, 0x36, 0x2c, 0xb1, 0x0b, 0x85, 0xaf, 0x4a, 0xf9, 0x22, 0x9b, 0x51, 0x2a, 0x3f, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x67, 0xcf, 0x08, 0xe9, 0xa8, 0x01, 0x00, 0x00,
}