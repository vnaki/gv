package validator

import "regexp"

type Message struct {
	ok bool
	value string
}

/**
 * create a new message
 */
func NewMessage (ok bool) *Message {
	return &Message{ok:ok, value:"Please set the prompt message!"}
}

/**
 * set the message content
 */
func (m *Message) Message (v string) {
	m.value = v
}

type Validator struct {
	Messages []*Message
}

/**
 * create a default validator
 */
func NewValidator () *Validator {
	return &Validator{Messages:[]*Message{}}
}

/**
 * Requirements are not empty strings
 */
func (vd *Validator) Require (str string) *Message {
	return vd.MatchBool(str != "")
}

/**
 * string min & max length
 */
func (vd *Validator) Size (v string, minSize int, maxSize int) *Message {
	l := len(v)
	return vd.MatchBool(l >= minSize && l <= maxSize)
}

/**
 * string min length
 */
func (vd *Validator) MinSize (v string, minSize int) *Message {
	return vd.MatchBool(len(v) >= minSize)
}

/**
 * string max length
 */
func (vd *Validator) MaxSize (v string, maxSize int) *Message {
	return vd.MatchBool(len(v) <= maxSize)
}

/**
 * number min value
 */
func (vd *Validator) MinInt (num int, min int) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxInt (num int, max int) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number min value
 */
func (vd *Validator) MinInt8 (num int8, min int8) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxInt8 (num int8, max int8) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number max value
 */
func (vd *Validator) MinInt16 (num int16, min int16) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxInt16 (num int16, max int16) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number min value
 */
func (vd *Validator) MinInt32 (num int32, min int32) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxInt32 (num int32, max int32) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number min value
 */
func (vd *Validator) MinInt64 (num int64, min int64) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxInt64 (num int64, max int64) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number min value
 */
func (vd *Validator) MinFloat32 (num float32, min float32) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxFloat32 (num float32, max float32) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * number min value
 */
func (vd *Validator) MinFloat64 (num float64, min float64) *Message {
	return vd.MatchBool(num >= min)
}

/**
 * number max value
 */
func (vd *Validator) MaxFloat64 (num float64, max float64) *Message {
	return vd.MatchBool(num <= max)
}

/**
 * validate email
 */
func (vd *Validator) Email (email string) *Message {
	return vd.Match(`^[[:alnum:]]+(-[[:alnum:]]+)*@[[:alnum:]]+(\.[[:alpha:]]+){1,2}$`, email)
}

/**
 * validate mobile
 */
func (vd *Validator) Mobile(mobile string) *Message {
	return vd.Match(`^1[3578]\d{9}$`, mobile)
}

/**
 * validate zip code
 */
func (vd *Validator) ZipCode(code string) *Message {
	return vd.Match(`^[1-9]\d{5}$`, code)
}

/**
 * validate ipv4 address
 */
func (vd *Validator) IPv4 (ipv4 string) *Message {
	return vd.Match(`^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$`, ipv4)
}

/**
 * validate domain
 * eg: domain.com or domain.com.cn or www.domain.com.cn or www.domain-test.com
 */
func (vd *Validator) Domain(domain string) *Message {
	return vd.Match(`^[[:alnum:]]+(-[[:alnum:]]+)*(\.[[:alpha:]]{2,3}){1,2}$`, domain)
}

/**
 * validate date
 */
func (vd *Validator) Date(date string) *Message {
	return vd.Match(`^[12]\d{3}[-/](0[1-9]|1[0-2])[-/](0[1-9]|[1-2][0-9]|30|31)$`, date)
}

/**
 * validate date time
 */
func (vd *Validator) Datetime(datetime string) *Message {
	return vd.Match(`^[12]\d{3}[-/](0[1-9]|1[0-2])[-/](0[1-9]|[1-2][0-9]|30|31)(\s{1,3}(23|22|21|20|[0-1][0-9]):[0-5][0-9]:[0-5][0-9])$`, datetime)
}

/**
 * validate letter(A-Z or a-z)
 */
func (vd *Validator) Letter (str string) *Message {
	return vd.Match(`^[[:alpha:]]+$`, str)
}

/**
 * validate number or letter (0-9 or A-Z or a-z)
 */
func (vd *Validator) Alpha (str string) *Message {
	return vd.Match(`^[[:alnum:]]+$`, str)
}

/**
 * validate complex string, by number(0-9) or letter (A-Za-z) or underline(_)
 */
func (vd *Validator) AlphaDash (str string) *Message {
	return vd.Match(`^[[:alnum:]]+(_[[:alnum:]]+)*$`, str)
}

/**
 * validate string & number
 */
func (vd *Validator) Number(num string) *Message {
	return vd.Match(`^-?\d+$`, num)
}

/**
 * validate string & negative number (<0)
 */
func (vd *Validator) Negative(num string) *Message {
	return vd.Match(`^-[1-9]\d*$`, num)
}

/**
 * validate string & positive number (>0)
 */
func (vd *Validator) Positive(num string) *Message {
	return vd.Match(`^[1-9]\d*$`, num)
}

/**
 * validate string & float number
 */
func (vd *Validator) Float(num string) *Message {
	return vd.Match(`^-?\d+\.\d+$`, num)
}

/**
 * Determine whether a string is in an array
 */
func (vd *Validator) RangeString (str string, array []string) *Message {
	ok := false

	for _, v := range array {
		if v == str {
			ok = true
			break
		}
	}

	return vd.MatchBool(ok)
}

/**
 * Determine whether a number is in an array
 */
func (vd *Validator) RangeInt (num int, array []int) *Message {
	ok := false

	for _, v := range array {
		if v == num {
			ok = true
			break
		}
	}

	return vd.MatchBool(ok)
}

/**
 * Determine whether a float number is in an array
 */
func (vd *Validator) RangeFloat (num float64, array []float64) *Message {
	ok := false

	for _, v := range array {
		if v == num {
			ok = true
			break
		}
	}

	return vd.MatchBool(ok)
}

/**
 * regular expression matching
 */
func (vd *Validator) Match (pattern string, str string) *Message {
	m := NewMessage(regexp.MustCompile(pattern).MatchString(str))

	vd.Messages = append(vd.Messages, m)

	return m
}

/**
 * regular expression matching
 */
func (vd *Validator) MatchBool (bool bool) *Message {
	m := NewMessage(bool)

	vd.Messages = append(vd.Messages, m)

	return m
}

/**
 * verification results
 */
func (vd *Validator) Validate () (bool, string) {
	m := vd.GetMessages()
	return len(m) == 0, m[0]
}

/**
 * return message slice
 */
func (vd *Validator) GetMessages () []string {
	m := make([]string, 0)

	for _, msg := range vd.Messages {
		if false == msg.ok {
			m = append(m, msg.value)
		}
	}

	return m
}