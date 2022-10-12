package shortner

type URLService struct {
	elements string
	COUNTER  int
	LTOS     map[string]int
	STOL     map[int]string
}

func (s *URLService) LongToShort(url string) (string, int, bool) {
	var shorturl string
	var existing bool
	if s.LTOS == nil {
		s.LTOS = map[string]int{}
		s.STOL = map[int]string{}
		s.elements = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s.COUNTER = 1000000000
	}
	if id, ok := s.LTOS[url]; ok {
		shorturl = s.STOL[id]
		existing = ok
	} else {
		shorturl = s.base10ToBase62(s.COUNTER)
		s.COUNTER++
		s.LTOS[url] = s.COUNTER
		s.STOL[s.COUNTER] = shorturl
	}
	return "http://tiny.url/" + shorturl, s.COUNTER, existing
}

func (s *URLService) ShortToLong(url string) string {
	url = url[len("http://tiny.url/"):]
	var n int = s.base62ToBase10(url)
	return s.STOL[n]
}

func (s *URLService) base62ToBase10(str string) int {
	var n int = 0
	for i := 0; i < len(str); i++ {
		n = n*62 + s.convert(str[i])
	}
	return n
}

func (s *URLService) convert(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 10)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 36
	}
	return -1
}

func (s *URLService) base10ToBase62(n int) string {
	var sb string
	for n != 0 {
		sb = string(s.elements[n%62]) + sb
		n /= 62
	}
	for len(sb) != 7 {
		sb = string('0') + sb
	}
	return string(sb)
}
