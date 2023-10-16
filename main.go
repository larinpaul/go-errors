//// 2023/10/16 // 21:57 //

// // The Error Interface

// // Go pograms express errors with error values. An Error is any type that implements the simple built-int
// // error interface:
// type error interface {
// 	Error() string
// }

// // When something can go wrong in a function, that function should return an error as its last return value.
// // Any code that calls a function that can return an error should handle errors by testing whether the error is
// // nil.

// // Atoi converts a stringified number to an integer
// i, err := strconv.Aroi("42b")
// if err != nil {
// 	fmt.Println("couldn't convert:", err)
// 	// because "42b" isn't a valid integer, we print:
// 	// couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
// 	// Note:
// 	// 'parsing "42b": invalid syntax' is returned by the .Error() method
// 	return
// }
// // if we get here, then
// // i was converted successfully

// // A nil error denotes success; a non-nil error denotes failure.

//// Assignment

// We offer a product that allows businesses that use Textio to send pairs of messages to couples. It is mostly
// used by flower shops and movie theaters.

// Complete the sendSMSMToCouple function. It sohuld send 2 messages, first to the customer, then to the
// customer's spouse.

// 1. Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0.0 and the error.
// 2. Do the same sfor the msgToSpouse
// 3. If both messages are sent successfully, retunr the total cost of the messages added together.

// When you return a non-nil error in Go, it's coventional to return the "zero" values of all otehr return values.

package main

import (
	"fmt"
)

func sendSMSMToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return costSpouse + cost, nil
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSMToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gitf.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fanstastic.",
	)
}
