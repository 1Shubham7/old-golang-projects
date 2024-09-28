package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    str := strings.ToUpper(customer)
    return "Welcome to the Tech Palace, " + str
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
    msg := welcomeMsg
    formattedString := fmt.Sprintf("%s\n%s\n%s", stars,msg, stars)
    return formattedString
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    newMsg := strings.Replace(oldMsg, "*", "", -1)
    trimedSpaces := strings.TrimSpace(newMsg)
    return trimedSpaces
}
