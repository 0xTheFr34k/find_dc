package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/tabwriter"
)

// Function to find the domain controller using nslookup
func findDomainController(domain string, dnsServer string) (string, error) {
	cmd := exec.Command("nslookup", "-type=srv", fmt.Sprintf("_ldap._tcp.dc._msdcs.%s", domain), dnsServer)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", nil
	}

	// Parse the output to find the domain controller
	re := regexp.MustCompile(`\b(\w+\.` + regexp.QuoteMeta(domain) + `)\b`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) > 0 {
		return matches[1], nil
	}

	return "", nil
}

func main() {
	// Define the help flag
	helpFlag := flag.Bool("help", false, "Display help information")
	hFlag := flag.Bool("h", false, "Display help information")
	flag.Parse()

	// Display help information if --help or -h is provided
	if *helpFlag || *hFlag {
		printHelp()
		return
	}

	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s+\d+\s+(\w+).*\(domain:([^)]+)`)

	var entries []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			ip := strings.ToLower(matches[1])
			hostname := strings.ToLower(matches[2])
			domain := strings.ToLower(matches[3])

			dc, err := findDomainController(domain, ip)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error finding domain controller for %s: %v\n", domain, err)
				entries = append(entries, fmt.Sprintf("%s\t%s.%s\t%s\t", ip, hostname, domain, hostname))
			} else {
				dc = strings.Replace(dc, "_msdcs.", "", 1)
				entries = append(entries, fmt.Sprintf("%s\t%s.%s\t%s\t%s", ip, hostname, domain, hostname, dc))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	if len(entries) > 0 {
		fmt.Println("Add the following lines to /etc/hosts:")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		for _, entry := range entries {
			fmt.Fprintln(w, entry)
		}
		w.Flush()
	} else {
		fmt.Println("No valid IP, hostname, and domain pairs found in the input.")
	}
}

// Function to display help information
func printHelp() {
	fmt.Println(`Usage: find_dc [OPTIONS]

Description:
  A Go tool to parse 'nxc' output, find domain controllers, and generate /etc/hosts entries.

Options:
  -h, --help    Display this help message and exit

Examples:
  nxc smb 192.168.56.10-12 | find_dc
  nxc smb 192.168.56.10-12 | find_dc --help

Output:
  The tool prints /etc/hosts entries in a table format.`)
}
