package main

type ServiceZ1 struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}


func (s ServiceZ1) getName() string {
	return s.description
}

func (s ServiceZ1) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
