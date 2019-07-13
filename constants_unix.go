// +build !windows

package cli

// Add default constants for non-windows environments
const (
	defaultPrompt             = "$"
	defaultEnvSetCmd          = "export"
	defaultAssignmentOperator = "="
)
