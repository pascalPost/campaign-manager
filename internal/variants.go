package internal

func GenerateVariants(config *Input) []map[string]string {
	variants := generateFromMatrix(config.Matrix)
	variants = removeExcludes(variants, config.Exclude)
	variants = addIncludes(variants, config.Include)

	return variants
}

// generateVariants generates all variants of a matrix
func generateFromMatrix(matrix map[string][]string) []map[string]string {
	// This code generates variants using an iterative approach.
	// It iterates through all combinations efficiently without
	// using recursion, thus offering better performance,
	// especially for larger matrices.

	// TODO check the matrix input for uniqueness (values & keys)

	// TODO add allocations to enhance performance

	var keys []string
	var keyIndex []int
	for k := range matrix {
		keys = append(keys, k)
		keyIndex = append(keyIndex, 0)
	}

	var result []map[string]string

	for {
		var variant = make(map[string]string)
		for i, key := range keys {
			variant[key] = matrix[key][keyIndex[i]]
		}
		result = append(result, variant)

		incrementIndex := len(keys) - 1
		for incrementIndex >= 0 {
			keyIndex[incrementIndex]++
			if keyIndex[incrementIndex] < len(matrix[keys[incrementIndex]]) {
				break
			}
			keyIndex[incrementIndex] = 0
			incrementIndex--
		}
		if incrementIndex < 0 {
			break
		}
	}

	return result
}

func removeExcludes(variants []map[string]string, excludes []map[string]string) []map[string]string {
	// parse all variants and collect the ones that match the exclude map
	var idx2delete []uint
	for _, exclude := range excludes {
		for i, variant := range variants {
			for k, v := range exclude {
				if variant[k] == v {
					idx2delete = append(idx2delete, uint(i))
				}
			}
		}
	}

	//delete all marked variants

	// TODO create own class for variants that contains all keys.
	// TODO use the keys to check that all given exclude keys are valid.

	//variants = slices.DeleteFunc(variants, func(variant map[string]string) bool {
	//	for _, exclude := range excludes {
	//		for k, v := range exclude {
	//			if variant[k] != v {
	//				return false
	//			}
	//		}
	//		return true
	//	}
	//	return false
	//})

	return variants
}

func addIncludes(variants []map[string]string, includes []map[string]string) []map[string]string {

	// TODO check if all keys are specified
	// TODO check if all values are unique

	// check the includes map for validity
	for _, include := range includes {
		variants = append(variants, include)
	}

	return variants
}
