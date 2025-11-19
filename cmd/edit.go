package cmd

import (
	"fmt"
	"nanobanana/pkg/filehandler"
	"nanobanana/pkg/gemini"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	editOutput string
)

var editCmd = &cobra.Command{
	Use:   "edit [image-path] [instruction]",
	Short: "Edit an existing image with natural language instructions",
	Long: `Modify an existing image using natural language instructions.

Examples:
  nanobanana edit photo.png "make it black and white"
  nanobanana edit landscape.png "add a rainbow in the sky"
  nanobanana edit portrait.png "change background to beach" --output=./edited`,
	Args: cobra.ExactArgs(2),
	RunE: runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&editOutput, "output", "o", "", "Output path for edited image (default: input_edited.png)")
}

func runEdit(cmd *cobra.Command, args []string) error {
	imagePath := args[0]
	instruction := args[1]

	fmt.Printf("Loading image: %s\n", imagePath)

	// Load image as base64
	imageBase64, err := filehandler.LoadImageAsBase64(imagePath)
	if err != nil {
		return fmt.Errorf("failed to load image: %w", err)
	}

	// Create Gemini client
	client, err := gemini.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	fmt.Printf("Editing image with instruction: %s\n", instruction)

	// Generate edited image
	editedImageData, err := client.GenerateContentWithImage(instruction, imageBase64)
	if err != nil {
		return fmt.Errorf("failed to edit image: %w", err)
	}

	// Determine output path
	outputPath := editOutput
	if outputPath == "" {
		// Generate default output path
		ext := filepath.Ext(imagePath)
		base := strings.TrimSuffix(imagePath, ext)
		outputPath = base + "_edited" + ext
	}

	outputPath = filehandler.EnsureUniqueFilename(outputPath)

	// Save edited image
	if err := filehandler.SaveImage(editedImageData, outputPath); err != nil {
		return fmt.Errorf("failed to save edited image: %w", err)
	}

	fmt.Printf("✓ Edited image saved to: %s\n", outputPath)

	return nil
}
