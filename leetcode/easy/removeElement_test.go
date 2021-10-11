package main

import "testing"

func TestRemoveElement(t *testing.T) {
    if removeElement([]int{3, 2, 2, 3}, 3) != 2 {
        t.Error(`removeElement([]int{3, 2, 2, 3}) != 2`)
    }

    if removeElement([]int{0,1,2,2,3,0,4,2}, 2) != 5 {
        t.Error(`removeElement([]int{0,1,2,2,3,0,4,2}) != 5`)
    }
}
