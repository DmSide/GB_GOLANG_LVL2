package task8

import (
	"testing"
)

func TestGetUniqueAndDoubles(t *testing.T) {
	t.Run("GetUniqueAndDoubles", func(t *testing.T) {
		done := make(chan struct{})
		close(done)
		chOut := make(chan MyFileInfo)
		all := []MyFileInfo{
			{
				FilePath: "mock",
				name:     "1",
				size:     0,
			},
			{
				FilePath: "mock",
				name:     "2",
				size:     0,
			},
		}
		for _, val := range all {
			chOut <- val
		}

		uniques, doubles := GetUniqueAndDoubles(done, chOut)
		if doubles != nil && len(uniques) == 1 {
			t.Error("___")
		}
	})
}
