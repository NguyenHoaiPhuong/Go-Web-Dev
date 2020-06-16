package main

import (
	"fmt"

	"github.com/ttacon/libphonenumber"
)

func main() {
	var (
		countryCode = new(int32)
		number      = new(uint64)
	)
	*countryCode = 84
	*number = 868024671
	phone := &libphonenumber.PhoneNumber{
		CountryCode:    countryCode,
		NationalNumber: number,
	}
	res := libphonenumber.IsValidNumber(phone)
	fmt.Printf("%+v\n", res)
}
