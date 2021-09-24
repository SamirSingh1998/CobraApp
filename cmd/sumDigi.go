/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// sumDigiCmd represents the sumDigi command
var sumDigiCmd = &cobra.Command{
	Use:   "sumDigi",
	Short: "Helps us to add odd and even numbers",
	Long:  `Helps us to add odd and even numbers with the help of -e and -o flags`,
	Run: func(cmd *cobra.Command, args []string) {
		f1Status, err1 := cmd.Flags().GetBool("even")
		f2Status, err2 := cmd.Flags().GetBool("odd")
		if err1 != nil {
			fmt.Println("Error found", err1)
		} else if err2 != nil {
			fmt.Println("Error found", err2)
		} else {
			if f1Status {
				addEven(args)
			} else if f2Status {
				fmt.Println("Hello from f2 status")
				addOdd(args)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sumDigiCmd)
	sumDigiCmd.Flags().BoolP("even", "e", false, "Sum of even digits")
	sumDigiCmd.Flags().BoolP("odd", "o", false, "Sum of odd digits")
}
func addEven(args []string) {
	var sum float64
	for _, v := range args {
		switch string(v) {
		case "+":
			continue
		default:
			if a, _ := strconv.ParseFloat(string(v), 64); math.Mod(a, 2.0) == 0 {
				sum += a
			}
		}
	}
	fmt.Println("Sum of even digits is:", sum)
}
func addOdd(args []string) {
	var sum float64
	for _, v := range args {
		switch string(v) {
		case "+":
			continue
		default:
			if a, _ := strconv.ParseFloat(string(v), 64); math.Mod(a, 2.0) != 0 {
				sum += a
			}
		}
	}
	fmt.Println("Sum of odd digits is:", sum)
}
