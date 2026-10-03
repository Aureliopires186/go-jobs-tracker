


Google chopping price tracker:


package main

import (
  "fmt"
  "log"
  "os"

  g "github.com/serpapi/google-search-results-golang"
)

func main() {
  apiKey := os.Getenv("SERPAPI_KEY")
  if apiKey == "" {
    log.Fatal("SERPAPI_KEY is not set")
  }

  params := map[string]string{
    "engine": "google_shopping",
    "q": "golang book",
  }

  search := g.NewGoogleSearch(params, apiKey)
  results, err := search.GetJSON()
  if err!= nil {
    log.Fatalf("failed to get shopping results: %v", err)
  }

  shoppingResults, ok := results["shopping_results"].([]interface{})
  if!ok {
    log.Printf("No shopping_results found: %+v", results)
    return
  }

  fmt.Printf("Found %d products\n", len(shoppingResults))
  for _, item := range shoppingResults {
    if m, ok := item.(map[string]interface{}); ok {
      fmt.Printf("- %v: %v\n", m["title"], m["price"])
    }
  }
}

