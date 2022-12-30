package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number, because it is divisible by 2"},
		{"not prime by definition", 0, false, "0 is not a prime, by definition"},
		{"not prime by definition", 1, false, "1 is not a prime, by definition"},
		{"negative number cannot be prime", -1, false, "Negative numbers are not prime by definition"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if msg != e.msg {
			t.Errorf("Wrong message returned: %s", msg)
		}
	}

}
