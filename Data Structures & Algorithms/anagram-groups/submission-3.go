

func groupAnagrams(strs []string) [][]string {
    group := map[[26]int][]string{}

    byt_a := byte('a')

    buildSign := func(s string)  [26]int {
        var sign [26]int
        for _ , char := range s{
            sign[byte(char) - byt_a]++
        }
        return sign
    }

    for _, str := range strs {
        getSign := buildSign(str)
        group[getSign] = append(group[getSign], str)
    }

    grouplist := make([][]string, len(group), len(group))

    idx := 0
    for _,val := range group{
        grouplist[idx]= val
        idx++
    }

    return grouplist
}
