package main

type Service1 struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s Service1) getName1() string {
	return s.description
}

func (s Service1) getCost1(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

