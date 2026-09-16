package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		printError("Usage: aibom scan <path_to_file>")
		os.Exit(1)
	}

	command := os.Args[1]
	filePath := os.Args[2]

	if command != "scan" {
		printError("Unknown command. Use 'scan'.")
		os.Exit(1)
	}

	// Find the Python script relative to the executable's location
	exePath, err := os.Executable()
	if err != nil {
		printError(fmt.Sprintf("Error getting executable path: %v", err))
		os.Exit(1)
	}
	exeDir := filepath.Dir(exePath)
	pythonScript := filepath.Join(exeDir, "engine", "detector.py")

	// Run the Python engine
	cmd := exec.Command("python", pythonScript, filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		printError(fmt.Sprintf("Error running scanner: %v | Output: %s", err, string(output)))
		os.Exit(1)
	}

	// Parse JSON
	var result ScanResult
	if err := json.Unmarshal(output, &result); err != nil {
		printError(fmt.Sprintf("Error parsing scanner output: %v | Raw: %s", err, string(output)))
		os.Exit(1)
	}

	if result.Error != "" {
		printError(result.Error)
		os.Exit(1)
	}

	// Print the Report
	fmt.Printf("🔍 Scanning %s for AI-generated code...\n\n", filePath)
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

func printError(msg string) {
	res := ScanResult{Error: msg}
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}