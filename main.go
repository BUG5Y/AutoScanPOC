package main

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "AutoScan"
	app.Usage = "A simple CLI application"

	// Define flags
	businessFlag := cli.StringFlag{
		Name:  "b",
		Usage: "Business argument (mandatory)",
	}

	domainFlag := cli.StringFlag{
		Name:  "d",
		Usage: "Domain argument (mandatory)",
	}

	strictFlag := cli.BoolFlag{
		Name:  "s",
		Usage: "Strict argument (optional)",
	}

	helpFlag := cli.BoolFlag{
		Name:  "h",
		Usage: "Help argument (optional)",
	}

	// Set up the CLI command
	app.Flags = []cli.Flag{&businessFlag, &domainFlag, &strictFlag, &helpFlag}

	app.Action = func(c *cli.Context) error {
		// Check for mandatory arguments
		if !c.IsSet("b") || !c.IsSet("d") {
			fmt.Println("Error: Business (-b) and Domain (-d) are mandatory arguments")
			return nil
		}

		// Retrieve values
		business := c.String("b")
		domain := c.String("d")
		strict := c.Bool("s")

		// Your application logic here
		bigNet()

		wg.Add(2)
		// Launch goroutines for webEnum and ipEnum
		go webEnum()
		go ipEnum()

		// Wait for all goroutines to finish
		wg.Wait()

		// Print values for demonstration
		fmt.Printf("Business: %s\n", business)
		fmt.Printf("Domain: %s\n", domain)
		fmt.Printf("Strict: %t\n", strict)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Main logic
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func bigNet() {
	for loopCount = 1; loopCount <= 2; loopCount++ {
		// Start with a list of known URLs & Company names
		companyNames := []string{"Company1", "Company2", "Company3"}
		domainsFileName := "domains.txt"
		subsFileName := "subs.txt"
		hostsFileName := "hosts.txt"

		// Loop through company names
		for _, company := range companyNames {
			runCommand(currentFunction(), "whoxyrm", company)
			// add info to designated files
			runCommand(currentFunction(), "asnLookup", company)
			// Pull domains and IPs with ASN, add to URLs and create an IP file
			pauseScript()
		}

		// Loop through domains
		domains, err := readLines(domainsFileName)
		if err != nil {
			fmt.Println("Error reading domains file:", err)
			return
		}

		for _, domain := range domains {
			runCommand(currentFunction(), "findomain", domain)
			runCommand(currentFunction(), "assetfinder", domain)
			// add to a subs.txt file
			pauseScript()
		}

		// Loop through subs
		subs, err := readLines(subsFileName)
		if err != nil {
			fmt.Println("Error reading subs file:", err)
			return
		}

		for _, sub := range subs {
			runCommand(currentFunction(), "host", sub)
			// Use host command to find sites are pointed to
			// Remove sites that are pointed to akima, amazon, etc.
			// Add to a hosts.txt
			pauseScript()
		}

		// Loop through hosts
		hosts, err := readLines(hostsFileName)
		if err != nil {
			fmt.Println("Error reading hosts file:", err)
			return
		}

		// Use httpx using the asn flag and webserver flag. Add asns to file
		runCommand(currentFunction(), "httpx", "-asn", "-webserver", "-i", hostsFileName)
		parseStatus()
		// Sort 400 type codes to a new file

		// Additional logic for loop count and scope
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

var wg sync.WaitGroup

func webEnum() {
	defer wg.Done()

	// Assuming 4xx.txt is a file with directories to be enumerated
	directories, err := readLines("4xx.txt")
	if err != nil {
		fmt.Println("Error reading 4xx.txt file:", err)
		return
	}

	for _, directory := range directories {
		go func(dir string) {
			runCommand(currentFunction(), "cewl", dir)
			runCommand(currentFunction(), "fuff", dir)
			runCommand(currentFunction(), "fuff", dir)

			// Additional logic to check directories for 200 or 500 status and update targets.txt
			// if any directory is 200 or 500 print an alert and add it to targets.txt file

		}(directory)
	}
}

func ipEnum() {
	defer wg.Done()

	// Assuming ipRanges is a slice of IP ranges to be processed
	ipRanges := []string{"192.168.1.0/24", "10.0.0.0/24"}

	for _, ipRange := range ipRanges {
		go func(rangeIP string) {
			// Ping the range for live hosts
			runCommand("ping", "-c", "1", rangeIP)

			// Additional logic to add live hosts to liveIP.txt
			// ...

			// Perform masscan of the top 100 ports
			runCommand("masscan", rangeIP, "--top-ports", "100")

			// Additional logic to parse masscan results and add to a file or database
			// ...

			// Perform nmap scan on open ports
			runCommand("nmap", "-p", "open_ports", rangeIP)

			// Additional logic to add nmap scan info to a file
			// ...

			// Parse file for web pages and perform subdomain enumeration
			// Additional logic for parsing and subdomain enumeration
			// ...

		}(ipRange)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helpers and Parsers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Parse ASNs
func parseASN (){
}
// Parse IPs and IP Ranges
func parseIP (){
}

func parseRange (){
}

// Parse status codes
func parseStatus (){
}

// track progress % or active run time
func trackProg (){
}

// Pause scripts
func pauseScript() bool {
	fmt.Println("Script paused. Do you want to continue? (y/n)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()

	switch response {
	case "y", "Y":
		return true
	case "n", "N":
		fmt.Println("Make necessary adjustments to the file and type 'c' to continue.")
		scanner.Scan()
		continueResponse := scanner.Text()
		return continueResponse == "c" || continueResponse == "C"
	default:
		fmt.Println("Invalid response. Please type 'y' to continue or 'n' to make adjustments.")
		return pauseScript()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Commands 
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Whoxyrm
// findomain
// assetfinder
// nmap
// masscan
// httpx
// cewl
// fuff or similar

func commandErrorMessage(stderr bytes.Buffer, program string) string {
    message := string(stderr.Bytes())

    if len(message) == 0 {
        message = "the command doesn't exist: " + program + "\n"
    }

    return message
}

func currentFunction() string {
    counter, _, _, success := runtime.Caller(1)

    if !success {
        println("functionName: runtime.Caller: failed")
        os.Exit(1)
    }

    return runtime.FuncForPC(counter).Name()
}

func printCommandError(stderr bytes.Buffer, callerFunc string, program string, args ...string) {
    printCommandErrorUbication(callerFunc, program, args...)
    fmt.Fprintf(os.Stderr, "%s", commandErrorMessage(stderr, program))
}

func printCommandErrorUbication(callerFunc string, program string, args ...string) {
    format := "error at: %s: %s %s\n"
    argsJoined := strings.Join(args, " ")
    fmt.Fprintf(os.Stderr, format, callerFunc, program, argsJoined)
}

func runCommand(callerFunc string, program string, args ...string) {
    command := exec.Command(program, args...)
    var stderr bytes.Buffer
    command.Stderr = &stderr
    fail := command.Run()

    if fail != nil {
        printCommandError(stderr, callerFunc, program, args...)
        os.Exit(1)
    }
}
