package caspertransaction

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"strconv"
)

func byteArrayCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

//reverseBytes endian reverse
func reverseBytes(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//reverseHexToBytes decode a hex string to an byte array,then change the endian
func reverseHexToBytes(hexVar string) ([]byte, error) {
	if len(hexVar)%2 == 1 {
		return nil, errors.New("Invalid TxHash!")
	}
	ret, err := hex.DecodeString(hexVar)
	if err != nil {
		return nil, err
	}
	return reverseBytes(ret), nil
}

//reverseBytesToHex change the endian of the input byte array then encode it to hex string
func reverseBytesToHex(bytesVar []byte) string {
	return hex.EncodeToString(reverseBytes(bytesVar))
}

// uint16 转换为小端字节数组
func uint16ToLittleEndianBytes(data uint16) []byte {
	tmp := [2]byte{}
	binary.LittleEndian.PutUint16(tmp[:], data)
	return tmp[:]
}

// 小端数组转无符号16位整数
func littleEndianBytesToUint16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

// uint32 转换为小端字节数组
func uint32ToLittleEndianBytes(data uint32) []byte {
	tmp := [4]byte{}
	binary.LittleEndian.PutUint32(tmp[:], data)
	return tmp[:]
}
func uint32ToLittleEndianBytesShortVersion(data uint32) []byte {
	tmp := [4]byte{}
	binary.LittleEndian.PutUint32(tmp[:], data)
	return tmp[:]
}

func littleEndianBytesToUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func uint64ToLittleEndianBytes(data uint64) []byte {
	tmp := [8]byte{}
	binary.LittleEndian.PutUint64(tmp[:], data)
	return tmp[:]
}

func littleEndianBytesToUint64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func uintToShortByte(n uint64) []byte {
	hexStr := strconv.FormatUint(n, 16)

	//log.Info("hexStr:",hexStr)
	if len(hexStr)%2 > 0 {
		hexStr = "0" + hexStr
	}
	var bytesData []byte
	strLen := len(hexStr)
	for index := 0; index < strLen; index += 2 {
		num, _ := strconv.ParseUint(hexStr[index:index+2], 16, 64)
		bytesData = append(bytesData, byte(num))
	}
	bytesData = reverseBytes(bytesData)
	bytesData = append([]byte{byte(len(bytesData))}, bytesData...)

	return bytesData
}
