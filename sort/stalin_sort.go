package sort

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
