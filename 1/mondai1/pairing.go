package pairing

import (
	"math/rand"
)

func MakePair(rnd *rand.Rand, strs []string) [][]string {
  perm := rand.Perm(len(strs))
  pairs  := make([][]string, 0)
  for i := 0; i < len(strs); i+=2 {
    pairs = append(pairs, []string{strs[perm[i]], strs[perm[i+1]]})
  }
  return pairs
}
