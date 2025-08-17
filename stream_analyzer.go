package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Stats struct {
	total       int
	errors      int
	statusCount map[int]int
	ipCount     map[string]int
}

func newStats() *Stats {
	return &Stats{
		statusCount: map[int]int{},
		ipCount:     map[string]int{},
	}
}

func main() {
	// Works with Nginx/Apache "combined" format
	re := regexp.MustCompile(`^(\S+) \S+ \S+ \[[^\]]+\] "([^"]*)" (\d{3}) \S+`)

	stats := newStats()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Print a summary every 5 seconds
	go func() {
		for range ticker.C {
			printStats(stats)
		}
	}()

	// Read lines from stdin (we'll pipe the log in)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		m := re.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		ip := m[1]
		code, _ := strconv.Atoi(m[3])
		stats.total++
		stats.statusCount[code]++
		if code >= 400 {
			stats.errors++
		}
		stats.ipCount[ip]++
	}
	printStats(stats) // final print when you stop it
}

func printStats(s *Stats) {
	fmt.Println("=== Live Log Stats ===")
	fmt.Printf("Total: %d  Errors(4xx/5xx): %d\n", s.total, s.errors)

	// Status codes
	type kv struct{ k, v int }
	var codes []kv
	for k, v := range s.statusCount {
		codes = append(codes, kv{k, v})
	}
	sort.Slice(codes, func(i, j int) bool { return codes[i].k < codes[j].k })
	fmt.Print("Status: ")
	for i, c := range codes {
		if i > 0 { fmt.Print(" | ") }
		fmt.Printf("%d=%d", c.k, c.v)
	}
	fmt.Println()

	// Top IPs
	type ipkv struct{ ip string; n int }
	var ips []ipkv
	for ip, n := range s.ipCount { ips = append(ips, ipkv{ip, n}) }
	sort.Slice(ips, func(i, j int) bool {
		if ips[i].n == ips[j].n { return ips[i].ip < ips[j].ip }
		return ips[i].n > ips[j].n
	})
	fmt.Println("Top IPs:")
	for i := 0; i < len(ips) && i < 5; i++ {
		fmt.Printf("  %-22s %d\n", ips[i].ip, ips[i].n)
	}
	fmt.Println()
}
