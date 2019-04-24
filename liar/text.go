package liar

// names
type PersonTruth struct {
	name string
}

func (t PersonTruth) Hash() string {
	return ""
}

func (t PersonTruth) Lie() string {
	// FORMAT: Firstname Lastname <email@email.domain>
	return ""
}

func (t PersonTruth) FirstNameLie() string {
	return ""
}

func (t PersonTruth) LastNameLie() string {
	return ""
}

func (t PersonTruth) FullNameLie() string {
	return ""
}

func (t PersonTruth) EmailLie() string {
	return ""
}

// business names
type CompanyTruth struct {
	name string
}

func (t CompanyTruth) Hash() string {
	return ""
}

func (t CompanyTruth) Lie() string {
	return ""
}
