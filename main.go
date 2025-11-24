package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"runtime"

	"github.com/BurntSushi/toml"
)

const (
	ColorGreen = "\033[32m"
	ColorRed   = "\033[31m"
	ColorReset = "\033[0m"
)

type DevCheckConfig struct {
	Versions    map[string]string `toml:"versions"`
	Env         EnvConfig         `toml:"env"`
	Connections map[string]string `toml:"connections"`
}

type EnvConfig struct {
	Keys []string `toml:"keys"`
}

func main() {
	var conf DevCheckConfig
	if _, err := toml.DecodeFile("devcheck.toml", &conf); err != nil {
		log.Fatalf("Error reading devcheck.toml: %v", err)
	}

	fmt.Println("--- 🔍 Starting DevCheck ---")
	
	allPassed := true

	fmt.Println("\n[Versions]")
	for tool, constraint := range conf.Versions {
		if !checkVersion(tool, constraint) {
			allPassed = false
		}
	}

	fmt.Println("\n[Environment]")
	for _, key := range conf.Env.Keys {
		if !checkEnv(key) {
			allPassed = false
		}
	}

	fmt.Println("\n[Connections/Commands]")
	for service, cmd := range conf.Connections {
		if !checkCommand(service, cmd) {
			allPassed = false
		}
	}

	fmt.Println("\n----------------------------")
	if allPassed {
		fmt.Printf("%s✅ Setup Complete! You are ready to code.%s\n", ColorGreen, ColorReset)
	} else {
		fmt.Printf("%s❌ Setup Failed. Fix the issues above.%s\n", ColorRed, ColorReset)
		os.Exit(1) 
	}
}
func checkVersion(tool, constraint string) bool {
	cmd := exec.Command(tool, "--version")
	output, err := cmd.Output()
	
	if err != nil {
		printResult(tool, fmt.Sprintf("Not found or error running '%s --version'", tool), false)
		return false
	}

	outStr := string(output)
	if strings.Contains(outStr, constraint) {
		printResult(tool, fmt.Sprintf("Found %s (Matched '%s')", strings.TrimSpace(outStr), constraint), true)
		return true
	}

	printResult(tool, fmt.Sprintf("Found %s (Expected '%s')", strings.TrimSpace(outStr), constraint), false)
	return false
}

func checkEnv(key string) bool {
	val := os.Getenv(key)
	if val != "" {
		printResult(key, "Present", true)
		return true
	}
	printResult(key, "Missing", false)
	return false
}

func checkCommand(name, command string) bool {
	command=os.ExpandEnv(command)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {

		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	
	if err := cmd.Run(); err != nil {
		printResult(name, fmt.Sprintf("Failed to run: %s", command), false)
		return false
	}
	printResult(name, "Connection Successful", true)
	return true
}

func printResult(name, message string, pass bool) {
	if pass {
		fmt.Printf("  %-15s %s[PASSED]%s %s\n", name, ColorGreen, ColorReset, message)
	} else {
		fmt.Printf("  %-15s %s[FAILED]%s %s\n", name, ColorRed, ColorReset, message)
	}
}