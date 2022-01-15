package main

import "fmt"

func main() {
	// dis is array in golang
	fourNum := [4]int{0, 1, 2, 3}
	// array is "immutable" ** fix ขนาดจ้า
	fmt.Println("fourNum: ", fourNum)
	// ในโลกของ go ใช้ slice มากกว่า

	nums := make([]int, 4) // วิธีประกาศ slice
	// nums := make([]int, 4, 6)
	fmt.Println("nums: ", nums) // ผลลัพธ์ที่ได้เป็น slice ที่ประกอบ ด้วย zero(initial) value ของ int คือ 0
	// โครงสร้างของ slice
	// []int => [ptr][len][cap]
	// // ptr คือ pointer ที่ชี้ไปยัง array สักแห่งใน ram
	// // len คือ สมาชิก array ตัวปััจจุบัน
	// // cap คือ ขนาดความจุของ array ตัวปััจจุบัน
	// // // ถ้า len เท่ากับ cap แล้้ว go จะ hard copy array แล้วทำการเพิ่ม cap

	// // อยากเพิ่มของใน slice ใช้ append
	nums = append(nums, 4)
	fmt.Println("nums after append 4: ", nums)

	sliceNum := nums[1:3]
	fmt.Println("sliceNum : ", sliceNum)

	// // วิธี cast array 2 slice
	zeroArray := [5]int{}
	zeroSlice := zeroArray[:] // <--- convert here !!!!!
	zeroSlice = append(zeroSlice, 1)
	fmt.Println("zeroSlice: ", zeroSlice)

	abc()
	efg()
	cde()
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
	s = s[:3]
	fmt.Println(s)
	// [apple banana coconut]
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	s = s[4:]

	fmt.Println(s)
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed
	s = s[2:5]

	fmt.Println(s)
	// * [coconut durian elderberries]
}
