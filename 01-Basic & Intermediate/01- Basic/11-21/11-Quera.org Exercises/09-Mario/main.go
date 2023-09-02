// ماریو و آنلاین شاپ


package main

import "fmt"

func main() {
	coinMario, gemMario, coinSpecial, gemSpecial, rate := GetUserInput()

	Gold(coinMario, gemMario, coinSpecial, gemSpecial, rate)

}

func GetUserInput() (
	coinMario,
	gemMario,
	coinSpecial,
	gemSpecial,
	rate float64) {

	_, err1 := fmt.Scanf("%f %f", &coinMario, &gemMario)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	_, err2 := fmt.Scanf("%f %f", &coinSpecial, &gemSpecial)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	_, err3 := fmt.Scanf("%f", &rate)
	if err3 != nil {
		fmt.Println("Error:", err3)
		return
	}
	return
}

func Gold(coinMario, gemMario, coinSpecial, gemSpecial, rate float64) {
	if coinMario >= coinSpecial && gemMario >= gemSpecial {
		fmt.Println("Yes")
	} else if coinMario < coinSpecial && gemMario > gemSpecial {
		for gemMario != gemSpecial {
			coinMario += 1 * rate
			if coinMario >= coinSpecial && gemMario >= gemSpecial {
				fmt.Println("Yes")
				break
			}
			gemMario--
		}
	} else if coinMario > coinSpecial && gemMario < gemSpecial {

		myGold := (coinMario - coinSpecial) / rate
		mySilver := int(myGold)
		gemMario := float64(mySilver)

		if coinMario >= coinSpecial && gemMario >= gemSpecial {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	} else {
		fmt.Println("No")
	}

}
