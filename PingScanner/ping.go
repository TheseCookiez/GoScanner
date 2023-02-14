package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// This function creates a list of IP addresses to ping
// It takes a subnet as an argument and returns a list of IP addresses
func create_list(subnet string) []string {
	// Create an empty list of IP addresses
	ip_list := []string{}
	// Loop through the IP addresses and append them to the list
	for i := 1; i < 255; i++ {
		ip_addresses := subnet + "." + fmt.Sprint(i)
		ip_list = append(ip_list, ip_addresses)
	}
	// Return the list of IP addresses
	return ip_list
}

// This is the main function
func main() {
	// Get the arguments from the command line
	args := os.Args[1:]
	// Check if the arguments are less than 2
	// If they are less than 2, print the usage and exit
	if len(args) < 2 {
		fmt.Println("Usage: ./ping <subnet> <port>\n\n Example: ./ping 192.168.1 80")
		os.Exit(1)
	}
	// Start the timer
	start := time.Now()
	// Create a wait group
	var wg sync.WaitGroup
	// Create a list of IP addresses to ping from the create_list function
	list_to_ping := create_list(args[0])
	// Loop through the list of IP addresses and ping them
	for _, v := range list_to_ping {
		// Add 1 to the wait group
		wg.Add(1)
		// Create goroutines to ping the IP addresses
		go func(v string) {
			port := args[1]
			// Set the timeout to 1 second to wait for a response
			timeout := time.Duration(1 * time.Second)
			// Connect to the IP address over tcp and check if it is responding
			_, err := net.DialTimeout("tcp", v+":"+port, timeout)
			// If the IP address is not responding, print the error
			if err != nil {
				fmt.Printf("%s Host is not responding on %s\n", v, port)
				// If the IP address is responding, print the IP address and the port
			} else {
				fmt.Printf("%s %s %s\n", v, "responding on port:", port)
			}
			// Subtract 1 from the wait group
			wg.Done()
		}(v)
	}
	// Wait for all the goroutines to finish
	wg.Wait()
	// Stop the timer
	end := time.Now()
	// Print the time taken to ping all the IP addresses
	fmt.Printf("Time taken: %v", end.Sub(start))
}
