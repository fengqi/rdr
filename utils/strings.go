package utils

func KeyPrefixDistinct(key string) string {
	k := ""
	var buf []rune
	for _, c := range key {
		if c >= 48 && c <= 57 { //48 == "0" 57 == "9"
			buf = append(buf, '0')
		} else {
			if len(buf) > 0 {
				k += "000"
			}
			k += string(c)
			buf = buf[:0]
		}
	}
	if len(buf) > 0 {
		k += "000"
	}
	return k
}
