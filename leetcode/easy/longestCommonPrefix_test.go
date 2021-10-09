package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
    strs1 := []string{"flower","flow","flight"}
    strs2 := []string{"dog","racecar","car"}

    if longestCommonPrefix(strs1) != "fl" {
        t.Error(`longestCommonPrefix({"flower","flow","flight"}) != "fl"`)
    }

    if longestCommonPrefix(strs2) != "" {
        t.Error(`longestCommonPrefix({"dog","racecar","car"}) != ""`)
    }
}
