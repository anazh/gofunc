package slice

//SRemove 从切片中删除元素
func SRemove(slice []string, elem string) []string {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return nil
	} else if sliceLen == 1 {
		if slice[0] == elem {
			return nil
		}
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return SRemove(slice, elem) //遍历删除
		}
	}
	return slice
}
