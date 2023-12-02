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
		bigNet
		webEnum
		ipEnum

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
// Start with a list of known URLs & Company names
    for i in company names:
        runCommand(currentFunction(), "whoxyrm")
        // add info to designated files
        runCommand(currentFunction(), "asnLookup")
        // Pull domains and IPs with ASN, add to URLs and a create a IP file
        pauseScript()
    for i in domain.txt:
        runCommand(currentFunction(), "findomain")
        runCommand(currentFunction(), "assetfinder")
        // add to a subs.txt file
    for i in subs.txt:
        runCommand(currentFunction(), "host")
        // Use host command to find sites are pointed to
        // Remove sites that are pointed to akima, amazon, etc.
        // Add to a hosts.txt
    for i in hosts.txt:
    // Use httpx using the asn flag and webserver flag. Add asn’s to file
        parseStatus()
        // Sort 400 type codes to a new file
    if loop count < 1:
        repeat loop
    if loop count = 2:
    // If scope is strict remove all URLs that do not strictly adhere
}

// Zoning in round 1
// implement as a go routine
func webEnum () {
    for i in 4xx.txt:
        runCommand(currentFunction(), "cewl")
        runCommand(currentFunction(), "fuff") // Use cewl to retrieve words from the website and use those words to fuzz
        runCommand(currentFunction(), "fuff") // Use a shorter to mid-size off the shelf wordlist
        if any directory is 200 or 500 print an alert and add it to targets.txt file
}

// Zoning in round 2
// implement as a go routine
func ipEnum () {
    for ip ranges:
        ping the range for live hosts
        for live hosts found add to a liveIP.txt file
    for IPs:
        perform masscan of the top 100 ports
        add open ports and their associated IP to a file or database
    for open ports found perform nmap scan:
        add info to a file
        parse file for web pages
        compare URLs to those that we have. Add unique URLs to a file and perform subdomain enumeration on those URLs.
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
