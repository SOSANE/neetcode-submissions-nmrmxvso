func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    var answer []int
    for i, v := range(nums) {
        if a, ok := m[target-v]; ok {
            answer = append(answer, a, i)
        }
        m[v] = i
    }
    return answer
}
