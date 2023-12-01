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

// Zoning in round 1

// Zoning in round 2

// Function to parse ASNs

// Function to parse IPs and IP Ranges

// Function to parse status codes

// Function to track progress % or active run time

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
