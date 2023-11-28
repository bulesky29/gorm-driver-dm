/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_919 struct{}

var Dm_build_920 = &dm_build_919{}

func (Dm_build_922 *dm_build_919) Dm_build_921(dm_build_923 []byte, dm_build_924 int, dm_build_925 byte) int {
	dm_build_923[dm_build_924] = dm_build_925
	return 1
}

func (Dm_build_927 *dm_build_919) Dm_build_926(dm_build_928 []byte, dm_build_929 int, dm_build_930 int8) int {
	dm_build_928[dm_build_929] = byte(dm_build_930)
	return 1
}

func (Dm_build_932 *dm_build_919) Dm_build_931(dm_build_933 []byte, dm_build_934 int, dm_build_935 int16) int {
	dm_build_933[dm_build_934] = byte(dm_build_935)
	dm_build_934++
	dm_build_933[dm_build_934] = byte(dm_build_935 >> 8)
	return 2
}

func (Dm_build_937 *dm_build_919) Dm_build_936(dm_build_938 []byte, dm_build_939 int, dm_build_940 int32) int {
	dm_build_938[dm_build_939] = byte(dm_build_940)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 8)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 16)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 24)
	dm_build_939++
	return 4
}

func (Dm_build_942 *dm_build_919) Dm_build_941(dm_build_943 []byte, dm_build_944 int, dm_build_945 int64) int {
	dm_build_943[dm_build_944] = byte(dm_build_945)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 8)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 16)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 24)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 32)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 40)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 48)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 56)
	return 8
}

func (Dm_build_947 *dm_build_919) Dm_build_946(dm_build_948 []byte, dm_build_949 int, dm_build_950 float32) int {
	return Dm_build_947.Dm_build_966(dm_build_948, dm_build_949, math.Float32bits(dm_build_950))
}

func (Dm_build_952 *dm_build_919) Dm_build_951(dm_build_953 []byte, dm_build_954 int, dm_build_955 float64) int {
	return Dm_build_952.Dm_build_971(dm_build_953, dm_build_954, math.Float64bits(dm_build_955))
}

func (Dm_build_957 *dm_build_919) Dm_build_956(dm_build_958 []byte, dm_build_959 int, dm_build_960 uint8) int {
	dm_build_958[dm_build_959] = byte(dm_build_960)
	return 1
}

func (Dm_build_962 *dm_build_919) Dm_build_961(dm_build_963 []byte, dm_build_964 int, dm_build_965 uint16) int {
	dm_build_963[dm_build_964] = byte(dm_build_965)
	dm_build_964++
	dm_build_963[dm_build_964] = byte(dm_build_965 >> 8)
	return 2
}

func (Dm_build_967 *dm_build_919) Dm_build_966(dm_build_968 []byte, dm_build_969 int, dm_build_970 uint32) int {
	dm_build_968[dm_build_969] = byte(dm_build_970)
	dm_build_969++
	dm_build_968[dm_build_969] = byte(dm_build_970 >> 8)
	dm_build_969++
	dm_build_968[dm_build_969] = byte(dm_build_970 >> 16)
	dm_build_969++
	dm_build_968[dm_build_969] = byte(dm_build_970 >> 24)
	return 3
}

func (Dm_build_972 *dm_build_919) Dm_build_971(dm_build_973 []byte, dm_build_974 int, dm_build_975 uint64) int {
	dm_build_973[dm_build_974] = byte(dm_build_975)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 8)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 16)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 24)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 32)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 40)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 48)
	dm_build_974++
	dm_build_973[dm_build_974] = byte(dm_build_975 >> 56)
	return 3
}

func (Dm_build_977 *dm_build_919) Dm_build_976(dm_build_978 []byte, dm_build_979 int, dm_build_980 []byte, dm_build_981 int, dm_build_982 int) int {
	copy(dm_build_978[dm_build_979:dm_build_979+dm_build_982], dm_build_980[dm_build_981:dm_build_981+dm_build_982])
	return dm_build_982
}

func (Dm_build_984 *dm_build_919) Dm_build_983(dm_build_985 []byte, dm_build_986 int, dm_build_987 []byte, dm_build_988 int, dm_build_989 int) int {
	dm_build_986 += Dm_build_984.Dm_build_966(dm_build_985, dm_build_986, uint32(dm_build_989))
	return 4 + Dm_build_984.Dm_build_976(dm_build_985, dm_build_986, dm_build_987, dm_build_988, dm_build_989)
}

