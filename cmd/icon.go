package cmd

import (
	"fmt"
	"imagemage/pkg/filehandler"
	"imagemage/pkg/gemini"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	iconSizes  string
	iconType   string
	iconOutput string
)

var iconCmd = &cobra.Command{
	Use:   "icon [description]",
	Short: "Generate app icons, favicons, and UI elements",
	Long: `Generate icons in multiple sizes for apps, websites, and UI elements.

Examples:
  imagemage icon "coffee cup logo"
  imagemage icon "rocket ship" --sizes="64,128,256" --type="app-icon"
  imagemage icon "hamburger menu" --type="ui-element"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runIcon,
}

func init() {
	rootCmd.AddCommand(iconCmd)

	iconCmd.Flags().StringVar(&iconSizes, "sizes", "64,128,256", "Comma-separated list of icon sizes")
	iconCmd.Flags().StringVar(&iconType, "type", "app-icon", "Icon type: app-icon, favicon, ui-element")
	iconCmd.Flags().StringVarP(&iconOutput, "output", "o", ".", "Output directory for icons")
}

func runIcon(cmd *cobra.Command, args []string) error {
	description := args[0]

	// Parse sizes
	sizeStrs := strings.Split(iconSizes, ",")
	sizes := make([]int, 0, len(sizeStrs))
	for _, s := range sizeStrs {
		size, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return fmt.Errorf("invalid size: %s", s)
		}
		sizes = append(sizes, size)
	}

	// Create enhanced prompt for icon generation
	prompt := fmt.Sprintf("Create a clean, professional %s icon: %s. The icon should be simple, recognizable, and work well at small sizes. Center the icon on a transparent or solid background.", iconType, description)

	// Create Gemini client
	client, err := gemini.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	fmt.Printf("Generating icon: %s\n", description)
	fmt.Printf("Type: %s\n", iconType)
	fmt.Printf("Sizes: %v\n", sizes)
	fmt.Println()

	// For now, generate one base icon
	// In a production version, you might want to generate optimized versions for each size
	fmt.Println("Generating base icon...")

	imageData, err := client.GenerateContent(prompt)
	if err != nil {
		return fmt.Errorf("failed to generate icon: %w", err)
	}

	// Save icons at different "sizes" (note: we're saving the same image with size indicators in filename)
	// In a real implementation, you might resize or regenerate for each size
	successCount := 0
	for _, size := range sizes {
		filename := filehandler.GenerateFilename(description, fmt.Sprintf("icon_%dx%d", size, size), 0)
		outputPath := filepath.Join(iconOutput, filename)
		outputPath = filehandler.EnsureUniqueFilename(outputPath)

		if err := filehandler.SaveImage(imageData, outputPath); err != nil {
			fmt.Printf("Error saving %dx%d icon: %v\n", size, size, err)
			continue
		}

		fmt.Printf("✓ Saved %dx%d icon to: %s\n", size, size, outputPath)
		successCount++
	}

	fmt.Printf("\nSuccessfully generated %d/%d icon sizes\n", successCount, len(sizes))
	fmt.Println("\nNote: The same base image was saved with different filenames.")
	fmt.Println("For production use, consider resizing these images to their target dimensions.")

	return nil
}
