package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🌞 Welcome to your Gratitude Journal 🌸")
	fmt.Println("1. Write a new gratitude entry")
	fmt.Println("2. View my past entries")
	fmt.Print("Choose 1 or 2: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		writeEntry(reader)
	} else if choice == "2" {
		viewEntries()
	} else {
		fmt.Println("Oops! I only understand 1 or 2. Try again next time 😅")
	}
}

func writeEntry(reader *bufio.Reader) {
	fmt.Print("What are you thankful for today? ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	now := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s - %s\n", now, input)

	file, err := os.OpenFile("journal.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Failed to write:", err)
	}

	fmt.Println("💖 Saved! You’re doing amazing 🙌🏽")
}

func viewEntries() {
	data, err := os.ReadFile("journal.txt")
	if err != nil {
		fmt.Println("Hmm… couldn’t read your journal. Maybe it’s empty? 🧐")
		return
	}

	fmt.Println("\n📔 Your Gratitude Entries:")
	fmt.Println("----------------------------")
	fmt.Println(string(data))
}
