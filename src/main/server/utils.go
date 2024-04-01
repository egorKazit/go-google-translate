package server

func getNested(nestedValue any, path []int) any {
	if _, ok := nestedValue.([]interface{}); !ok {
		return nil
	}
	// covert to nested type
	results := nestedValue.([]interface{})
	var result any
	// get length
	pathLen := len(path) - 1
	// get through path
	for index, value := range path {
		// if not last index -> handle results
		if pathLen > index {
			if len(results) < value {
				return nil
			}
			// re-covert to middle
			resultsIntermediate := results[value]
			if _, ok := resultsIntermediate.([]interface{}); ok {
				// reassign to results
				results = resultsIntermediate.([]interface{})
			} else {
				return nil
			}

		} else {
			// set result in case of last index
			if results == nil || len(results) <= value {
				return nil
			}
			result = results[value]
		}
	}
	return result
}
