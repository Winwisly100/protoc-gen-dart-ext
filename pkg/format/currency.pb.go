// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/format/currency.proto

package format // import "github.com/empirefox/protoc-gen-dart-ext/pkg/format"

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

// ISO 4217, publish date: 2018-08-29
type CurrencyV1 int32

const (
	CurrencyV1_XXX CurrencyV1 = 0
	CurrencyV1_ALL CurrencyV1 = 8
	CurrencyV1_DZD CurrencyV1 = 12
	CurrencyV1_ARS CurrencyV1 = 32
	CurrencyV1_AUD CurrencyV1 = 36
	CurrencyV1_BSD CurrencyV1 = 44
	CurrencyV1_BHD CurrencyV1 = 48
	CurrencyV1_BDT CurrencyV1 = 50
	CurrencyV1_AMD CurrencyV1 = 51
	CurrencyV1_BBD CurrencyV1 = 52
	CurrencyV1_BMD CurrencyV1 = 60
	CurrencyV1_BTN CurrencyV1 = 64
	CurrencyV1_BOB CurrencyV1 = 68
	CurrencyV1_BWP CurrencyV1 = 72
	CurrencyV1_BZD CurrencyV1 = 84
	CurrencyV1_SBD CurrencyV1 = 90
	CurrencyV1_BND CurrencyV1 = 96
	CurrencyV1_MMK CurrencyV1 = 104
	CurrencyV1_BIF CurrencyV1 = 108
	CurrencyV1_KHR CurrencyV1 = 116
	CurrencyV1_CAD CurrencyV1 = 124
	CurrencyV1_CVE CurrencyV1 = 132
	CurrencyV1_KYD CurrencyV1 = 136
	CurrencyV1_LKR CurrencyV1 = 144
	CurrencyV1_CLP CurrencyV1 = 152
	CurrencyV1_CNY CurrencyV1 = 156
	CurrencyV1_COP CurrencyV1 = 170
	CurrencyV1_KMF CurrencyV1 = 174
	CurrencyV1_CRC CurrencyV1 = 188
	CurrencyV1_HRK CurrencyV1 = 191
	CurrencyV1_CUP CurrencyV1 = 192
	CurrencyV1_CZK CurrencyV1 = 203
	CurrencyV1_DKK CurrencyV1 = 208
	CurrencyV1_DOP CurrencyV1 = 214
	CurrencyV1_SVC CurrencyV1 = 222
	CurrencyV1_ETB CurrencyV1 = 230
	CurrencyV1_ERN CurrencyV1 = 232
	CurrencyV1_FKP CurrencyV1 = 238
	CurrencyV1_FJD CurrencyV1 = 242
	CurrencyV1_DJF CurrencyV1 = 262
	CurrencyV1_GMD CurrencyV1 = 270
	CurrencyV1_GIP CurrencyV1 = 292
	CurrencyV1_GTQ CurrencyV1 = 320
	CurrencyV1_GNF CurrencyV1 = 324
	CurrencyV1_GYD CurrencyV1 = 328
	CurrencyV1_HTG CurrencyV1 = 332
	CurrencyV1_HNL CurrencyV1 = 340
	CurrencyV1_HKD CurrencyV1 = 344
	CurrencyV1_HUF CurrencyV1 = 348
	CurrencyV1_ISK CurrencyV1 = 352
	CurrencyV1_INR CurrencyV1 = 356
	CurrencyV1_IDR CurrencyV1 = 360
	CurrencyV1_IRR CurrencyV1 = 364
	CurrencyV1_IQD CurrencyV1 = 368
	CurrencyV1_ILS CurrencyV1 = 376
	CurrencyV1_JMD CurrencyV1 = 388
	CurrencyV1_JPY CurrencyV1 = 392
	CurrencyV1_KZT CurrencyV1 = 398
	CurrencyV1_JOD CurrencyV1 = 400
	CurrencyV1_KES CurrencyV1 = 404
	CurrencyV1_KPW CurrencyV1 = 408
	CurrencyV1_KRW CurrencyV1 = 410
	CurrencyV1_KWD CurrencyV1 = 414
	CurrencyV1_KGS CurrencyV1 = 417
	CurrencyV1_LAK CurrencyV1 = 418
	CurrencyV1_LBP CurrencyV1 = 422
	CurrencyV1_LSL CurrencyV1 = 426
	CurrencyV1_LRD CurrencyV1 = 430
	CurrencyV1_LYD CurrencyV1 = 434
	CurrencyV1_MOP CurrencyV1 = 446
	CurrencyV1_MWK CurrencyV1 = 454
	CurrencyV1_MYR CurrencyV1 = 458
	CurrencyV1_MVR CurrencyV1 = 462
	CurrencyV1_MUR CurrencyV1 = 480
	CurrencyV1_MXN CurrencyV1 = 484
	CurrencyV1_MNT CurrencyV1 = 496
	CurrencyV1_MDL CurrencyV1 = 498
	CurrencyV1_MAD CurrencyV1 = 504
	CurrencyV1_OMR CurrencyV1 = 512
	CurrencyV1_NAD CurrencyV1 = 516
	CurrencyV1_NPR CurrencyV1 = 524
	CurrencyV1_ANG CurrencyV1 = 532
	CurrencyV1_AWG CurrencyV1 = 533
	CurrencyV1_VUV CurrencyV1 = 548
	CurrencyV1_NZD CurrencyV1 = 554
	CurrencyV1_NIO CurrencyV1 = 558
	CurrencyV1_NGN CurrencyV1 = 566
	CurrencyV1_NOK CurrencyV1 = 578
	CurrencyV1_PKR CurrencyV1 = 586
	CurrencyV1_PAB CurrencyV1 = 590
	CurrencyV1_PGK CurrencyV1 = 598
	CurrencyV1_PYG CurrencyV1 = 600
	CurrencyV1_PEN CurrencyV1 = 604
	CurrencyV1_PHP CurrencyV1 = 608
	CurrencyV1_QAR CurrencyV1 = 634
	CurrencyV1_RUB CurrencyV1 = 643
	CurrencyV1_RWF CurrencyV1 = 646
	CurrencyV1_SHP CurrencyV1 = 654
	CurrencyV1_SAR CurrencyV1 = 682
	CurrencyV1_SCR CurrencyV1 = 690
	CurrencyV1_SLL CurrencyV1 = 694
	CurrencyV1_SGD CurrencyV1 = 702
	CurrencyV1_VND CurrencyV1 = 704
	CurrencyV1_SOS CurrencyV1 = 706
	CurrencyV1_ZAR CurrencyV1 = 710
	CurrencyV1_SSP CurrencyV1 = 728
	CurrencyV1_SZL CurrencyV1 = 748
	CurrencyV1_SEK CurrencyV1 = 752
	CurrencyV1_CHF CurrencyV1 = 756
	CurrencyV1_SYP CurrencyV1 = 760
	CurrencyV1_THB CurrencyV1 = 764
	CurrencyV1_TOP CurrencyV1 = 776
	CurrencyV1_TTD CurrencyV1 = 780
	CurrencyV1_AED CurrencyV1 = 784
	CurrencyV1_TND CurrencyV1 = 788
	CurrencyV1_UGX CurrencyV1 = 800
	CurrencyV1_MKD CurrencyV1 = 807
	CurrencyV1_EGP CurrencyV1 = 818
	CurrencyV1_GBP CurrencyV1 = 826
	CurrencyV1_TZS CurrencyV1 = 834
	CurrencyV1_USD CurrencyV1 = 840
	CurrencyV1_UYU CurrencyV1 = 858
	CurrencyV1_UZS CurrencyV1 = 860
	CurrencyV1_WST CurrencyV1 = 882
	CurrencyV1_YER CurrencyV1 = 886
	CurrencyV1_TWD CurrencyV1 = 901
	CurrencyV1_UYW CurrencyV1 = 927
	CurrencyV1_VES CurrencyV1 = 928
	CurrencyV1_MRU CurrencyV1 = 929
	CurrencyV1_STN CurrencyV1 = 930
	CurrencyV1_CUC CurrencyV1 = 931
	CurrencyV1_ZWL CurrencyV1 = 932
	CurrencyV1_BYN CurrencyV1 = 933
	CurrencyV1_TMT CurrencyV1 = 934
	CurrencyV1_GHS CurrencyV1 = 936
	CurrencyV1_SDG CurrencyV1 = 938
	CurrencyV1_UYI CurrencyV1 = 940
	CurrencyV1_RSD CurrencyV1 = 941
	CurrencyV1_MZN CurrencyV1 = 943
	CurrencyV1_AZN CurrencyV1 = 944
	CurrencyV1_RON CurrencyV1 = 946
	CurrencyV1_CHE CurrencyV1 = 947
	CurrencyV1_CHW CurrencyV1 = 948
	CurrencyV1_TRY CurrencyV1 = 949
	CurrencyV1_XAF CurrencyV1 = 950
	CurrencyV1_XCD CurrencyV1 = 951
	CurrencyV1_XOF CurrencyV1 = 952
	CurrencyV1_XPF CurrencyV1 = 953
	CurrencyV1_XBA CurrencyV1 = 955
	CurrencyV1_XBB CurrencyV1 = 956
	CurrencyV1_XBC CurrencyV1 = 957
	CurrencyV1_XBD CurrencyV1 = 958
	CurrencyV1_XAU CurrencyV1 = 959
	CurrencyV1_XDR CurrencyV1 = 960
	CurrencyV1_XAG CurrencyV1 = 961
	CurrencyV1_XPT CurrencyV1 = 962
	CurrencyV1_XTS CurrencyV1 = 963
	CurrencyV1_XPD CurrencyV1 = 964
	CurrencyV1_XUA CurrencyV1 = 965
	CurrencyV1_ZMW CurrencyV1 = 967
	CurrencyV1_SRD CurrencyV1 = 968
	CurrencyV1_MGA CurrencyV1 = 969
	CurrencyV1_COU CurrencyV1 = 970
	CurrencyV1_AFN CurrencyV1 = 971
	CurrencyV1_TJS CurrencyV1 = 972
	CurrencyV1_AOA CurrencyV1 = 973
	CurrencyV1_BGN CurrencyV1 = 975
	CurrencyV1_CDF CurrencyV1 = 976
	CurrencyV1_BAM CurrencyV1 = 977
	CurrencyV1_EUR CurrencyV1 = 978
	CurrencyV1_MXV CurrencyV1 = 979
	CurrencyV1_UAH CurrencyV1 = 980
	CurrencyV1_GEL CurrencyV1 = 981
	CurrencyV1_BOV CurrencyV1 = 984
	CurrencyV1_PLN CurrencyV1 = 985
	CurrencyV1_BRL CurrencyV1 = 986
	CurrencyV1_CLF CurrencyV1 = 990
	CurrencyV1_XSU CurrencyV1 = 994
	CurrencyV1_USN CurrencyV1 = 997
)

