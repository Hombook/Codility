package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type sockCount struct {
	clean int
	dirty int
}

func Solution(K int, C []int, D []int) int {
	// key: sock ID, value: sock count
	var sockCountMap = make(map[int]sockCount)
	var cleanPairs = 0
	var dirtyPairs = 0
	var washesLeft = K

	// Pair the clean socks first.
	for _, cleanSock := range C {
		var cleanSockCount = sockCountMap[cleanSock].clean
		cleanSockCount++
		if cleanSockCount == 2 {
			cleanPairs++
			cleanSockCount = 0
		}

		newSockCount := sockCount{cleanSockCount, 0}
		sockCountMap[cleanSock] = newSockCount
	}

	// No chance to clean any dirty socks, return directly
	if K <= 0 {
		return cleanPairs
	}

	// Find any possible 1-clean 1-dirty pairs.
	for _, dirtySock := range D {
		var cleanSockCount = sockCountMap[dirtySock].clean
		var dirtySockCount = sockCountMap[dirtySock].dirty
		dirtySockCount++
		// If we have 1 clean sock left, use 1 wash quota for the dirty 1.
		if cleanSockCount > 0 {
			washesLeft--
			cleanPairs++
			cleanSockCount = 0
			dirtySockCount = 0
		} else {
			// Collect the dirty pairs.
			if dirtySockCount == 2 {
				dirtyPairs++
				dirtySockCount = 0
			}
		}
		// No more washes left, there is no point of finding more dirty socks.
		if washesLeft <= 0 {
			break
		}
		newSockCount := sockCount{cleanSockCount, dirtySockCount}
		sockCountMap[dirtySock] = newSockCount
	}

	// Wash the dirty pairs.
	// Each pair requires 2 wash quotas.
	for washesLeft > 1 {
		if dirtyPairs > 0 {
			washesLeft = washesLeft - 2
			cleanPairs++
			dirtyPairs--
		} else {
			break
		}
	}

	return cleanPairs
}
