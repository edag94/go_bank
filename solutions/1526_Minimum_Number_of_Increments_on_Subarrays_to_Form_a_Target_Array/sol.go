package Minimum_Number_of_Increments_on_Subarrays_to_Form_a_Target_Array

/*
	Even though this is a hard question, the concept is simple

	Go from left to right, and for as long as the number to the right is smaller,
	you can include that in a subarray with the previous element, and therefore increment the sub array together.
	You'll still need as many operations as the max element, which is the leftmost element in the subarray, as if its
	not the max, you would just start a new sub array.
	Here's an example:

	[3,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]

	The answer here is 3. You can increment the whole array once, but the leftmost needs 3. Only if there's a next element that is greater, that things change.const

	[3,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2]

	Here, the answer is 4. Since there is a number at the end greater than the previous, at some point, it will be treated as its own subarray.
*/
func minNumberOperations(target []int) int {
	var operationCount int = target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			operationCount += target[i] - target[i-1]
		}
	}
	return operationCount
}
