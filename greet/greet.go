package greet

import "errors"

// We can use getGreeting in the same package only.
func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		err := errors.New("too early for greetings")
		return message, err
	}

	if hour < 12 {
		message = "Good Morning!"
	} else if hour < 6 {
		message = "Good Afternoon!"
	} else {
		message = "Good Evening!"
	}
	return message, nil
}

// GetGreeting allows to use getGreeting outside the package.
func GetGreeting(hour int) (string, error) {
	return getGreeting(hour)
}
