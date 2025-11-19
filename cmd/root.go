package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nanobanana",
	Short: "A CLI tool for generating and manipulating images using Google's Gemini 2.5 Flash Image model",
	Long: `Nanobanana is a professional CLI tool for creative image generation using Google's Gemini API.
It provides commands for text-to-image creation, image editing, photo restoration, icon generation,
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
	// Add global flags here if needed
}
