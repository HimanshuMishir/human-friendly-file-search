package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jdkato/prose/v2"
)

// ParsedQuery represents the structured result of a query
type ParsedQuery struct {
	FileType   string
	MaxSizeKB  int
	Modified   string
	Year       int
}

// extractFileSize parses text for size constraints (e.g., "under 100kb")
func extractFileSize(text string) (int, error) {
	sizePattern := regexp.MustCompile(`(\d+)\s?(b|bytes?|kb|kilobytes?|mb|megabytes?|gb|gigabytes?)`)
	match := sizePattern.FindStringSubmatch(text)
	if len(match) >= 3 {
		size, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		unit := strings.ToLower(match[2])
		switch unit {
		case "kb", "kilobyte", "kilobytes":
			return size, nil
		case "mb", "megabyte", "megabytes":
			return size * 1024, nil
		case "gb", "gigabyte", "gigabytes":
			return size * 1024 * 1024, nil
		case "b", "byte", "bytes":
			return size / 1024, nil // Convert bytes to KB
		default:
			return 0, fmt.Errorf("unknown size unit")
		}
	}
	return 0, fmt.Errorf("no size found")
}

// extractYearOrModified parses text for year or modification-related phrases
func extractYearOrModified(text string) (string, int) {
	// Check for year
	yearPattern := regexp.MustCompile(`\b(19|20)\d{2}\b`)
	yearMatch := yearPattern.FindStringSubmatch(text)
	if len(yearMatch) > 0 {
		year, _ := strconv.Atoi(yearMatch[0])
		return "", year
	}

	// Check for modification-related phrases
	if regexp.MustCompile(`last week`).MatchString(text) {
		return "last_week", 0
	}
	if regexp.MustCompile(`yesterday`).MatchString(text) {
		return "yesterday", 0
	}
	if regexp.MustCompile(`last month`).MatchString(text) {
		return "last_month", 0
	}

	return "", 0
}

// parseQuery processes the user query and extracts parameters
func parseQuery(query string) (*ParsedQuery, error) {
	doc, err := prose.NewDocument(query)
	if err != nil {
		return nil, err
	}

	parsed := &ParsedQuery{}

	// Extract entities (e.g., file types like "workflow file")
	for _, entity := range doc.Entities() {
		if entity.Label == "NN" { // Adjust label based on Prose's POS tagging
			parsed.FileType = entity.Text
		}
	}

	// Extract size
	size, err := extractFileSize(query)
	if err == nil {
		parsed.MaxSizeKB = size
	}

	// Extract year or modified date
	modified, year := extractYearOrModified(query)
	parsed.Modified = modified
	parsed.Year = year

	return parsed, nil
}

func main() {
	query := "Find workflow file under 110gb modified in year 2024"

	parsedQuery, err := parseQuery(query)
	if err != nil {
		log.Fatalf("Error parsing query: %v", err)
	}

	fmt.Printf("Parsed Query:\nFile Type: %s\nMax Size: %d KB\nModified: %s\nYear: %d\n",
		parsedQuery.FileType, parsedQuery.MaxSizeKB, parsedQuery.Modified, parsedQuery.Year)
}
