package myqoute

import "rsc.io/quote"

func GetQuote() string {
	return quote.Glass()
}
