package util

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func ConvertByteToUint32(buf []byte, offset int) uint32 {
	var v uint32
	v += uint32(buf[offset]) << 24
	v += uint32(buf[offset+1]) << 16
	v += uint32(buf[offset+2]) << 8
	v += uint32(buf[offset+3])
	return v
}

func SliceContains(slice []interface{}, data interface{}) bool {
	for _, v := range slice {
		if v == data {
			return true
		}
	}
	return false
}
