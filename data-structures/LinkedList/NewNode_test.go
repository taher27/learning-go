// ********RoostGPT********
/*
Test generated by RoostGPT for test testingGoCoverage using AI Type Open AI and AI Model gpt-4

ROOST_METHOD_HASH=NewNode_382369d7a5
ROOST_METHOD_SIG_HASH=NewNode_60e8401887

================================VULNERABILITIES================================
Vulnerability: Missing Import Statement
Issue: The code is missing the import statement for the package. This can lead to runtime errors, as the types or functions that the code relies on may not be available.
Solution: Include the necessary import statements at the beginning of the code. In this case, you may need to import a package that defines the 'node' structure.

Vulnerability: Undefined Struct
Issue: The 'node' struct is not defined in the provided code. This can lead to compile-time errors.
Solution: Define the 'node' struct before using it in the function. The struct definition should be included in the same package, or imported from an external package.

================================================================================
Scenario 1: NewNode creation with positive integer value

Details:
  Description: This test is meant to check the correct creation of a new node when a positive integer value is passed to the NewNode function. The created node should have the val field set to the passed integer and the next and prev fields set to nil.
Execution:
  Arrange: No arrangement is needed for this test.
  Act: Invoke the NewNode function with a positive integer value.
  Assert: Use Go testing facilities to verify that the created node's val field matches the passed integer and the next and prev fields are nil.
Validation:
  The choice of assertion is based on the expected behavior of the NewNode function. The function is designed to create a new node with the passed value and with next and prev fields set to nil. This test is important to ensure that the basic functionality of the NewNode function works as expected.

Scenario 2: NewNode creation with negative integer value

Details:
  Description: This test is meant to check the correct creation of a new node when a negative integer value is passed to the NewNode function. The created node should have the val field set to the passed integer and the next and prev fields set to nil.
Execution:
  Arrange: No arrangement is needed for this test.
  Act: Invoke the NewNode function with a negative integer value.
  Assert: Use Go testing facilities to verify that the created node's val field matches the passed integer and the next and prev fields are nil.
Validation:
  The choice of assertion is based on the expected behavior of the NewNode function. The function is designed to create a new node with the passed value and with next and prev fields set to nil. This test is important to ensure that the NewNode function works correctly with negative integers.

Scenario 3: NewNode creation with zero value

Details:
  Description: This test is meant to check the correct creation of a new node when zero is passed to the NewNode function. The created node should have the val field set to zero and the next and prev fields set to nil.
Execution:
  Arrange: No arrangement is needed for this test.
  Act: Invoke the NewNode function with zero value.
  Assert: Use Go testing facilities to verify that the created node's val field is zero and the next and prev fields are nil.
Validation:
  The choice of assertion is based on the expected behavior of the NewNode function. The function is designed to create a new node with the passed value and with next and prev fields set to nil. This test is important to ensure that the NewNode function works correctly with zero value.
*/

// ********RoostGPT********
package LinkedList

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name string
		val  int
	}{
		{"Positive integer", 5},
		{"Negative integer", -5},
		{"Zero value", 0},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := NewNode(tc.val)

			// Assert
			if result.val != tc.val {
				t.Errorf("Expected value to be %v, but got %v", tc.val, result.val)
			}

			if result.next != nil {
				t.Errorf("Expected next to be nil, but got %v", result.next)
			}

			if result.prev != nil {
				t.Errorf("Expected prev to be nil, but got %v", result.prev)
			}
		})
	}
}
