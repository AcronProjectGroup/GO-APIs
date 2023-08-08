package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (handler *server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Acron Projections !!!\n")
}


func (handler *server) bmi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		defer r.Body.Close()

		rawBody, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		
		body := &struct {
			Height float64 `json:"height"`
			Weight float64 `json:"weight"`
		}{}

		err = json.Unmarshal(rawBody, body)
		if err != nil {
			panic(err)
		}
		
		if body.Height < 40 || body.Height > 250 {
			fmt.Fprintf(w, "Invalid height provided: %.2f", body.Height)
			return
		}
		body.Height *= 0.01

		if body.Weight < 5 || body.Weight > 200 {
			fmt.Fprintf(w, "Invalid weight provided: %.2f", body.Weight)
			return
		}

		bmi := body.Weight / (body.Height * body.Height)  
		fmt.Fprintf(w, "Your BMI is = %.2f\n", bmi)

	default:
		fmt.Fprintf(w, "Sorry, only the post method is supported!!\n")
	}

}
