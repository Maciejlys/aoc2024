package main

import (
	_ "embed"
	"slices"
	"strings"
)

type diskSpace struct {
	ID   int
	Size int
}

type diskSpaces []diskSpace

func parse2(input string) diskSpaces {
	diskmap := strings.TrimSpace(input) + "0"

	files := make([]diskSpace, 0)
	for id := 0; id*2 < len(diskmap); id++ {
		size, free := int(diskmap[id*2]-'0'), int(diskmap[id*2+1]-'0')
		files = append(files, diskSpace{id, size}, diskSpace{-1, free})
	}
	return files
}

func (this diskSpaces) getChecksum() int {
	checksum := 0
	i := 0
	for _, f := range this {
		for range f.Size {
			if f.ID != -1 {
				checksum += i * f.ID
			}
			i++
		}
	}
	return checksum
}

func (this *diskSpaces) moveFiles() {
	for file := len((*this)) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			if (*this)[file].ID != -1 && (*this)[free].ID == -1 && (*this)[free].Size >= (*this)[file].Size {
				(*this) = slices.Insert((*this), free, (*this)[file])
				(*this)[file+1].ID = -1
				(*this)[free+1].Size = (*this)[free+1].Size - (*this)[file+1].Size
			}
		}
	}
}

func part2(input string) int {
	files := parse2(input)
	files.moveFiles()

	return files.getChecksum()
}
