func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var filter [26]int

    for i := range s{
        filter[s[i] - 'a']++
        filter[t[i] - 'a']--
    }

    for _,v := range filter {
        if v != 0 {
            return false
        }
    }

    return true
}
