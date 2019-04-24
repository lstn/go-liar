package liar

//ipv4

type IPv4Truth struct {
	IP string
}

func (t IPv4Truth) Hash() string {
	return ""
}

func (t IPv4Truth) Lie() string {
	return ""
}

//ipv6

type IPv6Truth struct {
	IP string
}

func (t IPv6Truth) Hash() string {
	return ""
}

func (t IPv6Truth) Lie() string {
	return ""
}

//mac

type MACTruth struct {
	MAC string
}

func (t MACTruth) Hash() string {
	return ""
}

func (t MACTruth) Lie() string {
	return ""
}