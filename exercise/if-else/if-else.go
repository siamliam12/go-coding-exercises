// --Summary:
// The existing program is usedf to restrict access to a resource base on dy of the week and user role. Currently, the program allows any user to access the resource. Implement the functionality needed to grant resource access using any combination of `if`, `else if`, and `else`.

// --Requirements:
// *Use the accessGranted() and accessDenided() functions to display informational messages
// *Acess at any time : Admin,Manager
// *Acess Weekends: Contractor
// *Access weekdays : Member
// *Acess Mondays,Wednesdays and Fridays : Guest

package main

import "fmt"

//days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

//user Roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func weekDay(day int) bool {
	return day <= 4
}
func accessGranted() {
	fmt.Println("Granted")
}
func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	fmt.Println("Starting if-else excercise...")
	today, role := Tuesday, Guest
	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekDay(today) {
		accessGranted()
	} else if role == Member && weekDay(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}

}
