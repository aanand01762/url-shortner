package shortner

type URLService struct {
	elements  string
	COUNTER   int
	LONGTOID  map[string]int
	IDTOSMALL map[int]string
}

func (s *URLService) LongToShort(url string) (string, int, bool) {
	var shorturl string
	var existing bool

	//Check if instance struct type does not exits,
	//then only initialise an instance which means its first entry
	if s.LONGTOID == nil {
		s.LONGTOID = map[string]int{}
		s.IDTOSMALL = map[int]string{}
		s.elements = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s.COUNTER = 1000000000
	}

	//Check if url already exists, update existing flag
	if id, ok := s.LONGTOID[url]; ok {
		shorturl = s.IDTOSMALL[id]
		existing = ok
	} else {

		//if url does not exist convert id or counter
		//of url to base62 encoded text
		shorturl = s.base10ToBase62(s.COUNTER)
		s.COUNTER++
		s.LONGTOID[url] = s.COUNTER
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
	return s.IDTOSMALL[n]
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
		sb = string(s.elements[n%62]) + sb
		n /= 62
	}
	for len(sb) != 7 {
		sb = string('0') + sb
	}
	return string(sb)
}
