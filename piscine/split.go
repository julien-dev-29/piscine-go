package piscine

func Split(s, sep string) []string {
	var res []string
	if sep == "" {
		return nil
	}
	start := 0

	for i := 0; i <= len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			res = append(res, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	res = append(res, s[start:])
	return res
}
