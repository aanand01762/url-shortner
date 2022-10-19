package shortner

type URLService struct {
	Elements  string
	COUNTER   int
	LONGTOID  map[string]int
	IDTOSMALL map[int]string
	IDTOLONG  map[int]string
}

func (s *URLService) LongToShort(url string) (string, int, bool) {
	var shorturl string
	var existing bool

	//Check if url already exists, update existing flag
	if id, ok := s.LONGTOID[url]; ok {
		shorturl = s.IDTOSMALL[id]
		existing = ok
	} else {

		//if url does not exist convert id or counter
		//of url to base62 encoded text
		s.COUNTER++
		shorturl = s.base10ToBase62(s.COUNTER)
		s.LONGTOID[url] = s.COUNTER
		s.IDTOLONG[s.COUNTER] = url
		s.IDTOSMALL[s.COUNTER] = shorturl
	}

	//Append the base62 encoded text to tinyurl
	return "http://tiny.url/" + shorturl, s.COUNTER, existing
}

func (s *URLService) ShortToLong(url string) string {

	//Return long url from IDTOSMALL hashmap using
	//encoded text as key
	url = url[len("http://tiny.url/"):]
	var n int = s.base62ToBase10(url)
	return s.IDTOLONG[n]
}

func (s *URLService) base62ToBase10(str string) int {

	//Convert base62 format to interger value in base10 format
	var n int = 0
	for i := 0; i < len(str); i++ {
		n = n*62 + s.convert(str[i])
	}
	return n
}

func (s *URLService) convert(c byte) int {

	//Return base10 integer representation of a char in base62
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

	//Convert base62 format to interger value in base10 format
	var sb string
	for n != 0 {
		sb = string(s.Elements[n%62]) + sb
		n /= 62
	}
	for len(sb) != 7 {
		sb = string('0') + sb
	}
	return string(sb)
}
