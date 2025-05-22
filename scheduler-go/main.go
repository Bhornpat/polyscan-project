//Orchestrator: calls Rust & Perl
package main

import (
	"fmt"
	"os/exec"
	"net/http"
	"encoding/json"
)

type ScanResult struct {
	Source  string   `json:"source"`
	Matches []string `json:"matches"`
}

func callRustScanner(path string) string {
	out, err := exec.Command("../scanner-core/target/release/scanner-core", path).Output()
	if err != nil {
		fmt.Println("Rust Error:", err)
	}
	return string(out)
}

func callPerlScanner(path string) string {
	out, err := exec.Command("../regex-module/regex_module.pl", path).Output()
	if err != nil {
		fmt.Println("Perl Error:", err)
	}
	return string(out)
}

func handler(w http.ResponseWriter, r *http.Request) {
	file := "../shared/data/test.env"
	perl := callPerlScanner(file)
	rust := callRustScanner(file)

	result := []ScanResult{
		{Source: "perl", Matches: []string{perl}},
		{Source: "rust", Matches: []string{rust}},
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/scan", handler)
	fmt.Println("Scheduler running at http://localhost:8001")
	http.ListenAndServe(":8001", nil)
}

