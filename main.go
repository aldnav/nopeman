package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type RedactRule struct {
	Name   string `yaml:"name"`
	Match  string `yaml:"match"`
	Target string `yaml:"target"` // Not in use right now
}

// Reads redact.yml file and returns a list of rules
func readRedactRules() []RedactRule {
	yamlFile, err := os.ReadFile("redact.yml")
	if err != nil {
		log.Fatalf("Error reading redact.yml file: %v", err)
	}

	var rules []RedactRule
	err = yaml.Unmarshal(yamlFile, &rules)
	if err != nil {
		log.Fatalf("Error unmarshalling redact.yml file: %v", err)
	}

	return rules
}

func evalRules(rules []RedactRule, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	width := 100
	var allOutput strings.Builder

	for _, rule := range rules {
		// Print the rule name left-aligned
		fmt.Printf("Checking %s", rule.Name)

		cmd := exec.Command("jp", "-f", filePath, rule.Match)
		output, err := cmd.Output()
		if err != nil {
			log.Fatalf("Error running command: %v", err)
		}

		// Calculate padding and create dot string
		padding := width - len("Checking "+rule.Name)

		if len(output) > 0 && string(output) != "[]\n" {
			message := "Matches found"
			dots := strings.Repeat(".", padding-len(message))
			fmt.Printf("%s\033[38;5;214m%s\033[0m\n", dots, message)

			allOutput.WriteString(fmt.Sprintf("%s\n\n", rule.Name))
			allOutput.WriteString(string(output))
			allOutput.WriteString("\n\n")
		} else {
			message := "No matches found"
			dots := strings.Repeat(".", padding-len(message))
			fmt.Printf("%s\033[38;5;46m%s\033[0m\n", dots, message)
		}
	}
	fmt.Println("")
	fmt.Println(allOutput.String())
}

func main() {
	if _, err := exec.LookPath("jp"); err != nil {
		fmt.Println("Error: 'jp' is not installed. Please install it from 'https://github.com/jmespath/jp#installing'")
		os.Exit(1)
	}

	rules := readRedactRules()
	if len(os.Args) < 2 {
		asciiArt, err := os.ReadFile("ascii-art.txt")
		if err != nil {
			log.Fatalf("Error reading ascii-art.txt file: %v", err)
		}
		fmt.Println(string(asciiArt))
		fmt.Println("Review \"sus\" content in a Postman collection")
		fmt.Println("Usage: nopeman <inputfile>")
		os.Exit(1)
	}
	filePath := os.Args[1]
	evalRules(rules, filePath)
}
