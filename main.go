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
// Whoxyrm with the company name(s) add into the known URLs
// Have script pause and have us check if all URLs are in target. y/n. If bad urls remove the bad URLs from the file and then continue.
// ASN lookup with company names. Create a file of ASNs
// Pull domains and IPs with ASN, add to URLs and a create a IP file
// Have script pause and check if all URLs are in target. y/n. If not remove the bad URLs from the file and then continue.
// Use findomain
// Use assetfinder
// Use host command to find sites are pointed to
// Remove sites that are pointed to akima, amazon, etc.
// Use httpx using the asn flag and webserver flag. Add asn’s to file
// Sort 400 type codes to a new file
// If loop has not been completed go back to step 4
// If scope is strict remove all URLs that do not strictly adhere

}

// Zoning in round 1
func webEnum () {
}

// Zoning in round 2
func ipEnum () {
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
