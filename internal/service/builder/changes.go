package builder

import (
	"prism/internal/model"
)

type Changes struct{}

func NewChanges() *Changes {
	return &Changes{}
}

// Making changes to the configuration file.
// Adds parameters and blocks specified in additional files,
// parameters from flags and environment variables.
func (s *Changes) SetChanges(
	config *model.TemplateBlock,
	changes *model.Changes,
) error {
	return nil
}
