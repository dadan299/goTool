package datahandle

import (
	"strconv"
)

func JoinInt (elems []int,sep string) string {
	switch len(elems) {
	case 0:
		return ""
	}

	gather := make([]any,0)

	for _, elem := range elems {
		z := any(elem)
		gather = append(gather,z)
	}

	return join(gather,sep)

}

func JoinInt64(elems []int64,sep string) string {
	switch len(elems) {
	case 0:
		return ""
	}

	gather := make([]any,0)

	for _, elem := range elems {
		z := any(elem)
		gather = append(gather,z)
	}

	return join(gather,sep)
}


func join (elems []any,sep string) string {

	resp := ""

	for i:=0;i<len(elems);i++ {

		str := ""
		v := elems[i]

		switch v.(type) {
		case string:
			str = v.(string)
		case int:
			str = strconv.Itoa(v.(int))
		case int64:
			str = strconv.FormatInt(v.(int64), 10)
		}

		if i == 0 {
			resp += str
		}else {
			resp += sep + str
		}
	}

	return resp
}