func (Dm_build_991 *dm_build_919) Dm_build_990(dm_build_992 []byte, dm_build_993 int, dm_build_994 []byte, dm_build_995 int, dm_build_996 int) int {
	dm_build_993 += Dm_build_991.Dm_build_961(dm_build_992, dm_build_993, uint16(dm_build_996))
	return 2 + Dm_build_991.Dm_build_976(dm_build_992, dm_build_993, dm_build_994, dm_build_995, dm_build_996)
}

func (Dm_build_998 *dm_build_919) Dm_build_997(dm_build_999 []byte, dm_build_1000 int, dm_build_1001 string, dm_build_1002 string, dm_build_1003 *DmConnection) int {
	dm_build_1004 := Dm_build_998.Dm_build_1136(dm_build_1001, dm_build_1002, dm_build_1003)
	dm_build_1000 += Dm_build_998.Dm_build_966(dm_build_999, dm_build_1000, uint32(len(dm_build_1004)))
	return 4 + Dm_build_998.Dm_build_976(dm_build_999, dm_build_1000, dm_build_1004, 0, len(dm_build_1004))
}

func (Dm_build_1006 *dm_build_919) Dm_build_1005(dm_build_1007 []byte, dm_build_1008 int, dm_build_1009 string, dm_build_1010 string, dm_build_1011 *DmConnection) int {
	dm_build_1012 := Dm_build_1006.Dm_build_1136(dm_build_1009, dm_build_1010, dm_build_1011)

	dm_build_1008 += Dm_build_1006.Dm_build_961(dm_build_1007, dm_build_1008, uint16(len(dm_build_1012)))
	return 2 + Dm_build_1006.Dm_build_976(dm_build_1007, dm_build_1008, dm_build_1012, 0, len(dm_build_1012))
}

func (Dm_build_1014 *dm_build_919) Dm_build_1013(dm_build_1015 []byte, dm_build_1016 int) byte {
	return dm_build_1015[dm_build_1016]
}

func (Dm_build_1018 *dm_build_919) Dm_build_1017(dm_build_1019 []byte, dm_build_1020 int) int16 {
	var dm_build_1021 int16
	dm_build_1021 = int16(dm_build_1019[dm_build_1020] & 0xff)
	dm_build_1020++
	dm_build_1021 |= int16(dm_build_1019[dm_build_1020]&0xff) << 8
	return dm_build_1021
}

func (Dm_build_1023 *dm_build_919) Dm_build_1022(dm_build_1024 []byte, dm_build_1025 int) int32 {
	var dm_build_1026 int32
	dm_build_1026 = int32(dm_build_1024[dm_build_1025] & 0xff)
	dm_build_1025++
	dm_build_1026 |= int32(dm_build_1024[dm_build_1025]&0xff) << 8
	dm_build_1025++
	dm_build_1026 |= int32(dm_build_1024[dm_build_1025]&0xff) << 16
	dm_build_1025++
	dm_build_1026 |= int32(dm_build_1024[dm_build_1025]&0xff) << 24
	return dm_build_1026
}

func (Dm_build_1028 *dm_build_919) Dm_build_1027(dm_build_1029 []byte, dm_build_1030 int) int64 {
	var dm_build_1031 int64
	dm_build_1031 = int64(dm_build_1029[dm_build_1030] & 0xff)
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 8
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 16
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 24
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 32
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 40
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 48
	dm_build_1030++
	dm_build_1031 |= int64(dm_build_1029[dm_build_1030]&0xff) << 56
	return dm_build_1031
}

func (Dm_build_1033 *dm_build_919) Dm_build_1032(dm_build_1034 []byte, dm_build_1035 int) float32 {
	return math.Float32frombits(Dm_build_1033.Dm_build_1049(dm_build_1034, dm_build_1035))
}

func (Dm_build_1037 *dm_build_919) Dm_build_1036(dm_build_1038 []byte, dm_build_1039 int) float64 {
	return math.Float64frombits(Dm_build_1037.Dm_build_1054(dm_build_1038, dm_build_1039))
}

func (Dm_build_1041 *dm_build_919) Dm_build_1040(dm_build_1042 []byte, dm_build_1043 int) uint8 {
	return uint8(dm_build_1042[dm_build_1043] & 0xff)
}

func (Dm_build_1045 *dm_build_919) Dm_build_1044(dm_build_1046 []byte, dm_build_1047 int) uint16 {
	var dm_build_1048 uint16
	dm_build_1048 = uint16(dm_build_1046[dm_build_1047] & 0xff)
	dm_build_1047++
	dm_build_1048 |= uint16(dm_build_1046[dm_build_1047]&0xff) << 8
	return dm_build_1048
}

