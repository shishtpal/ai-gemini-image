package cmd

import (
	"fmt"
	"nanobanana/pkg/filehandler"
	"nanobanana/pkg/gemini"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	generateCount       int
	generateOutput      string
	generateStyle       string
	generatePreview     bool
	generateAspectRatio string
)

var generateCmd = &cobra.Command{
	Use:   "generate [prompt]",
	Short: "Generate images from text descriptions",
	Long: `Generate one or more images from a text prompt using Gemini 2.5 Flash Image model.

Examples:
  nanobanana generate "watercolor painting of a fox in snowy forest"
  nanobanana generate "mountain landscape" --count=3 --output=./images
  nanobanana generate "cyberpunk city" --style="neon, futuristic"
  nanobanana generate "wide cinematic shot" --aspect-ratio="21:9"
  nanobanana generate "phone wallpaper" --aspect-ratio="9:16"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntVarP(&generateCount, "count", "c", 1, "Number of images to generate")
	generateCmd.Flags().StringVarP(&generateOutput, "output", "o", ".", "Output directory for generated images")
	generateCmd.Flags().StringVarP(&generateStyle, "style", "s", "", "Additional style guidance (e.g., 'watercolor', 'pixel-art')")
	generateCmd.Flags().BoolVarP(&generatePreview, "preview", "p", false, "Show preview information")
	generateCmd.Flags().StringVarP(&generateAspectRatio, "aspect-ratio", "a", "", "Aspect ratio (1:1, 16:9, 9:16, 4:3, 3:4, 3:2, 2:3, 21:9, 5:4, 4:5)")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	prompt := args[0]

	// Validate aspect ratio if provided
	if generateAspectRatio != "" {
		if err := gemini.ValidateAspectRatio(generateAspectRatio); err != nil {
			return err
		}
	}

	// Add style to prompt if specified
	fullPrompt := prompt
	if generateStyle != "" {
		fullPrompt = fmt.Sprintf("%s, style: %s", prompt, generateStyle)
	}

	// Create Gemini client
	client, err := gemini.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	fmt.Printf("Generating %d image(s) for: %s\n", generateCount, prompt)
	if generateStyle != "" {
		fmt.Printf("Style: %s\n", generateStyle)
	}
	if generateAspectRatio != "" {
		fmt.Printf("Aspect Ratio: %s\n", generateAspectRatio)
	}
	fmt.Println()

	successCount := 0
	for i := 1; i <= generateCount; i++ {
		if generateCount > 1 {
			fmt.Printf("[%d/%d] Generating image...\n", i, generateCount)
		} else {
			fmt.Println("Generating image...")
		}

		// Generate image
		imageData, err := client.GenerateContentWithOptions(fullPrompt, "", generateAspectRatio)
		if err != nil {
			fmt.Printf("Error generating image %d: %v\n", i, err)
			continue
		}

		// Generate filename
		var filename string
		if generateCount > 1 {
			filename = filehandler.GenerateFilename(prompt, "", i)
		} else {
			filename = filehandler.GenerateFilename(prompt, "", 0)
		}

		// Create output path
		outputPath := filepath.Join(generateOutput, filename)
		outputPath = filehandler.EnsureUniqueFilename(outputPath)

		// Save image
		if err := filehandler.SaveImage(imageData, outputPath); err != nil {
			fmt.Printf("Error saving image %d: %v\n", i, err)
			continue
		}

		fmt.Printf("✓ Saved to: %s\n", outputPath)
		successCount++
	}

	fmt.Printf("\nSuccessfully generated %d/%d images\n", successCount, generateCount)

	return nil
}
