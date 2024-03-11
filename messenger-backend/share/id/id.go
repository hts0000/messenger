package id

import (
	"strconv"
)

type UID uint64

func String2UID(uid string) (UID, error) {
	id, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		return 0, err
	}
	return UID(id), nil
}

func (id UID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
