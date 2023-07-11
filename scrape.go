package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// Job struct to hold job data
type Job struct {
	Position string `json:"position"`
	Company  string `json:"company"`
}

func main() {
	c := colly.NewCollector()

	var jobs []Job

	c.OnHTML(".result-card", func(e *colly.HTMLElement) {
		position := e.ChildText(".result-card__title")
		company := e.ChildText(".result-card__subtitle")

		job := Job{
			Position: position,
			Company:  company,
		}

		jobs = append(jobs, job)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.linkedin.com/jobs/search/?currentJobId=3638868938&keywords=software%20engineer&refresh=true")

	file, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jobs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scraping complete. Results saved to output.json.")

	// Display the scraped jobs
	fmt.Println("Scraped Jobs:")
	for _, job := range jobs {
		fmt.Printf("Position: %s\n", job.Position)
		fmt.Printf("Company: %s\n\n", job.Company)
	}
}