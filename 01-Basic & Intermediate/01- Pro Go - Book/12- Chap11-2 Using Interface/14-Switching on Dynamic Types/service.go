package main

type Service4 struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

func (s Service4) getName2() string {
	return s.description
}

func (s Service4) getCost2(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
