package regexptest

import (
	"fmt"
	"regexp"
)

func TestRegexp()  {

	//sampleRegex := regexp.MustCompile("[.]")
	//
	//match := sampleRegex.Match([]byte("."))
	//fmt.Println(match)
	//r1:= "abc"
	//r2:= "xyz"
	//sampleRegexAnd := regexp.MustCompile(r1+r2)
	//fmt.Println(sampleRegexAnd.Match([]byte("abc")))
	//fmt.Println(sampleRegexAnd.Match([]byte("abcxyz")))
	//sampleregexOr := regexp.MustCompile(r1+"|" +r2)
	//fmt.Println(sampleregexOr.Match([]byte("abc")))
	//fmt.Println(sampleregexOr.Match([]byte("xyz")))
	//fmt.Println("=======")
	//sampleRegexABC := regexp.MustCompile("^abc$")
	//fmt.Println(sampleRegexABC.Match([]byte("abcd")))
	//fmt.Println(sampleRegexABC.Match([]byte("1abcd")))
	//fmt.Println(sampleRegexABC.Match([]byte("abcd123")))
	//fmt.Println(sampleRegexABC.Match([]byte("abc")))
	fmt.Println("=======")
	//sampleRegexABCD := regexp.MustCompile("(?i)abc")
	//sampleRegexABCD := regexp.MustCompile(".")
	//sampleRegexABCD := regexp.MustCompile("(?s).")  // (?s) 与 "\n"也匹配 (?i) 不区分大小写
	//fmt.Printf("%t\n", sampleRegexABCD.Match([]byte("a")))
	//fmt.Printf("%t\n", sampleRegexABCD.Match([]byte("\n")))
	//fmt.Printf("%t\n", sampleRegexABCD.Match([]byte("ABCD")))
	//sampleRegex := regexp.MustCompile(`(\w+):([0-9]{1,3})`)
	//sampleRegex := regexp.MustCompile(`(?P<name>\w+):(?P<age>[0-9]\d{1,3})`)
	//input := "The names are John:21, simon:23, and Mike:19"
	//result := sampleRegex.ReplaceAllString(input, "$age:$name")
	//result := sampleRegex.Match([]byte("http"))
	//fmt.Println(result)
	//sampleRegex := regexp.MustCompile("(a+?)(a*)") // [aaaaaaa a aaaaaa], aaaaaaa 为整体，a 为 (a+?)，aaaaaa为 (a*)
	//result := sampleRegex.FindStringSubmatch("aaaaaaa")
	//fmt.Println(result, len(result))
	//sampleRegex = regexp.MustCompile("(a*?)(a*)") // [aaaaaaa  aaaaaaa]
	//result = sampleRegex.FindStringSubmatch("aaaaaaa")
	//fmt.Println(result, len(result))
	//sampleRegex = regexp.MustCompile("(a*?)(a*?)")
	//result = sampleRegex.FindStringSubmatch("aaaaaaa")
	//fmt.Println(result, len(result))
	sampleRegex := regexp.MustCompile(`[+\-]?(?:(?:\d+)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?`)
	fmt.Printf("%t\n",sampleRegex.MatchString("1.2"))
	fmt.Printf("%t\n",sampleRegex.MatchString(".12"))
	fmt.Printf("%t\n",sampleRegex.MatchString("12"))
	fmt.Printf("%t\n",sampleRegex.MatchString("12."))
	fmt.Printf("%t\n",sampleRegex.MatchString("+1.2"))
	fmt.Printf("%t\n",sampleRegex.MatchString("-1.2"))
	fmt.Printf("%t\n",sampleRegex.MatchString("12e3"))
}
