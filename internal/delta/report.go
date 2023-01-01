package delta

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
)

const (
	inserted = "inserted"
	copied   = "copied"
	deleted  = "deleted"

	insertedIcon = "+"
	copiedIcon   = ">"
	deletedIcon  = "-"
)

type Report struct {
	chunk     *SingleDelta
	operation string
}

var colors = map[string]func(format string, a ...interface{}){
	inserted: color.Green,
	copied:   color.Blue,
	deleted:  color.Red,
}

func GenerateDeltaReport(delta *Delta) {
	report := []*Report{}

	for _, x := range delta.Inserted {
		report = append(report, &Report{
			chunk:     x,
			operation: inserted,
		})
	}
	for _, x := range delta.Copied {
		report = append(report, &Report{
			chunk:     x,
			operation: copied,
		})
	}
	for _, x := range delta.Deleted {
		report = append(report, &Report{
			chunk:     x,
			operation: deleted,
		})
	}

	sort.Slice(report, func(i, j int) bool {
		return report[i].chunk.Start < report[j].chunk.Start
	})

	var spaces []string
	for _, item := range report {
		fillColoredLine := colors[item.operation]
		chunks := fmt.Sprintf(" [ %v - %v ] ", item.chunk.Start, item.chunk.End)
		if item.operation == inserted {
			fillColoredLine(fmt.Sprintf("%s %v %v", insertedIcon, chunks, string(item.chunk.DiffBytes)))
		}

		if item.operation == deleted {
			if len(item.chunk.DiffBytes) != 0 {
				fillColoredLine(fmt.Sprintf("%v %v %v", deletedIcon, chunks, string(item.chunk.DiffBytes)))
			} else {
				for i := 0; i < item.chunk.End-item.chunk.Start; i++ {
					spaces = append(spaces, "-")
				}
				fillColoredLine(fmt.Sprintf("%v %v %v", deletedIcon, chunks, strings.Join(spaces, "")))
			}
		}

		if item.operation == copied {
			if len(item.chunk.DiffBytes) != 0 {
				fillColoredLine(fmt.Sprintf("%v %v %v", copiedIcon, chunks, string(item.chunk.DiffBytes)))
			} else {
				for i := 0; i < item.chunk.End-item.chunk.Start; i++ {
					spaces = append(spaces, "-")
				}
				fillColoredLine(fmt.Sprintf("%v %v %v", copiedIcon, chunks, strings.Join(spaces, "")))
			}
		}
		spaces = nil
	}
}
