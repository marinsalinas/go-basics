package greet

// We can use getGreeting in the same package only.
func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning!"
	} else if hour < 6 {
		return "Good Afternoon!"
	} else {
		return "Good Evening!"
	}
}

// GetGreeting allows to use getGreeting outside the package.
func GetGreeting(hour int) string {
	return getGreeting(hour)
}