func (Dm_build_1050 *dm_build_919) Dm_build_1049(dm_build_1051 []byte, dm_build_1052 int) uint32 {
	var dm_build_1053 uint32
	dm_build_1053 = uint32(dm_build_1051[dm_build_1052] & 0xff)
	dm_build_1052++
	dm_build_1053 |= uint32(dm_build_1051[dm_build_1052]&0xff) << 8
	dm_build_1052++
	dm_build_1053 |= uint32(dm_build_1051[dm_build_1052]&0xff) << 16
	dm_build_1052++
	dm_build_1053 |= uint32(dm_build_1051[dm_build_1052]&0xff) << 24
	return dm_build_1053
}

func (Dm_build_1055 *dm_build_919) Dm_build_1054(dm_build_1056 []byte, dm_build_1057 int) uint64 {
	var dm_build_1058 uint64
	dm_build_1058 = uint64(dm_build_1056[dm_build_1057] & 0xff)
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 8
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 16
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 24
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 32
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 40
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 48
	dm_build_1057++
	dm_build_1058 |= uint64(dm_build_1056[dm_build_1057]&0xff) << 56
	return dm_build_1058
}

func (Dm_build_1060 *dm_build_919) Dm_build_1059(dm_build_1061 []byte, dm_build_1062 int) []byte {
	dm_build_1063 := Dm_build_1060.Dm_build_1049(dm_build_1061, dm_build_1062)

	dm_build_1064 := make([]byte, dm_build_1063)
	copy(dm_build_1064[:int(dm_build_1063)], dm_build_1061[dm_build_1062+4:dm_build_1062+4+int(dm_build_1063)])
	return dm_build_1064
}

func (Dm_build_1066 *dm_build_919) Dm_build_1065(dm_build_1067 []byte, dm_build_1068 int) []byte {
	dm_build_1069 := Dm_build_1066.Dm_build_1044(dm_build_1067, dm_build_1068)

	dm_build_1070 := make([]byte, dm_build_1069)
	copy(dm_build_1070[:int(dm_build_1069)], dm_build_1067[dm_build_1068+2:dm_build_1068+2+int(dm_build_1069)])
	return dm_build_1070
}

func (Dm_build_1072 *dm_build_919) Dm_build_1071(dm_build_1073 []byte, dm_build_1074 int, dm_build_1075 int) []byte {

	dm_build_1076 := make([]byte, dm_build_1075)
	copy(dm_build_1076[:dm_build_1075], dm_build_1073[dm_build_1074:dm_build_1074+dm_build_1075])
	return dm_build_1076
}

func (Dm_build_1078 *dm_build_919) Dm_build_1077(dm_build_1079 []byte, dm_build_1080 int, dm_build_1081 int, dm_build_1082 string, dm_build_1083 *DmConnection) string {
	return Dm_build_1078.Dm_build_1172(dm_build_1079[dm_build_1080:dm_build_1080+dm_build_1081], dm_build_1082, dm_build_1083)
}

func (Dm_build_1085 *dm_build_919) Dm_build_1084(dm_build_1086 []byte, dm_build_1087 int, dm_build_1088 string, dm_build_1089 *DmConnection) string {
	dm_build_1090 := Dm_build_1085.Dm_build_1049(dm_build_1086, dm_build_1087)
	dm_build_1087 += 4
	return Dm_build_1085.Dm_build_1077(dm_build_1086, dm_build_1087, int(dm_build_1090), dm_build_1088, dm_build_1089)
}

func (Dm_build_1092 *dm_build_919) Dm_build_1091(dm_build_1093 []byte, dm_build_1094 int, dm_build_1095 string, dm_build_1096 *DmConnection) string {
	dm_build_1097 := Dm_build_1092.Dm_build_1044(dm_build_1093, dm_build_1094)
	dm_build_1094 += 2
	return Dm_build_1092.Dm_build_1077(dm_build_1093, dm_build_1094, int(dm_build_1097), dm_build_1095, dm_build_1096)
}

func (Dm_build_1099 *dm_build_919) Dm_build_1098(dm_build_1100 byte) []byte {
	return []byte{dm_build_1100}
}

func (Dm_build_1102 *dm_build_919) Dm_build_1101(dm_build_1103 int8) []byte {
	return []byte{byte(dm_build_1103)}
}

func (Dm_build_1105 *dm_build_919) Dm_build_1104(dm_build_1106 int16) []byte {
	return []byte{byte(dm_build_1106), byte(dm_build_1106 >> 8)}
}

func (Dm_build_1108 *dm_build_919) Dm_build_1107(dm_build_1109 int32) []byte {
	return []byte{byte(dm_build_1109), byte(dm_build_1109 >> 8), byte(dm_build_1109 >> 16), byte(dm_build_1109 >> 24)}
}

