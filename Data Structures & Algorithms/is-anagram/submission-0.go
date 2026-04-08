func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
 
    so := make(map[rune]int)
    to := make(map[rune]int)

    for i, v := range(s) {
        so[v]++
        to[rune(t[i])]++
    }
    
    for k, v := range(so) {
        if v != to[k] {
            return false
        }
    }

    return true
}
