Tactic:
1. Scan through the tree by level.
2. When evaluating the nodes on the current level, collect their child nodes to the "next level nodes" list. (3 -> 20 & 10 -> 1, 15.  NextLevelNodes: [20, 1, 15])
3. A turn will happen only if its parent is from the opposite direction. (10 to 1)
4. Turn count shall be passed to the child, as those turns have already happened when we hit this node.<br />
![alt text](tree.png?raw=true)