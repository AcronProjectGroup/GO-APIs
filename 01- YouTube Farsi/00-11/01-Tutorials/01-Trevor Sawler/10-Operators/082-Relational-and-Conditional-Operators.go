package main

func main() {
	second := 20
	minute := 3

	// for readable code after six months, it's better to use pranteses
	if (minute < 59) && ((second + 1) > 59) {
		minute++
	}

}
