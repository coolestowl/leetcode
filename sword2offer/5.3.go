package sword2offer

func FirstUniqChar(s string) byte {
	counter := make(map[byte]int)
	bts := []byte(s)
	for _, b := range bts {
		counter[b]++
	}

	result := byte(' ')
	for _, b := range bts {
		if counter[b] == 1 {
			result = b
			break
		}
	}
	return result
}
