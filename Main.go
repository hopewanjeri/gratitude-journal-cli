package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🌞 Welcome to your Gratitude Journal 🌸")
	fmt.Print("What are you thankful for today? ")

	input, _ := reader.ReadString('\n')

	// Format the entry
	now := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s - %s\n", now, input)

	// Save to file
	file, err := os.OpenFile("journal.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Failed to write:", err)
	}

	fmt.Println("💖 Saved! Have a beautiful day ✨")
}
