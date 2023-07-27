package main

type Service2 struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service2) getName2() string {
	return s.description
}

func (s Service2) getCost2(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

