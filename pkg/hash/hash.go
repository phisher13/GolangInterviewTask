package hash

import (
	hashids "github.com/speps/go-hashids/v2"
)


func HashFunction(data string) string {
	hd := hashids.NewData()
	hd.Salt = data
	hd.MinLength = 1
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{52, 1, 32, 993})
	
	return e
}