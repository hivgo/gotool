package util

import "strconv"

func AtoBool(str string) (bool, error) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}

func AtoUInt16(str string) (uint16, error) {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint16(num), nil
}

func AtoInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

func AtoUInt32(str string) (uint32, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

func AtoUInt(str string) (uint, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}

func AtoInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(num), nil
}
