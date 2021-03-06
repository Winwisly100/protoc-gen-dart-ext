// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pgde/format/currency.proto

package format

import (
	fmt "fmt"
	_ "github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/zero"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ISO 4217, publish date: 2018-08-29
type Currency int32

const (
	Currency_XXX Currency = 0
	Currency_ALL Currency = 8
	Currency_DZD Currency = 12
	Currency_ARS Currency = 32
	Currency_AUD Currency = 36
	Currency_BSD Currency = 44
	Currency_BHD Currency = 48
	Currency_BDT Currency = 50
	Currency_AMD Currency = 51
	Currency_BBD Currency = 52
	Currency_BMD Currency = 60
	Currency_BTN Currency = 64
	Currency_BOB Currency = 68
	Currency_BWP Currency = 72
	Currency_BZD Currency = 84
	Currency_SBD Currency = 90
	Currency_BND Currency = 96
	Currency_MMK Currency = 104
	Currency_BIF Currency = 108
	Currency_KHR Currency = 116
	Currency_CAD Currency = 124
	Currency_CVE Currency = 132
	Currency_KYD Currency = 136
	Currency_LKR Currency = 144
	Currency_CLP Currency = 152
	Currency_CNY Currency = 156
	Currency_COP Currency = 170
	Currency_KMF Currency = 174
	Currency_CRC Currency = 188
	Currency_HRK Currency = 191
	Currency_CUP Currency = 192
	Currency_CZK Currency = 203
	Currency_DKK Currency = 208
	Currency_DOP Currency = 214
	Currency_SVC Currency = 222
	Currency_ETB Currency = 230
	Currency_ERN Currency = 232
	Currency_FKP Currency = 238
	Currency_FJD Currency = 242
	Currency_DJF Currency = 262
	Currency_GMD Currency = 270
	Currency_GIP Currency = 292
	Currency_GTQ Currency = 320
	Currency_GNF Currency = 324
	Currency_GYD Currency = 328
	Currency_HTG Currency = 332
	Currency_HNL Currency = 340
	Currency_HKD Currency = 344
	Currency_HUF Currency = 348
	Currency_ISK Currency = 352
	Currency_INR Currency = 356
	Currency_IDR Currency = 360
	Currency_IRR Currency = 364
	Currency_IQD Currency = 368
	Currency_ILS Currency = 376
	Currency_JMD Currency = 388
	Currency_JPY Currency = 392
	Currency_KZT Currency = 398
	Currency_JOD Currency = 400
	Currency_KES Currency = 404
	Currency_KPW Currency = 408
	Currency_KRW Currency = 410
	Currency_KWD Currency = 414
	Currency_KGS Currency = 417
	Currency_LAK Currency = 418
	Currency_LBP Currency = 422
	Currency_LSL Currency = 426
	Currency_LRD Currency = 430
	Currency_LYD Currency = 434
	Currency_MOP Currency = 446
	Currency_MWK Currency = 454
	Currency_MYR Currency = 458
	Currency_MVR Currency = 462
	Currency_MUR Currency = 480
	Currency_MXN Currency = 484
	Currency_MNT Currency = 496
	Currency_MDL Currency = 498
	Currency_MAD Currency = 504
	Currency_OMR Currency = 512
	Currency_NAD Currency = 516
	Currency_NPR Currency = 524
	Currency_ANG Currency = 532
	Currency_AWG Currency = 533
	Currency_VUV Currency = 548
	Currency_NZD Currency = 554
	Currency_NIO Currency = 558
	Currency_NGN Currency = 566
	Currency_NOK Currency = 578
	Currency_PKR Currency = 586
	Currency_PAB Currency = 590
	Currency_PGK Currency = 598
	Currency_PYG Currency = 600
	Currency_PEN Currency = 604
	Currency_PHP Currency = 608
	Currency_QAR Currency = 634
	Currency_RUB Currency = 643
	Currency_RWF Currency = 646
	Currency_SHP Currency = 654
	Currency_SAR Currency = 682
	Currency_SCR Currency = 690
	Currency_SLL Currency = 694
	Currency_SGD Currency = 702
	Currency_VND Currency = 704
	Currency_SOS Currency = 706
	Currency_ZAR Currency = 710
	Currency_SSP Currency = 728
	Currency_SZL Currency = 748
	Currency_SEK Currency = 752
	Currency_CHF Currency = 756
	Currency_SYP Currency = 760
	Currency_THB Currency = 764
	Currency_TOP Currency = 776
	Currency_TTD Currency = 780
	Currency_AED Currency = 784
	Currency_TND Currency = 788
	Currency_UGX Currency = 800
	Currency_MKD Currency = 807
	Currency_EGP Currency = 818
	Currency_GBP Currency = 826
	Currency_TZS Currency = 834
	Currency_USD Currency = 840
	Currency_UYU Currency = 858
	Currency_UZS Currency = 860
	Currency_WST Currency = 882
	Currency_YER Currency = 886
	Currency_TWD Currency = 901
	Currency_UYW Currency = 927
	Currency_VES Currency = 928
	Currency_MRU Currency = 929
	Currency_STN Currency = 930
	Currency_CUC Currency = 931
	Currency_ZWL Currency = 932
	Currency_BYN Currency = 933
	Currency_TMT Currency = 934
	Currency_GHS Currency = 936
	Currency_SDG Currency = 938
	Currency_UYI Currency = 940
	Currency_RSD Currency = 941
	Currency_MZN Currency = 943
	Currency_AZN Currency = 944
	Currency_RON Currency = 946
	Currency_CHE Currency = 947
	Currency_CHW Currency = 948
	Currency_TRY Currency = 949
	Currency_XAF Currency = 950
	Currency_XCD Currency = 951
	Currency_XOF Currency = 952
	Currency_XPF Currency = 953
	Currency_XBA Currency = 955
	Currency_XBB Currency = 956
	Currency_XBC Currency = 957
	Currency_XBD Currency = 958
	Currency_XAU Currency = 959
	Currency_XDR Currency = 960
	Currency_XAG Currency = 961
	Currency_XPT Currency = 962
	Currency_XTS Currency = 963
	Currency_XPD Currency = 964
	Currency_XUA Currency = 965
	Currency_ZMW Currency = 967
	Currency_SRD Currency = 968
	Currency_MGA Currency = 969
	Currency_COU Currency = 970
	Currency_AFN Currency = 971
	Currency_TJS Currency = 972
	Currency_AOA Currency = 973
	Currency_BGN Currency = 975
	Currency_CDF Currency = 976
	Currency_BAM Currency = 977
	Currency_EUR Currency = 978
	Currency_MXV Currency = 979
	Currency_UAH Currency = 980
	Currency_GEL Currency = 981
	Currency_BOV Currency = 984
	Currency_PLN Currency = 985
	Currency_BRL Currency = 986
	Currency_CLF Currency = 990
	Currency_XSU Currency = 994
	Currency_USN Currency = 997
)

var Currency_name = map[int32]string{
	0:   "XXX",
	8:   "ALL",
	12:  "DZD",
	32:  "ARS",
	36:  "AUD",
	44:  "BSD",
	48:  "BHD",
	50:  "BDT",
	51:  "AMD",
	52:  "BBD",
	60:  "BMD",
	64:  "BTN",
	68:  "BOB",
	72:  "BWP",
	84:  "BZD",
	90:  "SBD",
	96:  "BND",
	104: "MMK",
	108: "BIF",
	116: "KHR",
	124: "CAD",
	132: "CVE",
	136: "KYD",
	144: "LKR",
	152: "CLP",
	156: "CNY",
	170: "COP",
	174: "KMF",
	188: "CRC",
	191: "HRK",
	192: "CUP",
	203: "CZK",
	208: "DKK",
	214: "DOP",
	222: "SVC",
	230: "ETB",
	232: "ERN",
	238: "FKP",
	242: "FJD",
	262: "DJF",
	270: "GMD",
	292: "GIP",
	320: "GTQ",
	324: "GNF",
	328: "GYD",
	332: "HTG",
	340: "HNL",
	344: "HKD",
	348: "HUF",
	352: "ISK",
	356: "INR",
	360: "IDR",
	364: "IRR",
	368: "IQD",
	376: "ILS",
	388: "JMD",
	392: "JPY",
	398: "KZT",
	400: "JOD",
	404: "KES",
	408: "KPW",
	410: "KRW",
	414: "KWD",
	417: "KGS",
	418: "LAK",
	422: "LBP",
	426: "LSL",
	430: "LRD",
	434: "LYD",
	446: "MOP",
	454: "MWK",
	458: "MYR",
	462: "MVR",
	480: "MUR",
	484: "MXN",
	496: "MNT",
	498: "MDL",
	504: "MAD",
	512: "OMR",
	516: "NAD",
	524: "NPR",
	532: "ANG",
	533: "AWG",
	548: "VUV",
	554: "NZD",
	558: "NIO",
	566: "NGN",
	578: "NOK",
	586: "PKR",
	590: "PAB",
	598: "PGK",
	600: "PYG",
	604: "PEN",
	608: "PHP",
	634: "QAR",
	643: "RUB",
	646: "RWF",
	654: "SHP",
	682: "SAR",
	690: "SCR",
	694: "SLL",
	702: "SGD",
	704: "VND",
	706: "SOS",
	710: "ZAR",
	728: "SSP",
	748: "SZL",
	752: "SEK",
	756: "CHF",
	760: "SYP",
	764: "THB",
	776: "TOP",
	780: "TTD",
	784: "AED",
	788: "TND",
	800: "UGX",
	807: "MKD",
	818: "EGP",
	826: "GBP",
	834: "TZS",
	840: "USD",
	858: "UYU",
	860: "UZS",
	882: "WST",
	886: "YER",
	901: "TWD",
	927: "UYW",
	928: "VES",
	929: "MRU",
	930: "STN",
	931: "CUC",
	932: "ZWL",
	933: "BYN",
	934: "TMT",
	936: "GHS",
	938: "SDG",
	940: "UYI",
	941: "RSD",
	943: "MZN",
	944: "AZN",
	946: "RON",
	947: "CHE",
	948: "CHW",
	949: "TRY",
	950: "XAF",
	951: "XCD",
	952: "XOF",
	953: "XPF",
	955: "XBA",
	956: "XBB",
	957: "XBC",
	958: "XBD",
	959: "XAU",
	960: "XDR",
	961: "XAG",
	962: "XPT",
	963: "XTS",
	964: "XPD",
	965: "XUA",
	967: "ZMW",
	968: "SRD",
	969: "MGA",
	970: "COU",
	971: "AFN",
	972: "TJS",
	973: "AOA",
	975: "BGN",
	976: "CDF",
	977: "BAM",
	978: "EUR",
	979: "MXV",
	980: "UAH",
	981: "GEL",
	984: "BOV",
	985: "PLN",
	986: "BRL",
	990: "CLF",
	994: "XSU",
	997: "USN",
}

var Currency_value = map[string]int32{
	"XXX": 0,
	"ALL": 8,
	"DZD": 12,
	"ARS": 32,
	"AUD": 36,
	"BSD": 44,
	"BHD": 48,
	"BDT": 50,
	"AMD": 51,
	"BBD": 52,
	"BMD": 60,
	"BTN": 64,
	"BOB": 68,
	"BWP": 72,
	"BZD": 84,
	"SBD": 90,
	"BND": 96,
	"MMK": 104,
	"BIF": 108,
	"KHR": 116,
	"CAD": 124,
	"CVE": 132,
	"KYD": 136,
	"LKR": 144,
	"CLP": 152,
	"CNY": 156,
	"COP": 170,
	"KMF": 174,
	"CRC": 188,
	"HRK": 191,
	"CUP": 192,
	"CZK": 203,
	"DKK": 208,
	"DOP": 214,
	"SVC": 222,
	"ETB": 230,
	"ERN": 232,
	"FKP": 238,
	"FJD": 242,
	"DJF": 262,
	"GMD": 270,
	"GIP": 292,
	"GTQ": 320,
	"GNF": 324,
	"GYD": 328,
	"HTG": 332,
	"HNL": 340,
	"HKD": 344,
	"HUF": 348,
	"ISK": 352,
	"INR": 356,
	"IDR": 360,
	"IRR": 364,
	"IQD": 368,
	"ILS": 376,
	"JMD": 388,
	"JPY": 392,
	"KZT": 398,
	"JOD": 400,
	"KES": 404,
	"KPW": 408,
	"KRW": 410,
	"KWD": 414,
	"KGS": 417,
	"LAK": 418,
	"LBP": 422,
	"LSL": 426,
	"LRD": 430,
	"LYD": 434,
	"MOP": 446,
	"MWK": 454,
	"MYR": 458,
	"MVR": 462,
	"MUR": 480,
	"MXN": 484,
	"MNT": 496,
	"MDL": 498,
	"MAD": 504,
	"OMR": 512,
	"NAD": 516,
	"NPR": 524,
	"ANG": 532,
	"AWG": 533,
	"VUV": 548,
	"NZD": 554,
	"NIO": 558,
	"NGN": 566,
	"NOK": 578,
	"PKR": 586,
	"PAB": 590,
	"PGK": 598,
	"PYG": 600,
	"PEN": 604,
	"PHP": 608,
	"QAR": 634,
	"RUB": 643,
	"RWF": 646,
	"SHP": 654,
	"SAR": 682,
	"SCR": 690,
	"SLL": 694,
	"SGD": 702,
	"VND": 704,
	"SOS": 706,
	"ZAR": 710,
	"SSP": 728,
	"SZL": 748,
	"SEK": 752,
	"CHF": 756,
	"SYP": 760,
	"THB": 764,
	"TOP": 776,
	"TTD": 780,
	"AED": 784,
	"TND": 788,
	"UGX": 800,
	"MKD": 807,
	"EGP": 818,
	"GBP": 826,
	"TZS": 834,
	"USD": 840,
	"UYU": 858,
	"UZS": 860,
	"WST": 882,
	"YER": 886,
	"TWD": 901,
	"UYW": 927,
	"VES": 928,
	"MRU": 929,
	"STN": 930,
	"CUC": 931,
	"ZWL": 932,
	"BYN": 933,
	"TMT": 934,
	"GHS": 936,
	"SDG": 938,
	"UYI": 940,
	"RSD": 941,
	"MZN": 943,
	"AZN": 944,
	"RON": 946,
	"CHE": 947,
	"CHW": 948,
	"TRY": 949,
	"XAF": 950,
	"XCD": 951,
	"XOF": 952,
	"XPF": 953,
	"XBA": 955,
	"XBB": 956,
	"XBC": 957,
	"XBD": 958,
	"XAU": 959,
	"XDR": 960,
	"XAG": 961,
	"XPT": 962,
	"XTS": 963,
	"XPD": 964,
	"XUA": 965,
	"ZMW": 967,
	"SRD": 968,
	"MGA": 969,
	"COU": 970,
	"AFN": 971,
	"TJS": 972,
	"AOA": 973,
	"BGN": 975,
	"CDF": 976,
	"BAM": 977,
	"EUR": 978,
	"MXV": 979,
	"UAH": 980,
	"GEL": 981,
	"BOV": 984,
	"PLN": 985,
	"BRL": 986,
	"CLF": 990,
	"XSU": 994,
	"USN": 997,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}

func (Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_375cec4c1b9d61c6, []int{0}
}

func init() {
	proto.RegisterEnum("pgde.format.Currency", Currency_name, Currency_value)
}

func init() { proto.RegisterFile("pgde/format/currency.proto", fileDescriptor_375cec4c1b9d61c6) }

var fileDescriptor_375cec4c1b9d61c6 = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd5, 0x4f, 0x88, 0x55, 0x55,
	0x1c, 0x07, 0xf0, 0x6c, 0xd2, 0x91, 0xa9, 0xe0, 0x8b, 0xb4, 0x72, 0xd5, 0xa2, 0x55, 0xa4, 0x13,
	0xd9, 0x22, 0xa2, 0x45, 0xf7, 0xbe, 0x73, 0xff, 0xcc, 0xdc, 0x7b, 0xcf, 0xbd, 0x9e, 0x73, 0xef,
	0xbb, 0xef, 0xde, 0x55, 0x3a, 0x3e, 0x47, 0x29, 0x7d, 0xc3, 0xe3, 0x09, 0x16, 0x2d, 0x82, 0x4c,
	0x5c, 0x88, 0xb8, 0x30, 0x90, 0x88, 0x52, 0x33, 0x8b, 0xa1, 0x2c, 0xa6, 0xb2, 0x7f, 0xa6, 0xd3,
	0x64, 0x36, 0xa9, 0x99, 0x95, 0xc9, 0x24, 0x22, 0x12, 0x16, 0x2e, 0x24, 0x44, 0x22, 0x06, 0x69,
	0x11, 0xef, 0xf7, 0x7d, 0x8b, 0x36, 0xc3, 0x87, 0xef, 0x39, 0xdf, 0x73, 0xe6, 0x77, 0xe0, 0x71,
	0x47, 0x56, 0x4e, 0x4d, 0x6e, 0x68, 0x8f, 0x6e, 0xec, 0x74, 0xb7, 0xac, 0xeb, 0x8d, 0x4e, 0x6c,
	0xeb, 0x76, 0xdb, 0x5b, 0x27, 0x9e, 0x5f, 0x3d, 0xd5, 0xed, 0xf4, 0x3a, 0x2b, 0xee, 0xed, 0xaf,
	0xad, 0xe6, 0xda, 0xca, 0x07, 0x64, 0xe3, 0x0b, 0xed, 0x6e, 0x47, 0xfe, 0x70, 0xcb, 0xc3, 0x77,
	0xee, 0x1f, 0x59, 0xde, 0x18, 0xb4, 0x56, 0x0c, 0x8f, 0x0c, 0xb5, 0x5a, 0x2d, 0xdc, 0xd5, 0x87,
	0x13, 0xc7, 0x58, 0xde, 0x87, 0xaa, 0x15, 0xee, 0x93, 0xc4, 0x58, 0x3c, 0x28, 0x28, 0x14, 0x1e,
	0xea, 0xc3, 0xb5, 0x0a, 0x8f, 0x08, 0x42, 0x85, 0x47, 0x05, 0x2a, 0xc7, 0x63, 0xb2, 0x27, 0x51,
	0x58, 0x23, 0x89, 0xab, 0xf0, 0xb8, 0x20, 0x51, 0x78, 0x4a, 0x90, 0x6b, 0x3c, 0x2d, 0x48, 0x5d,
	0x28, 0x41, 0x99, 0x21, 0x14, 0xd4, 0x0a, 0x79, 0x1f, 0xd6, 0x55, 0xa8, 0x25, 0xd1, 0x0a, 0xcf,
	0xf4, 0x91, 0x24, 0x11, 0x36, 0x49, 0x32, 0xe6, 0xe3, 0xb9, 0x3e, 0xa2, 0xd0, 0xa0, 0xd7, 0x47,
	0xc3, 0x51, 0x78, 0x71, 0xc5, 0xf2, 0x91, 0xa1, 0x46, 0xd3, 0xc3, 0x8e, 0x25, 0x7d, 0x45, 0x95,
	0xc2, 0x2e, 0x51, 0x1c, 0x19, 0xec, 0x15, 0x35, 0xe2, 0x0c, 0xfb, 0x29, 0x5d, 0xe1, 0x75, 0x2a,
	0xcd, 0x30, 0xcd, 0x46, 0xe2, 0xe3, 0x28, 0x33, 0xd3, 0xc0, 0x71, 0x51, 0x68, 0x22, 0x9c, 0x64,
	0x56, 0x64, 0x98, 0xa5, 0xea, 0x08, 0x67, 0x45, 0x2a, 0x8a, 0x70, 0x81, 0x4a, 0x33, 0x5c, 0x12,
	0xd9, 0x66, 0x03, 0x57, 0x45, 0x5e, 0xee, 0xe2, 0x4f, 0xca, 0x68, 0xdc, 0x10, 0xf9, 0x51, 0x86,
	0xbf, 0xa8, 0x71, 0x85, 0xdb, 0xec, 0x8e, 0xfb, 0xd8, 0x79, 0x77, 0x5f, 0x41, 0xa2, 0xb0, 0x87,
	0x1a, 0xcb, 0x70, 0x98, 0xca, 0xd7, 0x62, 0x96, 0xd2, 0x3e, 0x4e, 0x51, 0x95, 0xc2, 0xbc, 0x28,
	0xcc, 0x03, 0x9c, 0xa3, 0x74, 0x8c, 0x8b, 0x54, 0xa4, 0xb0, 0x40, 0x15, 0x3e, 0xae, 0x88, 0xc6,
	0x6c, 0x84, 0x6b, 0x94, 0x36, 0xb8, 0x4e, 0x29, 0x83, 0x1b, 0x94, 0x31, 0xb8, 0x49, 0xad, 0x55,
	0xb8, 0x45, 0xc5, 0x16, 0x8b, 0xa2, 0xf1, 0x44, 0x61, 0xc7, 0x90, 0x28, 0xab, 0xb0, 0x4b, 0x14,
	0xd5, 0x39, 0xf6, 0x30, 0x4b, 0x15, 0xf6, 0x32, 0xf3, 0x2c, 0xf6, 0x51, 0x59, 0x89, 0xfd, 0x94,
	0x29, 0xf1, 0x1a, 0x55, 0x2a, 0xbc, 0x41, 0x05, 0x16, 0x07, 0x45, 0xb1, 0x13, 0xe1, 0x10, 0xe5,
	0x66, 0x38, 0x42, 0xd9, 0x18, 0xd3, 0x94, 0x51, 0x38, 0x4a, 0x55, 0x0a, 0x33, 0xa2, 0x24, 0xcd,
	0x70, 0x82, 0x2a, 0x23, 0x9c, 0xa6, 0x2a, 0x83, 0x33, 0x54, 0xd3, 0xe0, 0x3c, 0x55, 0x18, 0x5c,
	0xa3, 0x5a, 0x1a, 0xd7, 0x29, 0x9d, 0xe3, 0x16, 0xa5, 0x62, 0xdc, 0xa6, 0x1c, 0x85, 0x45, 0x51,
	0x9a, 0x18, 0xbc, 0x74, 0x4f, 0x5f, 0xda, 0x51, 0xd8, 0x41, 0x65, 0x06, 0xbb, 0x45, 0x8e, 0x0e,
	0xb0, 0x8f, 0x2a, 0x03, 0xbc, 0x2a, 0x6a, 0x16, 0x4d, 0x1c, 0xe6, 0xbe, 0x5a, 0x61, 0x9a, 0x1a,
	0x4b, 0x71, 0x94, 0x0a, 0x34, 0x8e, 0x51, 0x69, 0x84, 0x39, 0x51, 0x16, 0x19, 0x9c, 0xa1, 0x1c,
	0x17, 0xe7, 0xa9, 0x20, 0xc2, 0x25, 0xaa, 0x0a, 0xb0, 0x40, 0x79, 0x1a, 0x57, 0xa8, 0x30, 0xc3,
	0x35, 0xd1, 0x5a, 0xc7, 0xe0, 0x8e, 0xc8, 0x14, 0x2e, 0x5e, 0x5e, 0x2a, 0x2a, 0x7d, 0xec, 0x14,
	0xd9, 0x30, 0xc3, 0x1e, 0xca, 0x31, 0x98, 0xa6, 0x1a, 0x06, 0x33, 0x54, 0x1c, 0xe3, 0x18, 0x15,
	0x28, 0x9c, 0x10, 0x35, 0xb5, 0xc2, 0x2c, 0xb3, 0xd4, 0x62, 0x4e, 0x54, 0x3b, 0x06, 0xa7, 0x99,
	0xd9, 0x0c, 0x0b, 0x54, 0x1d, 0xe3, 0x26, 0xe5, 0x45, 0xb8, 0x25, 0x6a, 0x84, 0x3e, 0xfe, 0x66,
	0x56, 0x65, 0x58, 0x14, 0xe5, 0xa1, 0x8b, 0x7f, 0xa9, 0x34, 0xc3, 0xae, 0x65, 0xa2, 0x5c, 0x61,
	0xb7, 0xc8, 0xf1, 0x14, 0xf6, 0x32, 0xd3, 0x0a, 0xfb, 0x44, 0x45, 0xd0, 0xc2, 0x01, 0x51, 0x12,
	0x29, 0xbc, 0x23, 0xf2, 0x82, 0x0c, 0x33, 0xa2, 0xc0, 0xcd, 0xf0, 0x05, 0x1b, 0xb5, 0xc5, 0x1c,
	0x1b, 0x56, 0x61, 0x9e, 0xaa, 0x0a, 0x5c, 0xa6, 0x6a, 0x8b, 0x2b, 0xa2, 0xd2, 0xe6, 0xb8, 0x2d,
	0xaa, 0x3c, 0x83, 0x7f, 0xd8, 0x2d, 0x15, 0x5e, 0x19, 0x66, 0xa3, 0xc4, 0x9b, 0xa2, 0xa6, 0x67,
	0x71, 0x40, 0x94, 0x98, 0x02, 0x07, 0x45, 0x36, 0xd7, 0x38, 0x34, 0xcc, 0xdf, 0x79, 0x03, 0x6f,
	0x89, 0xea, 0x32, 0xc6, 0x61, 0x91, 0x5b, 0x69, 0xbc, 0x2d, 0xca, 0x93, 0x1c, 0x47, 0x44, 0x41,
	0x68, 0xf1, 0x2e, 0xbb, 0x2a, 0xc0, 0xf4, 0xe0, 0x8e, 0x31, 0xbc, 0x27, 0x32, 0x56, 0xe1, 0x7d,
	0xde, 0x51, 0x6b, 0x7c, 0x20, 0x72, 0x6a, 0x8d, 0x0f, 0xb9, 0x9a, 0x6a, 0xcc, 0xf0, 0xb6, 0xd0,
	0xc3, 0x47, 0x03, 0x95, 0xf8, 0x98, 0x77, 0x98, 0x0a, 0x9f, 0x88, 0x5a, 0x8e, 0x8f, 0x63, 0x54,
	0x43, 0xe1, 0x53, 0x2a, 0xf5, 0xf1, 0x19, 0x95, 0xf9, 0xf8, 0x9c, 0x72, 0x1d, 0x7c, 0x39, 0x90,
	0x8b, 0xe3, 0x03, 0x35, 0xf0, 0xd5, 0x40, 0x0a, 0x27, 0x06, 0xe7, 0x15, 0x38, 0x49, 0x29, 0x83,
	0xd9, 0x41, 0x16, 0xe0, 0xeb, 0xc1, 0x79, 0x39, 0xe6, 0xa8, 0xdc, 0xe2, 0x9b, 0x41, 0xa6, 0x70,
	0x8a, 0x2a, 0x1c, 0x7c, 0xcb, 0x77, 0x49, 0x4a, 0x7c, 0xc7, 0xc9, 0x8d, 0xc2, 0x3c, 0xa7, 0x0c,
	0x1c, 0x7c, 0xcf, 0x39, 0xd2, 0x02, 0x67, 0x38, 0xaf, 0xaf, 0x71, 0x96, 0x13, 0x8d, 0x5b, 0x9c,
	0x63, 0x96, 0x3a, 0xf8, 0x81, 0x6f, 0x1a, 0x68, 0xfc, 0xc8, 0x86, 0xf2, 0x71, 0x81, 0x99, 0x93,
	0xe0, 0x27, 0x91, 0x57, 0x18, 0xfc, 0xcc, 0x93, 0x5b, 0x4d, 0xfc, 0xc2, 0xd7, 0x75, 0x42, 0x5c,
	0xe4, 0xdb, 0x7b, 0x31, 0x7e, 0x65, 0x23, 0x6d, 0x62, 0x41, 0x94, 0xc5, 0x1a, 0xbf, 0x31, 0x33,
	0x31, 0x2e, 0xf3, 0xe4, 0xd8, 0xc7, 0x55, 0xfe, 0xf7, 0xb6, 0xc0, 0xef, 0x3c, 0xc5, 0x6a, 0xfc,
	0x31, 0xbc, 0x72, 0xe9, 0xfc, 0xf5, 0xd9, 0xa1, 0x25, 0xee, 0x93, 0xf5, 0x13, 0x93, 0x9b, 0x7b,
	0x9b, 0xb6, 0xad, 0x5f, 0x3d, 0xd1, 0xd9, 0x32, 0xda, 0xde, 0x32, 0xb5, 0xb9, 0xdb, 0xde, 0xd8,
	0xd9, 0x3e, 0x2a, 0x9f, 0xc6, 0x89, 0x55, 0x93, 0xed, 0xad, 0xab, 0x36, 0xac, 0xeb, 0xf6, 0x56,
	0xb5, 0xb7, 0xf7, 0x46, 0xa7, 0x9e, 0x9d, 0x1c, 0xfd, 0xdf, 0xa7, 0x76, 0xfd, 0x32, 0xd9, 0xb4,
	0xe6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xd2, 0x72, 0x72, 0x80, 0x07, 0x00, 0x00,
}
