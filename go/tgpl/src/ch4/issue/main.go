// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	// from new to old
	sort.Sort(sort.Reverse(result))
	items := result.Items

	m := sort.Search(len(items),
		func(i int) bool {
			return time.Since(items[i].CreatedAt) >= 30*24*time.Hour
		})
	if m < len(items) {
		fmt.Println("Less than a month:")
		for i := 0; i < m; i++ {
			fmt.Printf("#%-5d %5ddays %9.9s %.55s\n",
				items[i].Number, daysAgo(items[i].CreatedAt), items[i].User.Login, items[i].Title)
		}
		items = items[m:]
	}

	y := sort.Search(len(items),
		func(i int) bool {
			return time.Since(items[i].CreatedAt) >= 365*24*time.Hour
		})
	if y < len(items) {
		fmt.Println("Less than a year:")
		for i := 0; i < y; i++ {
			fmt.Printf("#%-5d %5ddays %9.9s %.55s\n",
				items[i].Number, daysAgo(items[i].CreatedAt), items[i].User.Login, items[i].Title)
		}
		items = items[y:]
	}

	if len(items) > 0 {
		fmt.Println("More than a year:")
		for _, item := range items {
			fmt.Printf("#%-5d %5ddays %9.9s %.55s\n",
				item.Number, daysAgo(item.CreatedAt), item.User.Login, item.Title)
		}
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
