package main

import "fmt"

func removeDuplicates(nums []int) int {
    for i:=0; i< len(nums) - 1; i++{
        if nums[i] == nums[i+1]{
            if len(nums) > 2{
                nums = append(nums[:i], nums[i+1:]...)
            } else {
                nums = []int{nums[i]}
            }
            i = -1
        }
    }
    return len(nums)
}
func main(){
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(removeDuplicates(nums))

}