# Graph Problems and Topological Sort: A Problem-Solving Guide

### Problem Type
- **Dependency Ordering Problem**: Tasks/items must be processed in an order that respects certain dependencies
- **Example**: Page 47 must be printed before Page 53 (written as 47|53)
- **Example**: Course A must be taken before Course B (A|B)
- **Example**: Library A must be compiled before Library B (A|B)
- **Example**: Task A must be executed before Task B (A|B) in workflow system

### Solution Approach
1. **Model as Graph**
   - Nodes: Individual pages
   - Directed Edges: Dependencies (if A|B, edge from A to B)

2. **Use Topological Sort**
   - Finds valid ordering that respects all dependencies
   - Detects if circular dependencies exist
   - Can validate if a given order is valid

## Key Graph Concepts

### 1. Graph Representation
- **Adjacency List**: Map where keys are nodes, values are lists of neighbors
  ```go
  graph := map[int][]int{
      1: [2,3],   // Node 1 points to nodes 2 and 3
      2: [4],     // Node 2 points to node 4
      3: [4],     // Node 3 points to node 4
      4: [],      // Node 4 has no outgoing edges
  }
  ```

### 2. In-Degree
- Count of incoming edges to each node
- Critical for identifying starting points
- Nodes with in-degree 0 have no dependencies
  ```go
  inDegree := map[int]int{
      1: 0,  // No nodes point to 1
      2: 1,  // One node points to 2
      3: 1,  // One node points to 3
      4: 2,  // Two nodes point to 4
  }
  ```

### 3. Directed Acyclic Graph (DAG)
- Graph with directed edges and no cycles
- Required for topological sort to work
- If cycle exists, no valid ordering is possible

## Common Problems Solvable with Topological Sort

1. **Build Systems**
   - Dependencies between source files
   - Compile order for libraries

2. **Course Prerequisites**
   - Students must take courses in correct order
   - Course A must be taken before Course B

3. **Task Scheduling**
   - Project tasks with dependencies
   - Assembly line ordering

4. **Package Management**
   - Software package dependencies
   - Installation order

5. **Data Processing Pipelines**
   - Steps must be executed in specific order
   - Data transformation dependencies

## Important Algorithm Patterns

### 1. Kahn's Algorithm Steps
```
1. Calculate in-degree for each node
2. Add nodes with in-degree 0 to queue
3. While queue not empty:
   - Remove node from queue
   - Add to result
   - Reduce in-degree of neighbors
   - Add neighbors with new in-degree 0 to queue
```

### 2. Cycle Detection
```
1. Count nodes in result
2. If count < total nodes, cycle exists
3. Alternative: Keep visited set during DFS (it's a cycle if node is visited again)
```

### 3. Validation Pattern
```
1. Create position map from sorted order
2. Check if each pair maintains relative position
3. For adjacent items A,B: position[A] < position[B]
```

## Problem-Solving Tips

1. **Modeling**
   - Identify nodes (individual items)
   - Identify edges (relationships/dependencies)
   - Draw the graph on paper first

2. **Edge Cases**
   - Empty graph
   - Single node
   - Disconnected components
   - Cycles
   - Multiple valid orders

3. **Optimization**
   - Use maps for O(1) lookups
   - Only process relevant nodes/edges
   - Avoid unnecessary validations

4. **Debugging**
   - Print graph structure
   - Track in-degree changes
   - Validate intermediate results

## Red Flags (When to Suspect Graph Problems)

1. Words/Phrases in Problem:
   - "depends on"
   - "must come before"
   - "prerequisites"
   - "ordering"
   - "sequence with conditions"

2. Problem Characteristics:
   - Items with relationships
   - Order matters
   - Dependencies between items
   - Validation of sequences
   - Finding valid arrangements

## Common Mistakes to Avoid

1. Not checking for cycles
2. Ignoring disconnected components in graph
3. Assuming single valid order
4. Not handling missing nodes
5. Incorrect graph construction
6. Over-constraining the problem
7. Not validating input
8. Assuming bi-directional relationships

Remember: Graph problems often appear in disguise. Look for relationships between items and dependencies in the problem statement. When in doubt, draw the relationships as a diagram - if it looks like a graph, it probably is one!


## Real-World Context

Consider these real scenarios where graph problems commonly appear:

1. **Software Deployment**
   ```
   ServiceA needs Database
   Database needs Configuration
   ServiceB needs ServiceA
   ```

2. **Recipe Steps**
   ```
   Beat eggs before adding flour
   Preheat oven before baking
   Mix dry ingredients before adding wet
   ```

3. **Document Processing**
   ```
   Sign form before submitting
   Get approval before signing
   Complete review before approval
   ```

## Common Mistakes with Examples

### 1. Not Checking for Cycles

#### Bad Example:
```go
// WRONG: No cycle detection
func processItems(dependencies map[string][]string) []string {
    var result []string
    queue := findNodesWithNoDependencies(dependencies)
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)
    }
    return result
}
```

#### Good Example:
```go
// RIGHT: With cycle detection
func processItems(dependencies map[string][]string) ([]string, error) {
    result := []string{}
    inDegree := calculateInDegree(dependencies)
    queue := findNodesWithNoDependencies(dependencies)
    
    nodesProcessed := 0
    totalNodes := len(dependencies)
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)
        nodesProcessed++
        
        for _, neighbor := range dependencies[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    if nodesProcessed != totalNodes {
        return nil, fmt.Errorf("cycle detected in dependencies")
    }
    return result, nil
}
```

