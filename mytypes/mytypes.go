package mytypes

type MyInt int

// Twice multiplies its receiver by 2 and returns the result
func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyStr string

// StrLen returns the length of a string
func (str MyStr) StrLen() int {
	return len(str)
}
