package main

type Service3 struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

func (s Service3) getName2() string {
	return s.description
}

func (s Service3) getCost2(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