### 2. Ignoring Missing Nodes

#### Bad Example:
```go
// WRONG: Assumes all nodes exist
func validateOrder(sequence []string, rules map[string][]string) bool {
    for node, deps := range rules {
        nodePos := findPosition(sequence, node)
        for _, dep := range deps {
            depPos := findPosition(sequence, dep)
            if nodePos < depPos {
                return false
            }
        }
    }
    return true
}
```

#### Good Example:
```go
// RIGHT: Handles missing nodes
func validateOrder(sequence []string, rules map[string][]string) bool {
    // Create position map
    positions := make(map[string]int)
    for i, node := range sequence {
        positions[node] = i
    }
    
    // Check only applicable rules
    for node, deps := range rules {
        nodePos, nodeExists := positions[node]
        if !nodeExists {
            continue // Skip rules for missing nodes
        }
        
        for _, dep := range deps {
            depPos, depExists := positions[dep]
            if !depExists {
                continue // Skip if dependent node is missing
            }
            
            if nodePos < depPos {
                return false
            }
        }
    }
    return true
}
```

### 3. Over-constraining Problems

#### Bad Example:
```go
// WRONG: Forces strict ordering even when unnecessary
func isValidSequence(seq []int, sorted []int) bool {
    orderMap := make(map[int]int)
    for i, v := range sorted {
        orderMap[v] = i
    }
    
    // Forces every pair to follow global order
    for i := 0; i < len(seq)-1; i++ {
        if orderMap[seq[i]] > orderMap[seq[i+1]] {
            return false
        }
    }
    return true
}
```

#### Good Example:
```go
// RIGHT: Only enforces necessary constraints
func isValidSequence(seq []int, rules []Rule) bool {
    positions := make(map[int]int)
    for i, v := range seq {
        positions[v] = i
    }
    
    // Only check pairs with actual dependencies
    for _, rule := range rules {
        posA, hasA := positions[rule.before]
        posB, hasB := positions[rule.after]
        
        if hasA && hasB && posA > posB {
            return false
        }
    }
    return true
}
```

## Common Gotchas and Solutions

### 1. Hidden Cycles

```
A → B → C → A   (obvious cycle)
A → B → C → D → B   (subtle cycle)
```

Solution:
```go
func detectCycle(graph map[string][]string) bool {
    visited := make(map[string]bool)
    recStack := make(map[string]bool)
    
    var dfs func(node string) bool
    dfs = func(node string) bool {
        visited[node] = true
        recStack[node] = true
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                if dfs(neighbor) {
                    return true
                }
            } else if recStack[neighbor] {
                return true
            }
        }
        
        recStack[node] = false
        return false
    }
    
    for node := range graph {
        if !visited[node] {
            if dfs(node) {
                return true
            }
        }
    }
    return false
}
```

### 2. Partial Ordering Traps

Problem:
```
Rules: A→B, C→D
Invalid Assumption: A must come before C
```

Solution:
```go
func buildDependencyGroups(rules []Rule) map[int][]int {
    groups := make(map[int][]int)
    
    // Track which nodes are related
    connected := make(map[int]map[int]bool)
    
    // Build connected components
    for _, rule := range rules {
        if _, exists := connected[rule.before]; !exists {
            connected[rule.before] = make(map[int]bool)
        }
        if _, exists := connected[rule.after]; !exists {
            connected[rule.after] = make(map[int]bool)
        }
        connected[rule.before][rule.after] = true
        connected[rule.after][rule.before] = true
    }
    
    // Now we can properly group related nodes
    groupID := 0
    processed := make(map[int]bool)
    
    for node := range connected {
        if !processed[node] {
            group := findConnectedNodes(node, connected, processed)
            groups[groupID] = group
            groupID++
        }
    }
    
    return groups
}
```

### 3. Edge Case Handling

Problem:
```
Empty input
Single node
All nodes independent
```

Solution:
```go
func processGraph(nodes []string, edges []Edge) ([]string, error) {
    // Handle empty input
    if len(nodes) == 0 {
        return []string{}, nil
    }
    
    // Handle single node
    if len(nodes) == 1 {
        return nodes, nil
    }
    
    // Handle independent nodes
    if len(edges) == 0 {
        return nodes, nil // Any order is valid
    }
    
    // Build graph
    graph := buildGraph(nodes, edges)
    
    // Process normal case
    return topologicalSort(graph)
}
```

## Best Practices Checklist

1. **Input Validation**
   - Check for empty input
   - Validate edge format
   - Verify node existence

2. **Graph Construction**
   - Initialize all nodes, even without edges
   - Build bidirectional maps for quick lookup
   - Track both incoming and outgoing edges

3. **Algorithm Implementation**
   - Include cycle detection
   - Handle disconnected components
   - Track visited nodes properly

4. **Result Validation**
   - Verify all nodes included
   - Check constraint satisfaction
   - Validate output format

Remember: The key to avoiding mistakes is systematic testing with different edge cases and maintaining clear data structures that make the state of your algorithm visible and debuggable.