var CurrencyV1_name = map[int32]string{
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
var CurrencyV1_value = map[string]int32{
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

func (x CurrencyV1) String() string {
	return proto.EnumName(CurrencyV1_name, int32(x))
}
func (CurrencyV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_currency_906048d5d19f7b4c, []int{0}
}

func init() {
	proto.RegisterEnum("format.CurrencyV1", CurrencyV1_name, CurrencyV1_value)
}

func init() {
	proto.RegisterFile("protos/format/currency.proto", fileDescriptor_currency_906048d5d19f7b4c)
}

var fileDescriptor_currency_906048d5d19f7b4c = []byte{
	// 1020 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x2c, 0xd5, 0x4f, 0x68, 0x1c, 0x55,
	0x1c, 0x07, 0x70, 0xdb, 0x6a, 0x12, 0x82, 0xc2, 0x97, 0x1c, 0x3c, 0x14, 0x0f, 0x1e, 0x3c, 0x89,
	0x69, 0x5a, 0xa3, 0x37, 0x0f, 0xce, 0xec, 0xdb, 0x99, 0x4d, 0x66, 0x76, 0x66, 0xf2, 0x66, 0x66,
	0x67, 0x77, 0x4f, 0xb6, 0x31, 0x4d, 0x8b, 0xcd, 0x1f, 0xd6, 0x2d, 0x54, 0xf0, 0x20, 0x18, 0x25,
	0x87, 0x12, 0x7a, 0x88, 0x50, 0x44, 0xb4, 0xad, 0xb1, 0x4a, 0xd0, 0x28, 0x51, 0xe3, 0xbf, 0xda,
	0xc6, 0x58, 0x6b, 0x4c, 0x6b, 0xad, 0x5a, 0x4b, 0x2c, 0xa1, 0x04, 0x89, 0xd2, 0x43, 0x91, 0x10,
	0x44, 0x82, 0x88, 0xf8, 0xf6, 0xf7, 0x9d, 0xcb, 0xf2, 0xe1, 0x3b, 0xef, 0xf7, 0xde, 0xbc, 0xb7,
	0xbf, 0xb7, 0xdb, 0xf9, 0xc0, 0x78, 0x63, 0xac, 0x39, 0xf6, 0x6c, 0xcf, 0xc1, 0xb1, 0xc6, 0xc8,
	0xfe, 0x66, 0xcf, 0xe0, 0xd1, 0x46, 0x63, 0x68, 0x74, 0xf0, 0xb9, 0x3d, 0x12, 0x77, 0xb5, 0x31,
	0xde, 0x7d, 0x7f, 0x3e, 0xea, 0xc8, 0xbe, 0xbd, 0xa3, 0xf2, 0xc1, 0xe7, 0x0f, 0xff, 0x77, 0x5f,
	0x67, 0x67, 0x21, 0x2f, 0xa9, 0xec, 0xeb, 0x6a, 0xef, 0xdc, 0x55, 0xad, 0x56, 0x71, 0x57, 0x0b,
	0x96, 0xef, 0xa3, 0xa3, 0x05, 0x55, 0x57, 0xb8, 0x57, 0x12, 0x1d, 0xe3, 0x41, 0x41, 0xaa, 0xf0,
	0x50, 0x0b, 0x76, 0xac, 0xf0, 0x88, 0xa0, 0xa4, 0xb0, 0x57, 0xa0, 0x12, 0x3c, 0x2a, 0x63, 0xca,
	0x0a, 0xbd, 0x92, 0xd8, 0x0a, 0x8f, 0x09, 0x4c, 0xf2, 0x84, 0x20, 0x09, 0xf0, 0xa4, 0x20, 0xb4,
	0xa1, 0x04, 0x59, 0x84, 0x92, 0xc0, 0xac, 0x95, 0xb4, 0x10, 0x9b, 0xaa, 0xba, 0x24, 0x81, 0xc2,
	0x53, 0x2d, 0x94, 0xcb, 0x1e, 0x0e, 0x49, 0xd2, 0xe7, 0xe0, 0x48, 0x0b, 0x5e, 0x49, 0xa3, 0xd9,
	0x42, 0xc1, 0x52, 0x78, 0xbe, 0xab, 0xc3, 0xa0, 0x52, 0xc4, 0xc4, 0x8e, 0x96, 0xbc, 0x9a, 0xc2,
	0xa4, 0xc8, 0xf7, 0x34, 0x4e, 0x88, 0x0a, 0x7e, 0x84, 0x93, 0x54, 0x50, 0xc3, 0x6b, 0x54, 0x18,
	0x61, 0x96, 0x15, 0x65, 0x07, 0x73, 0xcc, 0x74, 0x01, 0xe7, 0x44, 0x25, 0xed, 0xe1, 0x02, 0xb3,
	0x34, 0xc2, 0x22, 0x55, 0xf7, 0x70, 0x59, 0xa4, 0x3c, 0x0f, 0xd7, 0x28, 0x33, 0xcb, 0x0d, 0x51,
	0x5c, 0x29, 0xe0, 0x96, 0xa8, 0x98, 0xd8, 0xf8, 0x83, 0xd2, 0x01, 0x6e, 0x8b, 0x1c, 0x2f, 0xc2,
	0x9f, 0x54, 0xbf, 0xc2, 0x16, 0x6b, 0xfb, 0x1d, 0xbc, 0xbc, 0xb3, 0x25, 0xd7, 0x1c, 0xd0, 0x14,
	0xd5, 0x17, 0x61, 0x86, 0x4a, 0x06, 0xb0, 0x48, 0x05, 0x0e, 0x2e, 0x52, 0x66, 0x6f, 0xcb, 0xa2,
	0x52, 0xe2, 0xe2, 0x0a, 0x15, 0xf8, 0xb8, 0x4e, 0x79, 0x0a, 0xab, 0x54, 0xea, 0x60, 0x4d, 0xd4,
	0x17, 0x7b, 0x58, 0xa7, 0x02, 0x8d, 0x0d, 0x4a, 0x69, 0xdc, 0xa6, 0xb4, 0xc6, 0x1d, 0x6a, 0x40,
	0x61, 0x93, 0xf2, 0x63, 0x6c, 0x8b, 0xfa, 0xcd, 0x5b, 0x4d, 0xec, 0x12, 0x45, 0x35, 0x4c, 0x8a,
	0xbc, 0x7a, 0x82, 0x29, 0x66, 0xa1, 0xc2, 0x09, 0x66, 0xc5, 0x18, 0xd3, 0x54, 0x94, 0xe1, 0x24,
	0xa5, 0x33, 0xbc, 0x4a, 0x65, 0x0a, 0xaf, 0x53, 0x6e, 0x8c, 0xd3, 0x22, 0xdf, 0xf2, 0x70, 0x86,
	0xb2, 0x23, 0x9c, 0xa5, 0x62, 0x1f, 0xb3, 0x94, 0x56, 0x98, 0xa3, 0xcc, 0x7e, 0xe7, 0x45, 0x65,
	0x73, 0xce, 0xe7, 0xa9, 0xcc, 0xc3, 0x25, 0xaa, 0xa6, 0xb1, 0x42, 0x55, 0x34, 0xae, 0x52, 0xa9,
	0xc6, 0x3a, 0x55, 0x0d, 0xb0, 0x41, 0x05, 0x09, 0x36, 0x29, 0xe5, 0x63, 0x8b, 0x32, 0x8d, 0xb3,
	0x2d, 0x0a, 0xcb, 0x1a, 0x2f, 0xdc, 0xdd, 0x52, 0x60, 0xb2, 0x09, 0x2a, 0xd2, 0x38, 0x2e, 0xb2,
	0x02, 0x17, 0xd3, 0x54, 0xe6, 0xe2, 0x15, 0x51, 0x25, 0xad, 0x60, 0x86, 0xe3, 0x4c, 0xd3, 0xce,
	0x52, 0x7d, 0x21, 0xe6, 0x28, 0x37, 0xc0, 0x02, 0x15, 0x7a, 0x58, 0x12, 0x45, 0xa6, 0x13, 0x57,
	0x28, 0xcb, 0xc6, 0x55, 0xca, 0xf5, 0x70, 0x83, 0xaa, 0xb9, 0x58, 0xa5, 0x8a, 0x01, 0xd6, 0xa8,
	0x52, 0x84, 0x75, 0xd1, 0x80, 0xa5, 0xf1, 0x8f, 0x48, 0xa7, 0x36, 0x5e, 0xbc, 0x47, 0x94, 0x99,
	0xce, 0x11, 0xc5, 0x66, 0xdc, 0x14, 0x65, 0xc6, 0xcd, 0x52, 0x05, 0x8d, 0x79, 0xca, 0xdc, 0xe4,
	0x05, 0xca, 0x55, 0x38, 0x2f, 0xaa, 0x98, 0x6b, 0xb5, 0xc8, 0x2c, 0x8c, 0xb1, 0x24, 0xaa, 0x9b,
	0xda, 0x4b, 0xcc, 0xe2, 0x08, 0xab, 0x54, 0xdd, 0xc7, 0x1d, 0xaa, 0xe8, 0x61, 0x53, 0x54, 0x28,
	0x39, 0xf8, 0x8b, 0x59, 0x2d, 0xc2, 0xb6, 0x28, 0x29, 0xd9, 0xf8, 0x97, 0x32, 0xdf, 0xd1, 0x64,
	0x9b, 0x28, 0x51, 0x38, 0x2e, 0xb2, 0x8a, 0xa6, 0x4b, 0x98, 0x99, 0x75, 0xa7, 0x45, 0xa9, 0x5b,
	0xc5, 0x29, 0x51, 0xd9, 0x74, 0xec, 0xdb, 0xa2, 0xa2, 0x1b, 0x61, 0x5e, 0xe4, 0x9a, 0x8e, 0xf8,
	0x9c, 0x15, 0x75, 0xf3, 0x7e, 0xac, 0x30, 0xbf, 0x36, 0xcb, 0x54, 0x2d, 0xc5, 0x4d, 0xca, 0x3c,
	0x5d, 0x13, 0x65, 0x71, 0x82, 0x2d, 0x51, 0xad, 0xa8, 0xf1, 0x37, 0x6b, 0x4d, 0xd7, 0xbd, 0xd4,
	0xce, 0x8a, 0x0c, 0x6f, 0x88, 0x2a, 0xa6, 0x4f, 0x4f, 0x89, 0xca, 0x3a, 0xc5, 0x69, 0x51, 0x6c,
	0x7e, 0x91, 0xce, 0xb4, 0xf3, 0x9e, 0x17, 0xf0, 0xa6, 0xa8, 0x9e, 0xf9, 0x98, 0x11, 0xd9, 0xb5,
	0x00, 0x6f, 0x89, 0x92, 0x72, 0x82, 0xb3, 0x22, 0xb7, 0x14, 0xe3, 0x1d, 0xd6, 0x2a, 0x17, 0xb3,
	0xf9, 0x1a, 0x7d, 0x78, 0x57, 0xa4, 0xcd, 0x9b, 0xbe, 0xc7, 0x35, 0xea, 0x01, 0xde, 0x17, 0x59,
	0x46, 0x1f, 0xf0, 0x69, 0x18, 0x60, 0x9e, 0xab, 0x95, 0x8a, 0xf8, 0x30, 0x57, 0x86, 0x8f, 0xb8,
	0x86, 0xae, 0xe1, 0x63, 0x51, 0xd5, 0x72, 0xb0, 0x40, 0x15, 0x14, 0x3e, 0xa1, 0x42, 0x07, 0x9f,
	0x52, 0x91, 0x83, 0xcf, 0x28, 0xdb, 0xc2, 0x17, 0xb9, 0x6c, 0x9c, 0xcb, 0x55, 0xc0, 0x97, 0xb9,
	0xcc, 0x77, 0x9e, 0xcf, 0x97, 0xe2, 0x02, 0x65, 0xee, 0xfe, 0x62, 0x9e, 0xb9, 0xf8, 0x2a, 0x9f,
	0x2f, 0xc1, 0x12, 0x95, 0xc4, 0xf8, 0x3a, 0xcf, 0x14, 0x2e, 0x52, 0xa9, 0x85, 0x6f, 0x78, 0x2e,
	0xe5, 0x0c, 0xdf, 0x72, 0xe7, 0xe6, 0x86, 0x2e, 0x73, 0x97, 0xae, 0x85, 0xef, 0xb8, 0x8f, 0x30,
	0xc5, 0x0a, 0xf7, 0xeb, 0x04, 0xb8, 0xcc, 0x1d, 0xf5, 0xc7, 0xb8, 0xc2, 0x2c, 0xb4, 0xf0, 0x3d,
	0xcf, 0xd4, 0xdc, 0x8f, 0x1f, 0x58, 0xa1, 0x1c, 0x5c, 0x63, 0x66, 0x95, 0xf1, 0xa3, 0xa8, 0x68,
	0xee, 0xef, 0x4f, 0x9c, 0xb9, 0x5a, 0xc1, 0xcf, 0x3c, 0x5d, 0xab, 0x84, 0xeb, 0x3c, 0xfb, 0xa2,
	0x8f, 0x5f, 0x58, 0x11, 0x56, 0xb0, 0x2a, 0x8a, 0xfc, 0x00, 0xbf, 0x32, 0xd3, 0x3e, 0x6e, 0x72,
	0x66, 0xdf, 0xc1, 0x2d, 0xbe, 0x7d, 0x9c, 0xe2, 0x37, 0xce, 0x12, 0x07, 0xf8, 0xbd, 0x7d, 0x77,
	0xfb, 0xfc, 0xc6, 0xe2, 0xae, 0x9d, 0x1d, 0x3b, 0xec, 0xc7, 0xeb, 0xbd, 0xc3, 0x87, 0x9b, 0x87,
	0x8e, 0x1e, 0xd8, 0x33, 0x38, 0x36, 0xd2, 0x33, 0x34, 0x32, 0x7e, 0xb8, 0x31, 0x74, 0x70, 0xec,
	0x58, 0x8f, 0xfc, 0x3d, 0x0e, 0x76, 0x0f, 0x0f, 0x8d, 0x76, 0x3f, 0xbd, 0xbf, 0xd1, 0xec, 0x1e,
	0x3a, 0xd6, 0xec, 0x19, 0x7f, 0x66, 0x38, 0xff, 0x9b, 0x3d, 0xd0, 0x26, 0xcf, 0x7b, 0xff, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x41, 0x1d, 0x4b, 0x7e, 0x07, 0x00, 0x00,
}