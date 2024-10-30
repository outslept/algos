package stalinsort

/*
Stalin Sort is a humorous sorting algorithm that:
1. Iterates through array from left to right
2. If current element is less than previous - remove it
3. Else - eliminate elements that are out of order

Name after Stalin's purges, as it "removes" elements that don't follow the increasing order.

Visual representation:

    Initial array:  [1,2,4,3,6,5,7]
    Process:         ^ ^ ^ x ^ x ^
    Result:        [1,2,4,6,7]

    where x marks eliminated elements
*/


func StalinSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] >= result[len(result)-1] {
			result = append(result, arr[i])
		}
	}
	return result
}

/*
Time Complexity: O(n), where n is the number of elements

SpaceComplexity: O(n) in worst case
*/
