package helper

func Messages() map[string]interface{} {
	// Create a map to store frequently used variables
	frequentVars := make(map[string]interface{})

	// Store variables in the map
	frequentVars["invalid_password"] = "invalid_password"

	return frequentVars
}
