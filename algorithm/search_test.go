//author xinbing
//time 2018/10/16 10:40
//
package algorithm

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type TB struct {
	s string
}

func TestBinarySearch(t *testing.T) {
	tt := []*TB{
		{
			"a",
		},
		{
			"b",
		},
		{
			"c",
		},
		{
			"d",
		},
		{
			"e",
		},
	}
	value := "a"
	index := BinarySearch(len(tt), func(i int) int {
		return strings.Compare(tt[i].s, value)
	})
	if index == -1 {
		fmt.Println("not found " + value)
		return
	}
	fmt.Println(tt[index].s + " index is:" + strconv.Itoa(index))
}

func TestTT(t *testing.T) {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, email := range emails {
		index := strings.LastIndex(email, "@")
		localName := email[:index]
		localName = strings.Replace(localName, ".", "", -1)
		pIndex := strings.IndexRune(localName, '+')
		if pIndex != -1 {
			localName = localName[:pIndex]
		}
		m[localName+email[index:]] = 0
	}
	fmt.Println(m)
	return len(m)
}
