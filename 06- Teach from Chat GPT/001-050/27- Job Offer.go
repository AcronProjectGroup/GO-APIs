/*
    Important Source : https://sinalalenakhsh.notion.site/1-Chat-GPT-310ba8a7e074493e95d8c45bfe6aed86

   Notion Source: https://sinalalenakhsh.notion.site/12-First-Project-d828d312c2e44551bffa7c4c1589979c
   از کاربر ورودی میگیره
   تحلیل روی اطلاعات کاربر انجام میده
   اگر اطلاعات مطابق نظرش بود
   خروجی رو به کاربر نمایش میده
   اگر اطلاعات مطابق نظرش نبود
   یک خروجی ثابت رو به کاربر نمایش میده

       داستان کاریابی در آلمان

   variables:
       userName            string
       jobTitle            string
       userAge             int
       yearsExperience     int

   we will if condition

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userName, jobTitle string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the program to help immigrants to Germany!")
	fmt.Println("Please enter your personal and job-related information.")

	fmt.Print("What's You full name? ")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSuffix(userName, "\n")

	fmt.Printf("Hello %v \n", userName)

	fmt.Print("What do you looking for job? ")
	jobTitle, _ = reader.ReadString('\n')
	jobTitle = strings.TrimSuffix(jobTitle, "\n")

	if jobTitle == "programming lang" {
		fmt.Printf("congratulate %v , one job is ON for this: %v \n", userName, jobTitle)
	}

}
