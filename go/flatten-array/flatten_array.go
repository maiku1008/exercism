package flatten

// Flatten takes an input and flattens it
func Flatten(input interface{}) []interface{} {
	output := []interface{}{}

	for _, v := range input.([]interface{}) {
		switch v := v.(type) {
		case []interface{}:
			output = append(output, Flatten(v)...)
		case interface{}:
			output = append(output, v)
		}
	}
	return output
}
