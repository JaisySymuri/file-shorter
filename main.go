package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type Category struct {
	Name        string   `mapstructure:"name"`
	Rules       []string `mapstructure:"rules"`
	Destination string   `mapstructure:"destination"`
}

var categories []Category

func main() {
	for {
		fmt.Println("What do you want to tidy up?")
		fmt.Println("1. Downloads")
		fmt.Println("2. Academic")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tidyUp("downloads", "C:\\Users\\bests\\Downloads\\")
		case 2:
			tidyUp("academic", "C:\\Users\\bests\\Downloads\\[Academic\\")
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func tidyUp(configFile, Dir string) {
	// Load configuration from config.yml
	if err := loadConfig(configFile); err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Check for duplicate rules
	if err := checkConfig(); err != nil {
		fmt.Println("Config validation failed:", err)
		fmt.Println("Please fix the YAML file before proceeding.")
		return
	}

	files, err := os.ReadDir(Dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		filePath := filepath.Join(Dir, file.Name())
		moveFile(filePath, file.Name())
	}
}

func loadConfig(configFile string) error {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.UnmarshalKey("categories", &categories)
}

func checkConfig() error {
	ruleMap := make(map[string]string) // Maps rule to category name

	for _, category := range categories {
		localRules := make(map[string]bool) // To check duplicates within the same category

		for _, rule := range category.Rules {
			// Check if the rule is already in the same category
			if localRules[rule] {
				return fmt.Errorf("duplicate rule '%s' found in category '%s'", rule, category.Name)
			}
			localRules[rule] = true

			// Check if the rule is already assigned to another category
			if existingCategory, exists := ruleMap[rule]; exists {
				return fmt.Errorf("rule '%s' is used in both '%s' and '%s'", rule, existingCategory, category.Name)
			}
			ruleMap[rule] = category.Name
		}
	}

	return nil
}

func moveFile(filePath, fileName string) {
	lowerName := strings.ToLower(fileName)

	for _, category := range categories {
		for _, rule := range category.Rules {
			pattern := fmt.Sprintf(`.*%s.*`, regexp.QuoteMeta(rule))
			matched, _ := regexp.MatchString(strings.ToLower(pattern), lowerName)
			if matched {
				destinationPath := filepath.Join(category.Destination, fileName)
				if err := os.MkdirAll(category.Destination, os.ModePerm); err != nil {
					fmt.Println("Error creating destination folder:", err)
					return
				}

				if err := os.Rename(filePath, destinationPath); err != nil {
					fmt.Println("Error moving file:", err)
				} else {
					fmt.Printf("Moved %s to %s\n", fileName, category.Destination)
				}
				return // Move only once per file
			}
		}
	}
}
