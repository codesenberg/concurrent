package histogram

import "math/rand"

func generateStrings(seed int64, count, size int) []string {
	rand.Seed(seed)
	result := make([]string, count)
	randomBytes := make([]byte, size)
	for i := range result {
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
		// Make it a string from some random letters.
		for j := range randomBytes {
			randomBytes[j] = 'a' + (randomBytes[j] % 28)
		}
		result[i] = string(randomBytes)
	}
	return result
}
