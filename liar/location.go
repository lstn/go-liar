package liar

type LocationTruth struct {
	Country string
	State string
	City string
	StreetAddress string
	Postal string
}

func (t LocationTruth) Hash() string {
	return ""
}

func (t LocationTruth) Lie() string {
	return ""
}

func (t LocationTruth) CountryLie() string {
	return ""
}

func (t LocationTruth) StateLie() string {
	return ""
}

func (t LocationTruth) CityLie() string {
	return ""
}

func (t LocationTruth) StreetAddressLie() string {
	return ""
}

func (t LocationTruth) PostalLie() string {
	return ""
}