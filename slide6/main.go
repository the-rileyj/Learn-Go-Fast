package java_sucks

import "fmt"

func does_java_suck() bool {
	return true
}

func get_good_question() (string, error) {
	if does_java_suck() {
		return "Then why is it still used?", nil
	}
	return "Error in function", fmt.Errorf("Function broken")
}
