func isAnagram(s string, t string) bool {
    mp := map[rune]int{}

    for _,byt := range s {
        _, found := mp[byt]
        if !found {
            mp[byt] = 0
        }

        mp[byt]++
    }

    for _,byt := range t {
        val, found := mp[byt]
        if !found {
            return false
        }

        mp[byt] = val-1

        if val-1 == 0{
            delete(mp, byt)
        }
    }

    return len(mp) == 0
}