func (Dm_build_1111 *dm_build_919) Dm_build_1110(dm_build_1112 int64) []byte {
	return []byte{byte(dm_build_1112), byte(dm_build_1112 >> 8), byte(dm_build_1112 >> 16), byte(dm_build_1112 >> 24), byte(dm_build_1112 >> 32),
		byte(dm_build_1112 >> 40), byte(dm_build_1112 >> 48), byte(dm_build_1112 >> 56)}
}

func (Dm_build_1114 *dm_build_919) Dm_build_1113(dm_build_1115 float32) []byte {
	return Dm_build_1114.Dm_build_1125(math.Float32bits(dm_build_1115))
}

func (Dm_build_1117 *dm_build_919) Dm_build_1116(dm_build_1118 float64) []byte {
	return Dm_build_1117.Dm_build_1128(math.Float64bits(dm_build_1118))
}

func (Dm_build_1120 *dm_build_919) Dm_build_1119(dm_build_1121 uint8) []byte {
	return []byte{byte(dm_build_1121)}
}

func (Dm_build_1123 *dm_build_919) Dm_build_1122(dm_build_1124 uint16) []byte {
	return []byte{byte(dm_build_1124), byte(dm_build_1124 >> 8)}
}

func (Dm_build_1126 *dm_build_919) Dm_build_1125(dm_build_1127 uint32) []byte {
	return []byte{byte(dm_build_1127), byte(dm_build_1127 >> 8), byte(dm_build_1127 >> 16), byte(dm_build_1127 >> 24)}
}

func (Dm_build_1129 *dm_build_919) Dm_build_1128(dm_build_1130 uint64) []byte {
	return []byte{byte(dm_build_1130), byte(dm_build_1130 >> 8), byte(dm_build_1130 >> 16), byte(dm_build_1130 >> 24), byte(dm_build_1130 >> 32), byte(dm_build_1130 >> 40), byte(dm_build_1130 >> 48), byte(dm_build_1130 >> 56)}
}

