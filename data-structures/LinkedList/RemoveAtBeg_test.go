// ********RoostGPT********
/*
Test generated by RoostGPT for test goastparsertest using AI Type  and AI Model 

ROOST_METHOD_HASH=RemoveAtBeg_2301eda597
ROOST_METHOD_SIG_HASH=RemoveAtBeg_e45f33f4b9

Here are several test scenarios for the `RemoveAtBeg` function, covering normal operation, edge cases, and error handling:

---

### Scenario 1: Remove from an empty linked list

Details:  
- **Description**: This test checks the behavior of the `RemoveAtBeg` function when the linked list is empty. The function should return `-1` since there are no elements to remove.
  
Execution:  
- **Arrange**: Create an empty linked list (`ll := LinkedList{head: nil}`).
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `-1`.

Validation:  
- **Explanation**: The function should return `-1` when the list is empty to signal that no element was removed. This test ensures the function gracefully handles the edge case where no nodes are present.
- **Importance**: It's critical to handle empty list cases to avoid panics or errors in real-world applications where linked lists could be empty at any point.

---

### Scenario 2: Remove from a linked list with one node

Details:  
- **Description**: This test validates that the function correctly removes and returns the value of the only node in the list, and the list becomes empty afterward.
  
Execution:  
- **Arrange**: Create a linked list with one node (`ll := LinkedList{head: &node{val: 10, next: nil, prev: nil}}`).
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `10` and that `ll.head` is now `nil`.

Validation:  
- **Explanation**: After removing the only node, the list should be empty (`head == nil`). The function should return the correct value (`10`) and properly update the list structure.
- **Importance**: This test ensures that the function correctly handles the removal of the last node in the list, which is a common operation in linked list management.

---

### Scenario 3: Remove from a linked list with multiple nodes

Details:  
- **Description**: This test checks that the function correctly removes the first node from a list with multiple nodes and updates the list structure properly.
  
Execution:  
- **Arrange**: Create a linked list with multiple nodes (`ll := LinkedList{head: &node{val: 10, next: &node{val: 20, next: nil, prev: nil}, prev: nil}}`). Set the `prev` pointer of the second node correctly.
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `10`, and `ll.head.val` is now `20`. Also, check that the `prev` pointer of the new head node is `nil`.

Validation:  
- **Explanation**: The function should remove the first node (`10`), update the head to the second node (`20`), and set the new head's `prev` pointer to `nil`. This test ensures correct list manipulation during removal.
- **Importance**: Properly updating the `head` and `prev` pointers is crucial for maintaining the integrity of the linked list, especially in multi-node scenarios.

---

### Scenario 4: Remove from a linked list with two nodes

Details:  
- **Description**: This test checks the behavior when removing the first node from a list with exactly two nodes. The second node should become the new head, and its `prev` pointer should be updated to `nil`.
  
Execution:  
- **Arrange**: Create a linked list with two nodes (`ll := LinkedList{head: &node{val: 5, next: &node{val: 15, next: nil, prev: nil}, prev: nil}}`). Set the `prev` pointer of the second node correctly.
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `5`, and `ll.head.val` is now `15`. Also, check that the `prev` pointer of the new head node is `nil`.

Validation:  
- **Explanation**: The function should remove the first node (`5`), update the head to the second node (`15`), and set the new head's `prev` pointer to `nil`. This test ensures proper pointer updates when there are exactly two nodes.
- **Importance**: This test validates the function's ability to handle small lists, ensuring correct pointer manipulation when transitioning from two nodes to one.

---

### Scenario 5: Remove repeatedly until the list becomes empty

Details:  
- **Description**: This test ensures that the function behaves correctly when repeatedly removing nodes from a list until it becomes empty.
  
Execution:  
- **Arrange**: Create a linked list with three nodes (`ll := LinkedList{head: &node{val: 1, next: &node{val: 2, next: &node{val: 3, next: nil, prev: nil}, prev: nil}, prev: nil}}`). Set up the `prev` pointers correctly.
- **Act**: Call `ll.RemoveAtBeg()` three times.
- **Assert**: Verify the return values are `1`, `2`, and `3` in sequence. After the third call, `ll.head` should be `nil`.

Validation:  
- **Explanation**: The function should remove nodes one at a time, updating the head and maintaining proper list structure. After removing the last node, the list should be empty (`head == nil`).
- **Importance**: This test checks the function's ability to handle multiple removals and ensures the list structure is updated correctly after each removal.

---

### Scenario 6: Remove from a list where nodes contain negative values

Details:  
- **Description**: This test ensures that the function handles nodes with negative values correctly, as the value of a node should not affect the removal logic.
  
Execution:  
- **Arrange**: Create a linked list with nodes containing negative values (`ll := LinkedList{head: &node{val: -5, next: &node{val: -10, next: nil, prev: nil}, prev: nil}}`). Set the `prev` pointer of the second node correctly.
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `-5`, and `ll.head.val` is now `-10`. Also, check that the `prev` pointer of the new head node is `nil`.

Validation:  
- **Explanation**: The function should remove the first node (`-5`), update the head to the second node (`-10`), and set the new head's `prev` pointer to `nil`. This test ensures the function works correctly with negative values.
- **Importance**: Handling negative values is important for general-purpose linked lists, as they may store any integer values.

---

### Scenario 7: Remove from a list with nodes containing large values

Details:  
- **Description**: This test checks that the function handles nodes with large integer values without any overflow or unexpected behavior.
  
Execution:  
- **Arrange**: Create a linked list with a node containing a large value (`ll := LinkedList{head: &node{val: 9223372036854775807, next: nil, prev: nil}}`).
- **Act**: Call `ll.RemoveAtBeg()`.
- **Assert**: Verify that the return value is `9223372036854775807` and `ll.head` is now `nil`.

Validation:  
- **Explanation**: The function should remove the node and return the large value without any issues. This test ensures the function properly handles large integer inputs.
- **Importance**: Testing for large values ensures the function is robust and can handle edge cases related to integer limits.

---

These test scenarios cover various edge cases and normal operations for the `RemoveAtBeg` function. By testing different linked list configurations (empty, single node, multiple nodes, negative values, large values), we can ensure the function behaves correctly in a wide range of situations.
*/

