package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type testCase struct {
	Name      string
	Turns     int // Expected number of turns (lines with moves)
	TimeLimit time.Duration
	Path      string // Relative path to the test file
}

func TestValidMaps(t *testing.T) {
	tests := []testCase{
		// Audit maps
		{Name: "example00.txt", Turns: 6, Path: "maps/audit/example00.txt"},
		{Name: "example01.txt", Turns: 8, Path: "maps/audit/example01.txt"},
		{Name: "example02.txt", Turns: 11, Path: "maps/audit/example02.txt"},
		{Name: "example03.txt", Turns: 6, Path: "maps/audit/example03.txt"},
		{Name: "example04.txt", Turns: 6, Path: "maps/audit/example04.txt"},
		{Name: "example05.txt", Turns: 8, Path: "maps/audit/example05.txt"},
		{Name: "example06.txt", TimeLimit: 90 * time.Second, Path: "maps/audit/example06.txt"},
		{Name: "example07.txt", TimeLimit: 150 * time.Second, Path: "maps/audit/example07.txt"},

		// Bhandari maps
		{Name: "bow.txt", Turns: 6, Path: "maps/bhandari/bow.txt"},
		{Name: "zhangir.txt", Turns: 6, Path: "maps/bhandari/zhangir.txt"},

		// Custom maps
		{Name: "extra-tails.txt", Turns: 4, Path: "maps/custom/extra-tails.txt"},
		{Name: "nrblzn.txt", Turns: 20, Path: "maps/custom/nrblzn.txt"},

		// Default maps
		{Name: "big_1.txt", Turns: 50, Path: "maps/default/big_1.txt"},
		{Name: "big_2.txt", Turns: 72, Path: "maps/default/big_2.txt"},
		{Name: "big_5.txt", Turns: 67, Path: "maps/default/big_5.txt"},
		{Name: "big_6.txt", Turns: 76, Path: "maps/default/big_6.txt"},
		{Name: "big_7.txt", Turns: 46, Path: "maps/default/big_7.txt"},
		{Name: "loop.txt", Turns: 3, Path: "maps/default/loop.txt"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Get absolute path to the test file
			absPath, err := filepath.Abs(test.Path)
			if err != nil {
				t.Fatalf("Failed to get absolute path: %v", err)
			}

			// Run the program
			cmd := exec.Command("go", "run", "./cmd/main.go", absPath)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			// Set timeout if specified
			if test.TimeLimit > 0 {
				timer := time.AfterFunc(test.TimeLimit, func() {
					cmd.Process.Kill()
				})
				defer timer.Stop()
			}

			err = cmd.Run()
			if err != nil {
				t.Fatalf("Program failed: %v\nOutput:\n%s", err, out.String())
			}

			// Count turns in output
			output := out.String()
			turnCount := 0
			for _, line := range strings.Split(output, "\n") {
				if strings.HasPrefix(line, "L") {
					turnCount++
				}
			}
			// Verify turn count if expected
			if test.Turns > 0 {
				if turnCount == 0 {
					t.Errorf("Expected %d turns, got 0. Full output:\n%s", test.Turns, output)
				} else if turnCount != test.Turns {
					t.Errorf("Expected %d turns, got %d", test.Turns, turnCount)
					t.Logf("Full output for %s:\n%s", test.Name, output)

				}
			}
		})
	}
}
