import (
    "maps"
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    group := map[[26]int][]string{}

    byt_a := byte('a')
    for _, str := range strs {
        var sign [26]int

        for _ , char := range str{
            sign[byte(char) - byt_a]++
        }

        group[sign] = append(group[sign], str)
    }

    return slices.Collect(maps.Values(group))
}