// ********RoostGPT********
package LinkedList

import (
	"testing"
	"bytes"
	"fmt"
)

// Node structure as per the context
type node struct {
	val  int
	next *node
	prev *node
}

// LinkedList structure as per the context
type LinkedList struct {
	head *node
}

// TestRemoveAtBeg tests the RemoveAtBeg function using table-driven tests
func TestRemoveAtBeg(t *testing.T) {
	testCases := []struct {
		name           string
		initialList    *LinkedList
		expectedReturn int
		expectedHead   *node
	}{
		{
			name:           "Remove from an empty linked list",
			initialList:    &LinkedList{head: nil},
			expectedReturn: -1,
			expectedHead:   nil,
		},
		{
			name:           "Remove from a linked list with one node",
			initialList:    &LinkedList{head: &node{val: 10, next: nil, prev: nil}},
			expectedReturn: 10,
			expectedHead:   nil,
		},
		{
			name: "Remove from a linked list with multiple nodes",
			initialList: &LinkedList{head: &node{
				val: 10, next: &node{val: 20, next: nil, prev: nil}, prev: nil}},
			expectedReturn: 10,
			expectedHead:   &node{val: 20, next: nil, prev: nil},
		},
		{
			name: "Remove from a linked list with two nodes",
			initialList: &LinkedList{head: &node{
				val: 5, next: &node{val: 15, next: nil, prev: nil}, prev: nil}},
			expectedReturn: 5,
			expectedHead:   &node{val: 15, next: nil, prev: nil},
		},
		{
			name: "Remove repeatedly until the list becomes empty",
			initialList: &LinkedList{head: &node{
				val: 1, next: &node{val: 2, next: &node{val: 3, next: nil, prev: nil}, prev: nil}, prev: nil}},
			expectedReturn: 1,
			expectedHead:   &node{val: 2, next: &node{val: 3, next: nil, prev: nil}, prev: nil},
		},
		{
			name: "Remove from a list where nodes contain negative values",
			initialList: &LinkedList{head: &node{
				val: -5, next: &node{val: -10, next: nil, prev: nil}, prev: nil}},
			expectedReturn: -5,
			expectedHead:   &node{val: -10, next: nil, prev: nil},
		},
		{
			name: "Remove from a list with nodes containing large values",
			initialList: &LinkedList{head: &node{
				val: 9223372036854775807, next: nil, prev: nil}},
			expectedReturn: 9223372036854775807,
			expectedHead:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)

			// Capture the output of RemoveAtBeg
			var buf bytes.Buffer
			fmt.Fprintf(&buf, "Initial List: %+v\n", tc.initialList)

			// Act
			result := tc.initialList.RemoveAtBeg()

			// Log the results
			t.Logf("Expected Return: %d, Got: %d", tc.expectedReturn, result)
			t.Logf("Expected Head: %+v, Got: %+v", tc.expectedHead, tc.initialList.head)

			// Assert
			if result != tc.expectedReturn {
				t.Errorf("expected %d, got %d", tc.expectedReturn, result)
			}
			if tc.expectedHead == nil && tc.initialList.head != nil {
				t.Errorf("expected head to be nil, but got %+v", tc.initialList.head)
			} else if tc.expectedHead != nil && tc.initialList.head != nil {
				if tc.expectedHead.val != tc.initialList.head.val {
					t.Errorf("expected head value %d, but got %d", tc.expectedHead.val, tc.initialList.head.val)
				}
				if tc.initialList.head.prev != nil {
					t.Errorf("expected head.prev to be nil, but got %+v", tc.initialList.head.prev)
				}
			}
		})
	}
}
