package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// SetVersionInfo sets the version info from main package
func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	date = d
	rootCmd.Version = version
}

var rootCmd = &cobra.Command{
	Use:     "imagemage",
	Version: version,
	Short:   "A CLI tool for generating and manipulating images using Google's Gemini image models",
	Long: `Imagemage is a focused CLI tool for image generation using Google's Gemini API.

Supports multiple Gemini models:
  • Gemini 3 Pro Image (default) - High-quality 4K generation
  • Gemini 2.5 Flash Image (--frugal) - Faster, cheaper generation

Features include text-to-image creation, image editing, photo restoration, icon generation,
pattern creation, visual narratives, and technical diagrams.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Cobra automatically adds --version flag when Version is set
}
