package interview_problems

type Node struct {
	Value int
	Next  *Node
}

func CreateList() *Node {

	// Creating a circular linked list for testing
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}
	node5 := &Node{Value: 5}
	node6 := &Node{Value: 6}
	node7 := &Node{Value: 7}
	node8 := &Node{Value: 8}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node4 // Creating the loop

	return node1
}

/*
this is the solution i gave in interview
although I have not completed the code fully
so this is time complexcity O(n) space complextiy O(n)
*/
func IsCircularUsingMap() bool {

	head := CreateList()

	visited := make(map[*Node]bool)
	current := head

	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}

	return false

}

/*
actual expected answer
time complexcity O(n) space complextiy O(1)

	This code defines a linked list and checks if it is circular using the Floyd's Cycle-Finding Algorithm
	(also known as the Tortoise and Hare algorithm).
	The isCircular function returns true if the linked list is circular and false otherwise.
	The main function creates a circular linked list and tests the isCircular function.
*/
func IsCircularFloydCycleFindingAlgorithm(head *Node) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
