package main

import (
	"fmt"
	"log"
	"main/github"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d, %9.9s %9.9s %.55s\n",
			item.Number, item.User.Login, getAgeCategory(item.CreatedAt), item.Title)
	}
}

func getAgeCategory(date time.Time) string {
	today := time.Now()
	oneMonthBefore := today.AddDate(0, -1, 0)
	if date.After(oneMonthBefore) {
		return "<1 month"
	}
	oneYearBefore := today.AddDate(-1, 0, 0)
	if date.After(oneYearBefore) {
		return "<1 year"
	}
	return ">1 year"
}