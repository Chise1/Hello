package hello

import quoteV3 "rsc.io/quote/v3"

func Hello() string {
	return quoteV3.HelloV3()
}
func HelloV2() string {
	return "Hello,I'm back."
}
func HelloV3() string {
	return "Hello,V3."
}
func Proverb() string {
	return quoteV3.Concurrency()
}