func (Dm_build_1132 *dm_build_919) Dm_build_1131(dm_build_1133 []byte, dm_build_1134 string, dm_build_1135 *DmConnection) []byte {
	if dm_build_1134 == "UTF-8" {
		return dm_build_1133
	}

	if dm_build_1135 == nil {
		if e := dm_build_1177(dm_build_1134); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1133), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1135.encodeBuffer == nil {
		dm_build_1135.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1135.encode = dm_build_1177(dm_build_1135.getServerEncoding())
		dm_build_1135.transformReaderDst = make([]byte, 4096)
		dm_build_1135.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1135.encode; e != nil {

		dm_build_1135.encodeBuffer.Reset()

		n, err := dm_build_1135.encodeBuffer.ReadFrom(
			Dm_build_1191(bytes.NewReader(dm_build_1133), e.NewEncoder(), dm_build_1135.transformReaderDst, dm_build_1135.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1135.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1137 *dm_build_919) Dm_build_1136(dm_build_1138 string, dm_build_1139 string, dm_build_1140 *DmConnection) []byte {
	return Dm_build_1137.Dm_build_1131([]byte(dm_build_1138), dm_build_1139, dm_build_1140)
}

func (Dm_build_1142 *dm_build_919) Dm_build_1141(dm_build_1143 []byte) byte {
	return Dm_build_1142.Dm_build_1013(dm_build_1143, 0)
}

func (Dm_build_1145 *dm_build_919) Dm_build_1144(dm_build_1146 []byte) int16 {
	return Dm_build_1145.Dm_build_1017(dm_build_1146, 0)
}

func (Dm_build_1148 *dm_build_919) Dm_build_1147(dm_build_1149 []byte) int32 {
	return Dm_build_1148.Dm_build_1022(dm_build_1149, 0)
}

func (Dm_build_1151 *dm_build_919) Dm_build_1150(dm_build_1152 []byte) int64 {
	return Dm_build_1151.Dm_build_1027(dm_build_1152, 0)
}

func (Dm_build_1154 *dm_build_919) Dm_build_1153(dm_build_1155 []byte) float32 {
	return Dm_build_1154.Dm_build_1032(dm_build_1155, 0)
}

func (Dm_build_1157 *dm_build_919) Dm_build_1156(dm_build_1158 []byte) float64 {
	return Dm_build_1157.Dm_build_1036(dm_build_1158, 0)
}

func (Dm_build_1160 *dm_build_919) Dm_build_1159(dm_build_1161 []byte) uint8 {
	return Dm_build_1160.Dm_build_1040(dm_build_1161, 0)
}

func (Dm_build_1163 *dm_build_919) Dm_build_1162(dm_build_1164 []byte) uint16 {
	return Dm_build_1163.Dm_build_1044(dm_build_1164, 0)
}

func (Dm_build_1166 *dm_build_919) Dm_build_1165(dm_build_1167 []byte) uint32 {
	return Dm_build_1166.Dm_build_1049(dm_build_1167, 0)
}

func (Dm_build_1169 *dm_build_919) Dm_build_1168(dm_build_1170 []byte, dm_build_1171 string) []byte {
	if dm_build_1171 == "UTF-8" {
		return dm_build_1170
	}

	if e := dm_build_1177(dm_build_1171); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1170), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1173 *dm_build_919) Dm_build_1172(dm_build_1174 []byte, dm_build_1175 string, dm_build_1176 *DmConnection) string {
	return string(Dm_build_1173.Dm_build_1168(dm_build_1174, dm_build_1175))
}

func dm_build_1177(dm_build_1178 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1178); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1179 struct {
	dm_build_1180 io.Reader
	dm_build_1181 transform.Transformer
	dm_build_1182 error

	dm_build_1183                []byte
	dm_build_1184, dm_build_1185 int

	dm_build_1186                []byte
	dm_build_1187, dm_build_1188 int

	dm_build_1189 bool
}

const dm_build_1190 = 4096

func Dm_build_1191(dm_build_1192 io.Reader, dm_build_1193 transform.Transformer, dm_build_1194 []byte, dm_build_1195 []byte) *Dm_build_1179 {
	dm_build_1193.Reset()
	return &Dm_build_1179{
		dm_build_1180: dm_build_1192,
		dm_build_1181: dm_build_1193,
		dm_build_1183: dm_build_1194,
		dm_build_1186: dm_build_1195,
	}
}

func (dm_build_1197 *Dm_build_1179) Read(dm_build_1198 []byte) (int, error) {
	dm_build_1199, dm_build_1200 := 0, error(nil)
	for {

		if dm_build_1197.dm_build_1184 != dm_build_1197.dm_build_1185 {
			dm_build_1199 = copy(dm_build_1198, dm_build_1197.dm_build_1183[dm_build_1197.dm_build_1184:dm_build_1197.dm_build_1185])
			dm_build_1197.dm_build_1184 += dm_build_1199
			if dm_build_1197.dm_build_1184 == dm_build_1197.dm_build_1185 && dm_build_1197.dm_build_1189 {
				return dm_build_1199, dm_build_1197.dm_build_1182
			}
			return dm_build_1199, nil
		} else if dm_build_1197.dm_build_1189 {
			return 0, dm_build_1197.dm_build_1182
		}

		if dm_build_1197.dm_build_1187 != dm_build_1197.dm_build_1188 || dm_build_1197.dm_build_1182 != nil {
			dm_build_1197.dm_build_1184 = 0
			dm_build_1197.dm_build_1185, dm_build_1199, dm_build_1200 = dm_build_1197.dm_build_1181.Transform(dm_build_1197.dm_build_1183, dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1187:dm_build_1197.dm_build_1188], dm_build_1197.dm_build_1182 == io.EOF)
			dm_build_1197.dm_build_1187 += dm_build_1199

			switch {
			case dm_build_1200 == nil:
				if dm_build_1197.dm_build_1187 != dm_build_1197.dm_build_1188 {
					dm_build_1197.dm_build_1182 = nil
				}

				dm_build_1197.dm_build_1189 = dm_build_1197.dm_build_1182 != nil
				continue
			case dm_build_1200 == transform.ErrShortDst && (dm_build_1197.dm_build_1185 != 0 || dm_build_1199 != 0):

				continue
			case dm_build_1200 == transform.ErrShortSrc && dm_build_1197.dm_build_1188-dm_build_1197.dm_build_1187 != len(dm_build_1197.dm_build_1186) && dm_build_1197.dm_build_1182 == nil:

			default:
				dm_build_1197.dm_build_1189 = true

				if dm_build_1197.dm_build_1182 == nil || dm_build_1197.dm_build_1182 == io.EOF {
					dm_build_1197.dm_build_1182 = dm_build_1200
				}
				continue
			}
		}

		if dm_build_1197.dm_build_1187 != 0 {
			dm_build_1197.dm_build_1187, dm_build_1197.dm_build_1188 = 0, copy(dm_build_1197.dm_build_1186, dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1187:dm_build_1197.dm_build_1188])
		}
		dm_build_1199, dm_build_1197.dm_build_1182 = dm_build_1197.dm_build_1180.Read(dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1188:])
		dm_build_1197.dm_build_1188 += dm_build_1199
	}
}
