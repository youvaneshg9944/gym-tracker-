package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gymtracker/gym"
)

func readLine(in *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !in.Scan() {
		return "", false
	}
	return strings.TrimSpace(in.Text()), true
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- GYM MEMBERSHIP TRACKER ---")
		fmt.Println("1. Enroll Member")
		fmt.Println("2. Lookup Member")
		fmt.Println("3. Renew Subscription")
		fmt.Println("4. Revoke Membership")
		fmt.Println("5. Exit")

		choiceText, ok := readLine(in, "Enter your choice: ")
		if !ok {
			return
		}
		choice, err := strconv.Atoi(choiceText)
		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			id, ok := readLine(in, "Enter Member ID: ")
			if !ok {
				return
			}
			name, ok := readLine(in, "Enter Full Name: ")
			if !ok {
				return
			}
			plan, ok := readLine(in, "Enter Plan (Standard/Premium/VIP): ")
			if !ok {
				return
			}
			monthsText, ok := readLine(in, "Enter Months: ")
			if !ok {
				return
			}
			months, err := strconv.Atoi(monthsText)
			if err != nil {
				fmt.Println("Months must be a number.")
				continue
			}

			member := gym.Member{
				MemberID:        id,
				FullName:        name,
				PlanType:        plan,
				MonthsRemaining: months,
				IsActive:        true,
			}
			if err := gym.EnrollMember(member); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Member enrolled successfully")
			}

		case 2:
			id, ok := readLine(in, "Enter Member ID: ")
			if !ok {
				return
			}
			member, err := gym.LookupMember(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("\n--- MEMBER DETAILS ---")
			fmt.Println("ID:", member.MemberID)
			fmt.Println("Name:", member.FullName)
			fmt.Println("Plan:", member.PlanType)
			fmt.Println("Months:", member.MonthsRemaining)
			fmt.Println("Active:", member.IsActive)

		case 3:
			id, ok := readLine(in, "Enter Member ID: ")
			if !ok {
				return
			}
			monthsText, ok := readLine(in, "Enter Extra Months: ")
			if !ok {
				return
			}
			months, err := strconv.Atoi(monthsText)
			if err != nil {
				fmt.Println("Months must be a number.")
				continue
			}
			if err := gym.RenewSubscription(id, months); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Subscription renewed successfully")
			}

		case 4:
			id, ok := readLine(in, "Enter Member ID: ")
			if !ok {
				return
			}
			if err := gym.RevokeMembership(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Membership revoked successfully")
			}

		case 5:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
