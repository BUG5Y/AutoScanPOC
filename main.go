package main

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)
////////////
//Commands//
////////////
// Whoxyrm
// findomain
// assetfinder
// nmap
// masscan
// httpx
// cewl
// fuff or similar

// Take CLI Args of files containing domains and business names

// Function to pause the script. Ask if the found URLs are good and okay to continue. (y/n) If we say no then it will wait till we adjust the file and say to continue(c)

// Main loop (Casting a big net)
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

// Function to parse ASNs
func parseASN (){
}
// Function to parse IPs and IP Ranges
func parseIP (){
}

func parseRange (){
}

// Function to parse status codes
func parseStatus (){
}

// Function to track progress % or active run time
func trackProg (){
}

func main() {
    runCommand(currentFunction(), "ping", "-c1", "google.commm")
}

// Function to pause script
func pauseScript () {
    // Have script pause and have us check if all URLs are in target. y/n. If bad urls remove the bad URLs from the file and then continue.
}

////////////////////
//Calling Commands//
////////////////////

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
