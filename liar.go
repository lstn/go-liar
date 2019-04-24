package liar
import (
)

// type Liar interface {
// 	// text
// 	FirstName(fname string) string
// 	LastName(fname string) string
// 	FullName(fname string) string
// 	Email(fname string) string
// 	Company(bname string) string

// 	// net
// 	IPv4(ip string) string
// 	IPv6(ip string) string
// 	MAC(ip string) string

// 	// phone
// 	Phone()
// 	Phone(phone string) string
// 	Fax(fax string) string

// 	// location
// 	// Location(flocation string) string
// 	// Country(flocation string) string
// 	// State(flocation string) string
// 	// City(flocation string) string
// 	// Postal(flocation string) string
// }

type Liar struct {
	Secret string
}

type PersonParams struct {
	Name string
	Email string
}

type PhoneParams struct {
	Phone string
	Fax string
}


// person
func (l Liar) Person(p PersonParams) string {
	return ""
}

// company
func (l Liar) Company(name string) string {
	return ""
}

// phone
func (l Liar) Phone(p PhoneParams) string {
	return ""
}

// IPv4
func (l Liar) IPv4(ip string) string {
	return ""
}

// IPv6
func (l Liar) IPv6(ip string) string {
	return ""
}

// MAC
func (l Liar) MAC(mac string) string {
	return ""
}