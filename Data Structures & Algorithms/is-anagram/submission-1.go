func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var record [26]int

    for i:=0;i<len(s);i++ {
        record[s[i]-'a']++
        record[t[i]-'a']--
    }

    for i:=0;i<26;i++ {
        if record[i] != 0{
            return false
        }
    } 

    return true
}
