package gemini

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ImageConfig represents configuration for image generation
type ImageGenConfig struct {
	Defaults ImageGenDefaults `json:"defaults"`
}

// ImageGenDefaults represents default settings for image generation
type ImageGenDefaults struct {
	AspectRatio       string `json:"aspectRatio"`
	Resolution        string `json:"resolution"`
	Style             string `json:"style"`
	ColorScheme       string `json:"colorScheme"`
	AdditionalContext string `json:"additionalContext"`
}

// LoadConfig loads configuration from a file
func LoadConfig(configPath string) (*ImageGenConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config ImageGenConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	return &config, nil
}

// FindConfig searches for config files in order of precedence:
// 1. Specified config path (if provided)
// 2. ./image-gen.config.json (per-talk config)
// 3. ~/src/talks/image-gen.defaults.json (global defaults)
func FindConfig(specifiedPath string) (*ImageGenConfig, error) {
	// Try specified path first
	if specifiedPath != "" {
		return LoadConfig(specifiedPath)
	}

	// Try current directory for per-talk config
	localConfig := "image-gen.config.json"
	if _, err := os.Stat(localConfig); err == nil {
		return LoadConfig(localConfig)
	}

	// Try global defaults
	homeDir, err := os.UserHomeDir()
	if err == nil {
		globalConfig := filepath.Join(homeDir, "src", "talks", "image-gen.defaults.json")
		if _, err := os.Stat(globalConfig); err == nil {
			return LoadConfig(globalConfig)
		}
	}

	// No config found - return nil (not an error)
	return nil, nil
}

// ApplyConfigToPrompt applies configuration settings to a prompt
func (c *ImageGenConfig) ApplyToPrompt(prompt string) string {
	if c == nil {
		return prompt
	}

	fullPrompt := prompt

	// Add style
	if c.Defaults.Style != "" {
		fullPrompt = fmt.Sprintf("%s, %s", fullPrompt, c.Defaults.Style)
	}

	// Add color scheme
	if c.Defaults.ColorScheme != "" {
		fullPrompt = fmt.Sprintf("%s, colors: %s", fullPrompt, c.Defaults.ColorScheme)
	}

	// Add additional context
	if c.Defaults.AdditionalContext != "" {
		fullPrompt = fmt.Sprintf("%s, %s", fullPrompt, c.Defaults.AdditionalContext)
	}

	return fullPrompt
}

// GetAspectRatio returns the aspect ratio from config, or empty string if not set
func (c *ImageGenConfig) GetAspectRatio() string {
	if c == nil {
		return ""
	}
	return c.Defaults.AspectRatio
}

// GetResolution returns the resolution from config, or empty string if not set
func (c *ImageGenConfig) GetResolution() string {
	if c == nil {
		return ""
	}
	return c.Defaults.Resolution
}
