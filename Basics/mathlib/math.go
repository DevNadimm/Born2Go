package mathlib

// Add is a public function because it starts with an uppercase letter.
// It can be accessed from other packages.
func Add(num1, num2 int) int {
	return num1 + num2
}

// add is a private function because it starts with a lowercase letter.
// It is accessible only within the same package.
func multiply(num1, num2 int) int {
	return num1 * num2
}
