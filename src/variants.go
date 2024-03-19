package main

import "slices"

func GenerateVariants(config *Config) []map[string]string {
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
	// TODO check if all keys are specified

	// TODO add possibility to remove multiple entries at once
	for _, exclude := range excludes {
		for i, variant := range variants {
			found := true

			for k, v := range exclude {
				if variant[k] != v {
					found = false
					break
				}
			}

			if found {
				// this might not be the most efficient way to delete an entry;
				// order is retained
				variants = slices.Delete(variants, i, i+1)
			}
		}
	}

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
