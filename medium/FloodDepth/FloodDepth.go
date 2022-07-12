package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	var peakStart = 0
	var holePart = 0
	var lowestHeight = 100000001
	var deepestDepth = 0

	for i, height := range A {
		// A plane that's lower that the first peak: A potential flooding spot
		if peakStart > height {
			// Check if it's end of area
			if i == len(A)-1 {
				// End of area, check if it's relatively higher than the lowest part,
				// which will create flooded spots.
				var maxDepth = height - lowestHeight
				if maxDepth > deepestDepth {
					deepestDepth = maxDepth
				}
			} else {
				// Counting the deepest hole so far
				if holePart > 0 {
					var maxDepth = height - lowestHeight
					if maxDepth > deepestDepth {
						deepestDepth = maxDepth
					}
				}
				if height < lowestHeight {
					lowestHeight = height
				}
				holePart++
			}
		} else { // Another peak that close the flooding area.
			// If it's not peak to peak, there will be flooded spots.
			if holePart > 0 {
				var maxDepth = peakStart - lowestHeight
				if maxDepth > deepestDepth {
					deepestDepth = maxDepth
				}
			}

			// new starting point
			peakStart = height
			holePart = 0
			lowestHeight = 100000001
		}
	}

	return deepestDepth
}
