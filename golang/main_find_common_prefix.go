package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	count := lengthOfLongestSubstring([]string{"flower", "flow", "flight"})
}

func longestCommonPrefix(strs []string) string {
	longestPrefix := []string{}

	if len(strs) == 1 {
		ind := 0
		if len(strs[0]) == 1 {
			ind = 1
		} else if len(strs[0]) == 2 {
			ind = 2
		} else if len(strs[0]) == 1 {
			ind = 3
		}

		return strs[0][:ind]
	}

	for coreIndex, s := range strs {
		tempPrefix := ""
		// if coreIndex > 0 {
		//     continue
		// }
		for i, _ := range s {
			// if i > 2 {
			// 	continue
			// }
			_ = i

			for nestedIndex, nextStr := range strs {

				_ = coreIndex
				_ = nestedIndex
				for newStrCharIndex, _ := range nextStr {
					// if len(nextStr) == 1 && len(strs[0]) == 1 {
					//     continue
					// }
					// if len(s) >= newStrCharIndex {
					length := len(strs[0])
					newStrCharIndexNew := newStrCharIndex + 1
					if length >= newStrCharIndexNew {
						//if strs[0] != nextStr {
						if nestedIndex != 0 {
							fmt.Println("strs0: ", strs[0])
							fmt.Println("nextStr: ", nextStr)

							fmt.Println("strs0[:newStrCharIndexNew] ", strs[0][0:newStrCharIndexNew])
							fmt.Println("nextStr[:newStrCharIndexNew] ", nextStr[0:newStrCharIndexNew])
							firstSubStr := string(strs[0])[0:newStrCharIndexNew]
							secondSubStr := string(nextStr)[0:newStrCharIndexNew]

							_ = firstSubStr
							_ = secondSubStr

							if strs[0][:newStrCharIndexNew] == nextStr[:newStrCharIndexNew] {
								if nextStr[:newStrCharIndexNew] != "" {
									tempPrefix = nextStr[:newStrCharIndexNew]
								}
							}
						}

					}
				}
			}

			if !slices.Contains(longestPrefix, tempPrefix) {
				if len(tempPrefix) > 0 && tempPrefix != " " && coreIndex != 0 {
					longestPrefix = append(longestPrefix, tempPrefix)
				}
			}
		}
	}
	fmt.Println(longestPrefix)
	fmt.Println(len(longestPrefix))

	if len(longestPrefix) > 0 {
		return strings.Trim(longestPrefix[0], "")
	} else {
		return ""
	}
}
