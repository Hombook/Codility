package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type nodeInfo struct {
	node       *Tree
	isFromLeft bool
	turnCount  int
}

func Solution(T *Tree) int {
	// There is no tree, return 0.
	if T == nil {
		return 0
	}

	// Find the first level of nodes, if any.
	var nextNodes []nodeInfo
	if T.L != nil {
		newNodeInfo := nodeInfo{T.L, true, 0}
		nextNodes = append(nextNodes, newNodeInfo)
	}
	if T.R != nil {
		newNodeInfo := nodeInfo{T.R, false, 0}
		nextNodes = append(nextNodes, newNodeInfo)
	}

	// Keep the largest turn record as we scan through the tree.
	largestTurns := 0
	// Loop through the node list until there is none.
	for nextNodes != nil {
		// This will hold the nodes on the next level.
		var toBeNextNodes []nodeInfo
		// Scan each node.
		for _, currentNodeInfo := range nextNodes {
			var isFromLeft = currentNodeInfo.isFromLeft
			var turnCount = currentNodeInfo.turnCount
			// Evaluate the left node.
			if currentNodeInfo.node.L != nil {
				var turnCountL = turnCount
				// It's parent should be a right node in order to count it a 'turn'.
				if !isFromLeft {
					turnCountL++
					if turnCountL > largestTurns {
						largestTurns = turnCountL
					}
				}
				// Put the node info in the next level list.
				// The child node will inherit the turn count.
				newLNodeInfo := nodeInfo{currentNodeInfo.node.L, true, turnCountL}
				toBeNextNodes = append(toBeNextNodes, newLNodeInfo)
			}
			// Evaluate the right node.
			if currentNodeInfo.node.R != nil {
				var turnCountR = turnCount
				// It's parent should be a left node in order to count it a 'turn'.
				if isFromLeft {
					turnCountR++
					if turnCountR > largestTurns {
						largestTurns = turnCountR
					}
				}
				newLNodeInfo := nodeInfo{currentNodeInfo.node.R, false, turnCountR}
				toBeNextNodes = append(toBeNextNodes, newLNodeInfo)
			}
		}
		// Hand over the collected nodes to our scanner.
		nextNodes = toBeNextNodes
	}

	return largestTurns
}
