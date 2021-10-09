package main

import "testing"

func TestRomanToInt(t *testing.T) {
    if romanToInt("III") != 3 {
        t.Error(`romanToInt("III") != 3`)
    }

    if romanToInt("IX") != 9 {
        t.Error(`romanToInt("IX") != 9`)
    }

    if romanToInt("XCIX") != 99 {
        t.Error(`romanToInt("XCIX") != 99`)
    }
}
