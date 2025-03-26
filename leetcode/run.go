package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) <= 1 {
		fmt.Println("Need test number/name")
		return
	}
	name := os.Args[1]
	if name == "" {
		name = "547"
	}

	tests := map[string]func(){
		"1071":        test_gcdOfStrings,
		"reverseList": test_reverseList,
		"104":         test_maxDepth,
		"283":         test_moveZeroes,
		"1431":        test_kidsWithCandies,
		"605":         test_canPlaceFlowers,
		"53":          test_maxSubArray,
		"345":         test_reverseVowels,
		"151":         test_reverseWords,
		"238":         testProductExceptSelf,
		"334":         testIncreasingTriplet,
		"443":         testCompress,
		"392":         testIsSubsequence,
		"11":          testMaxArea,
		"1679":        testMaxOperations,
		"643":         testFindMaxAverage,
		"1456":        testMaxVowels,
		"1004":        testLongestOnes,
		"1493":        testLongestSubarray,
		"724":         testPivotIndex,
		"1732":        testLargestAltitude,
		"2215":        testFindDifference,
		"1207":        testUniqueOccurrences,
		"2352":        testEqualPairs,
		"2390":        testRemoveStars,
		"739":         testDailyTemperatures,
		"901":         testStockSpanner,
		"735":         testAsteroidCollision,
		"394":         testDecodeString,
		"216":         testCombinationSum3,
		"17":          testLetterCombinations,
		"1137":        testTribonacci,
		"746":         testMinCostClimbingStairs,
		"198":         testRob,
		"872":         test_leafSimilar,
		"1448":        test_goodNodes,
		"437":         test_pathSum,
		"1372":        test_longestZigZag,
		"236":         test_lowestCommonAncestor,
		"199":         test_rightSideView,
		"1161":        test_maxLevelSum,
		"700":         test_searchBST,
		"450":         test_deleteNode,
		"993":         test_recentCalls,
		"649":         test_predictPartyVictory,
		"2095":        test_deleteMiddle,
		"328":         test_oddEvenList,
		"2130":        test_pairSum,
		"841":         test_canVisitAllRooms,
		"547":         test_findCircleNum,
	}
	tests[name]()
	// if name == "mergeAlternately" {
	// 	fmt.Println(mergeAlternately("Hello", "world"))
	// } else if name == "1071" {
	// 	test_gcdOfStrings()
	// } else if name == "reverseList" {
	// 	test_reverseList()
	// } else if name == "104" {
	// 	test_maxDepth()
	// } else if name == "283" {
	// 	test_moveZeroes()
	// } else if name == "1431" {
	// 	test_kidsWithCandies()
	// } else if name == "605" {
	// 	test_canPlaceFlowers()
	// } else if name == "53" {
	// 	test_maxSubArray()
	// } else if name == "345" {
	// 	test_reverseVowels()
	// } else if name == "151" {
	// 	test_reverseWords()
	// } else if name == "238" {
	// 	testProductExceptSelf()
	// } else if name == "334" {
	// 	testIncreasingTriplet()
	// } else if name == "443" {
	// 	testCompress()
	// } else if name == "392" {
	// 	testIsSubsequence()
	// } else if name == "11" {
	// 	testMaxArea()
	// } else if name == "1679" {
	// 	testMaxOperations()
	// } else if name == "643" {
	// 	testFindMaxAverage()
	// } else if name == "1456" {
	// 	testMaxVowels()
	// } else if name == "1004" {
	// 	testLongestOnes()
	// } else if name == "1493" {
	// 	testLongestSubarray()
	// } else if name == "724" {
	// 	testPivotIndex()
	// } else if name == "1732" {
	// 	testLargestAltitude()
	// } else if name == "2215" {
	// 	testFindDifference()
	// } else if name == "1207" {
	// 	testUniqueOccurrences()
	// } else if name == "2352" {
	// 	testEqualPairs()
	// } else if name == "2390" {
	// 	testRemoveStars()
	// } else if name == "739" {
	// 	testDailyTemperatures()
	// } else if name == "901" {
	// 	testStockSpanner()
	// } else if name == "735" {
	// 	testAsteroidCollision()
	// } else if name == "394" {
	// 	testDecodeString()
	// } else if name == "216" {
	// 	testCombinationSum3()
	// } else if name == "17" {
	// 	testLetterCombinations()
	// } else if name == "1137" {
	// 	testTribonacci()
	// } else if name == "746" {
	// 	testMinCostClimbingStairs()
	// } else if name == "198" {
	// 	testRob()
	// } else if name == "872" {
	// 	test_leafSimilar()
	// } else if name == "1448" {
	// 	test_goodNodes()
	// } else if name == "437" {
	// 	test_pathSum()
	// } else if name == "1372" {
	// 	test_longestZigZag()
	// } else if name == "236" {
	// 	test_lowestCommonAncestor()
	// } else if name == "199" {
	// 	test_rightSideView()
	// } else if name == "1161" {
	// 	test_maxLevelSum()
	// } else if name == "700" {
	// 	test_searchBST()
	// } else if name == "450" {
	// 	test_deleteNode()
	// } else if name == "993" {
	// 	test_recentCalls()
	// } else if name == "649" {
	// 	test_predictPartyVictory()
	// } else if name == "2095" {
	// 	test_deleteMiddle()
	// } else if name == "328" {
	// 	test_oddEvenList()
	// } else if name == "2130" {
	// 	test_pairSum()
	// } else if name == "841" {
	// 	test_canVisitAllRooms()
	// } else if name == "547" {
	// 	test_findCircleNum()
	// }

}
