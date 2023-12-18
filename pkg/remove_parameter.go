package pkg

import (
	"prism/internal/model"
	"slices"
)

// Remove a parameter by name from a job configuration block.
func RemoveParameter(block *model.TemplateBlock, name string) {
	for index, item := range block.Parameter {
		for k := range item {
			if k == name {
				block.Parameter = slices.Delete(
					block.Parameter, index, index+1,
				)
				return
			}
		}
	}
}
