// Notion Source: https://sinalalenakhsh.notion.site/12-First-Project-d828d312c2e44551bffa7c4c1589979c

package main

import (
    "fmt"
)

func main() {
    var name, jobTitle string
    var age, yearsExperience int

    fmt.Println("Welcome to the program to help immigrants to Germany!")
    fmt.Println("Please enter your personal and job-related information.")

    fmt.Print("Name and last name: ")
    fmt.Scanln(&name)

    fmt.Print("Age: ")
    fmt.Scanln(&age)

    fmt.Print("Job title: ")
    fmt.Scanln(&jobTitle)

    fmt.Print("Years of work experience: ")
    fmt.Scanln(&yearsExperience)

    // Checking the person's eligibility for the desired job
    if jobTitle == "electrical engineer" && yearsExperience >= 2 {
        fmt.Println("You can work in Germany as an electrical engineer.")
    } else if jobTitle == "programmer" && yearsExperience >= 3 {
        fmt.Println("You can work in Germany as a programmer.")
    } else if jobTitle == "software engineer" && yearsExperience >= 5 {
        fmt.Println("You can work in Germany as a software engineer.")
    } else {
        fmt.Println("Unfortunately, you do not have the necessary qualifications for the desired job.")
    }
}