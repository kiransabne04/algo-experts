package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Block map[string]bool

func main() {
	blocks := make([]Block, 0)
	inputs := `[
    {
      "gym": true,
      "pool": false,
      "school": true,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": true,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": true
    },
    {
      "gym": true,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": false,
      "store": false
    },
    {
      "gym": false,
      "pool": false,
      "school": true,
      "store": false
    },
    {
      "gym": false,
      "pool": true,
      "school": false,
      "store": false
    }
  ]`

	err := json.Unmarshal([]byte(inputs), &blocks)
	if err != nil {
		fmt.Printf("parsing block failed \n", err)
	}
	fmt.Println(blocks)
	reqs := []string{"gym", "pool", "school", "store"}

	val := ApartmentHunting(blocks, reqs)
	fmt.Println(val)

}

func ApartmentHunting(blocks []Block, reqs []string) int {
	// Write your code here.
	fmt.Println(reqs)
	reqBlkDistance := [][]int{}
	for i, v := range reqs {
		fmt.Println(i, v)
		// for each requirement, we will caclucate distance

		blockDist := make([]int, len(blocks))
		countAway := math.MaxInt64
		fmt.Printf("countAway %v form position %d\n", v, i)
		// left to right pass
		countAway = math.MaxInt64
		for blockIdx := 0; blockIdx < len(blocks); blockIdx++ {
			blockVal, found := blocks[blockIdx][v]
			if blockVal && found {
				countAway = 0

			} else if countAway != math.MaxInt64 {
				countAway++
			}

			// update blockDist, to countAway val
			blockDist[blockIdx] = countAway
		}
		fmt.Println("blockDist 0> ", blockDist)
		// reverse pass
		countAway = math.MaxInt64
		for blockIdx := len(blocks) - 1; blockIdx >= 0; blockIdx-- {
			blockVal, found := blocks[blockIdx][v]
			if blockVal && found {
				countAway = 0
			} else if countAway != math.MaxInt64 {
				countAway++
			}

			// update blockDist, the elements should be max of both value in index in both above loop
			// blockDist[blockIdx] = countAway
			fmt.Println("reversePass i, v", blockIdx, countAway)
			blockDist[blockIdx] = min(blockDist[blockIdx], countAway)

		}

		fmt.Println("blockDist 0> ", blockDist)
		reqBlkDistance = append(reqBlkDistance, blockDist)
		fmt.Println("reqBlkDistance -> ", reqBlkDistance)
	}

	//Find the block with the smallest maximum distance across all requirements
	minDistance := math.MaxInt64
	bestBlock := -1

	for i := 0; i < len(blocks); i++ {
		maxDistanceAtBlock := 0
		for _, distances := range reqBlkDistance {
			maxDistanceAtBlock = max(maxDistanceAtBlock, distances[i])
		}

		if maxDistanceAtBlock < minDistance {
			minDistance = maxDistanceAtBlock
			bestBlock = i
		}
	}

	return bestBlock
	//return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
