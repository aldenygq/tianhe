package pkg


func Int64ToPointInt64(i int64) *int64 {
	ptrValue := new(int64)
    *ptrValue = i
	return  ptrValue
}