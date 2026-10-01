func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    filter := map[byte]int{}

    for i:=0; i<len(s); i++{
        if _ , ex := filter[s[i]]; !ex {
            filter[s[i]] = 1
        } else {
            filter[s[i]]+=1
        }

        if _ , ex := filter[t[i]]; !ex {
            filter[t[i]] = -1
        } else {
            filter[t[i]]-=1
        }
    }

    for _,v := range filter {
        if v != 0 {
            return false
        }
    }

    return true
}
