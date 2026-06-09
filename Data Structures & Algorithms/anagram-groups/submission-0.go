func buildSign(word string) [26]int {
    var sign [26]int

    for i:=0; i<len(word); i++{
        sign[word[i]-'a']++
    }

    return sign
}

func groupAnagrams(strs []string) [][]string {
    group := map[[26]int][]string{}

    for _, word:= range strs {
        sign := buildSign(word)

        _, found := group[sign]
        if !found {
            group[sign] = []string{}
        }

        group[sign] = append(group[sign], word)
    }

    groups := [][]string{}
    for _, val := range group {
        groups = append(groups, val)
    }

    return groups
}
