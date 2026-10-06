// Copied Code for testing out time

import "slices"

func groupAnagrams(strs []string) [][]string {    
    mapped := make(map[string][]string)
    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        mapped[string(runes)] = append(mapped[string(runes)], str)
    }

    res := [][]string{}
    for _, m := range mapped {
        res = append(res, m)
    }
    return res
}
