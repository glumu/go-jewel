package utils

func GenHeaders(requestId string, headers [][2]string) map[string]string {
	head := map[string]string{
		"Content-Type": `application/json;charset=UTF-8`,
		"X-Request-Id": requestId,
	}
	if headers != nil {
		for _, header := range headers {
			if header[1] != "" {
				head[header[0]] = header[1]
			}
		}
	}
	return head
}
