// Copyright (C) 2026 grusha4669
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License.

package main

import (
	"fmt"
	"net"
)

func IPv4ToISIS(ipStr string) string {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return "Error input."
	}

	s := fmt.Sprintf("%03d%03d%03d%03d", ip[0], ip[1], ip[2], ip[3])

	return fmt.Sprintf("%s.%s.%s", s[0:4], s[4:8], s[8:12])
}

func main() {
	var input string
	for {
		fmt.Print("IPv4: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error.", err)
			continue
		}

		result := IPv4ToISIS(input)

		if result == "Error input." {
			//result rewrited
			fmt.Println(result)
			continue
		}
		fmt.Println("ISIS:", result)
		fmt.Println("Success.")
	}
}
