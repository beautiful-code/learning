package alerting

import "fmt"

func SendWelcomeEmail(email string) {
	fmt.Printf("Successfully sent out welcome email to %s", email)
}
