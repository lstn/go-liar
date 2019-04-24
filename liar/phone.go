package liar

type PhoneTruth struct {
	PhoneNumber string
}

func (t PhoneTruth) Hash() string {
	return ""
}

func (t PhoneTruth) Lie() string {
	return ""
}