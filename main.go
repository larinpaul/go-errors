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

// package main

// import (
// 	"fmt"
// )

// func sendSMSMToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
// 	cost, err := sendSMS(msgToCustomer)
// 	if err != nil {
// 		return 0.0, err
// 	}

// 	costSpouse, err := sendSMS(msgToSpouse)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return costSpouse + cost, nil
// }

// func sendSMS(message string) (float64, error) {
// 	const maxTextLen = 25
// 	const costPerChar = .0002
// 	if len(message) > maxTextLen {
// 		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
// 	}
// 	return costPerChar * float64(len(message)), nil
// }

// func test(msgToCustomer, msgToSpouse string) {
// 	defer fmt.Println("========")
// 	fmt.Println("Message for customer:", msgToCustomer)
// 	fmt.Println("Message for spouse:", msgToSpouse)
// 	totalCost, err := sendSMSMToCouple(msgToCustomer, msgToSpouse)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("Total cost: $%.4f\n", totalCost)
// }

// func main() {
// 	test(
// 		"Thanks for coming in to our flower shop today!",
// 		"We hope you enjoyed your gitf.",
// 	)
// 	test(
// 		"Thanks for joining us!",
// 		"Have a good day.",
// 	)
// 	test(
// 		"Thank you.",
// 		"Enjoy!",
// 	)
// 	test(
// 		"We loved having you in!",
// 		"We hope the rest of your evening is absolutely fanstastic.",
// 	)
// }

//// 2023/10/17 // 14:08 //
// Formatting strings review

// A convenient way to format strings in Go is by using the standard library's
// fmt.Sprintf() function. It's a string interpolation function, similar to
// JavaScript's built-in template literals. The %v substring uses the type's
// default formatting, which is often what you want.

// Default values

// const name = "Kim"
// const age = 22
// s := fmt.Sprintf("%v is %v years old.", name, age)
// // s = "Kim is 22 years old."

// // The equivalent JavaScript code:
// const name = 'Kim'
// const age = 22
// s = `${name} is ${age} years old.`
// // s = "Kim is 22 years old."

// // Rounding floats
// fmt.Printf("I am %f years old", 10.521)
// // I am 10.523000 years old

// // The ".2" rounds the number to 2 decimal places
// fmt.Printf("I am %.2f years old", 10.523)
// // I am 10.53 years old

//// 14:14
// Assignment
// We need better error logs for our backend developers to help them debug
// their code.

// Complete the getSMSErrorString() function. It should return a string
// with this format:
// SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
// * COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
// * RECIPIENT is the stringfield representation of the recipient's phone number
// Be sure to include the $ symbol and the single quotes

// package main

// import "fmt"

// func getSMSMErrorString(cost float64, recipient string) string {
// 	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",
// 		cost,
// 		recipient,
// 	)
// }

// func test(cost float64, recipient string) {
// 	s := getSMSMErrorString(cost, recipient)
// 	fmt.Println(s)
// 	fmt.Println("=====================================")
// }

// func main() {
// 	test(1.4, "+1 (435) 555 0923")
// 	test(2.1, "+2 (702) 555 3452")
// 	test(32.1, "+1 (801) 555 7456")
// 	test(14.4, "+1 (234) 555 6545")
// }

// //// 15:15

// // The Error Interface

// // Because errors are jsut itnerfaces, you can build your own custom types that implement the error
// // interface. Here's an example of a userError struct that implements the error interface:

// type userError struct {
// 	name string
// }

// func (e userError) Error() string {
// 	return fmt.Sprintf("%v has a problem with their account", e.name)
// }

// // It can then be used as an error:
// func sendSMS(msg, userName string) error {
// 	if !canSendToUser(userName) {
// 		return userError{name: userName}
// 	}
// 	...
// }

// Assignment
// Our users are frequently trying to run custom analytics queries on their message deliverability metrics, and
// end up writing a bad query that tries to divide a number by zero. It's become such a problem, that we thing
// it would be best to make a specific type of error for division by zero.

// Update the code so taht hte divideError type implements the error itnerface. Its Error() method
// should just return a string formatted in the following way:
/// can not divide DIVIDEND by zero
// Where DIVIDENT is the actual dividend of the divideError. User the %v verb to format the dividend as a
// float.

// package main

// import (
// 	"fmt"
// )

// type divideError struct {
// 	dividend float64
// }

// func (de divideError) Error() string {
// 	return fmt.Sprintf("can not divide %v by zero", de.dividend)
// }

// // don't edit below this line

// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		// We convert the `divideError` struct to an `error` type by returning it
// 		// as an error. As an error type, when it's printed its default value
// 		// will be the result of the Error() method
// 		return 0, divideError{dividend: dividend}
// 	}
// 	return dividend / divisor, nil
// }

// func test(dividend, divisor float64) {
// 	defer fmt.Println("=============================")
// 	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
// 	quotient, err := divide(dividend, divisor)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Quotient: %.2f\n", quotient)
// }

// func main() {
// 	test(10, 0)
// 	test(10, 2)
// 	test(15, 30)
// 	test(6, 3)
// }

// //// 15:35

// // The ERrors Package

// // The Go standard library provides an "errors" package that makes it easy to deal with errors.

// // Read the godoc for the errors.New() function, but here's a simple example:
// var err error = errors.New("something went wrong")

// Assignment

// Twilio's software architects may have overcomplicated the requirements from the last coding assignment...
// oops. All we needed was a new generic error message that returns the string no dividing by 0 when a
// user a ttempts to get us to perform the taboo.

// Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads
// "no dividng by 0".

package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func test(x, y float64) {
	defer fmt.Println("==============================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
