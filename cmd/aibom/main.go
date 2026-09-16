package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Evidence struct {
	Pattern string `json:"pattern"`
	Line    int    `json:"line"`
	Snippet string `json:"snippet"`
}

type ScanResult struct {
	File          string     `json:"file"`
	AIProbability float64    `json:"ai_probability"`
	RiskLevel     string     `json:"risk_level"`
	Evidence      []Evidence `json:"evidence"`
	Error         string     `json:"error,omitempty"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: aibom scan <path_to_file>")
		os.Exit(1)
	}

	command := os.Args[1]
	filePath := os.Args[2]

	if command != "scan" {
		fmt.Println("❌ Unknown command. Use 'scan'.")
		os.Exit(1)
	}

	fmt.Printf("🔍 Scanning %s for AI-generated code...\n\n", filePath)

	// Run the Python engine
	cmd := exec.Command("python", "engine/detector.py", filePath)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ Error running scanner: %v\n", err)
		os.Exit(1)
	}

	// Parse JSON
	var result ScanResult
	if err := json.Unmarshal(output, &result); err != nil {
		fmt.Printf("❌ Error parsing scanner output: %v\n", err)
		os.Exit(1)
	}

	if result.Error != "" {
		fmt.Printf("❌ Error: %s\n", result.Error)
		os.Exit(1)
	}

	// Print the Report
	fmt.Printf("📄 File: %s\n", result.File)
	fmt.Printf("🤖 AI Probability: %.0f%%\n", result.AIProbability*100)
	fmt.Printf("⚠️  Risk Level: %s\n", result.RiskLevel)
	
	if len(result.Evidence) > 0 {
		fmt.Println("\n📋 Evidence Found:")
		for _, e := range result.Evidence {
			fmt.Printf("   - Line %d [%s]: %s...\n", e.Line, e.Pattern, e.Snippet)
		}
	} else {
		fmt.Println("\n✅ No obvious AI-generated patterns detected.")
	}
}