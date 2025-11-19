# Nanobanana

A professional CLI tool for generating and manipulating images using Google's Gemini 2.5 Flash Image model. Built with Go and Cobra, Nanobanana provides a standalone solution without requiring the underlying Gemini CLI.

## Features

- **Text-to-Image Generation** - Create images from descriptive text prompts
- **Image Editing** - Modify existing images with natural language instructions
- **Photo Restoration** - Enhance and repair old or damaged photos
- **Icon Generation** - Create app icons, favicons, and UI elements in multiple sizes
- **Pattern Creation** - Generate seamless patterns and textures
- **Visual Storytelling** - Produce sequential images for narratives
- **Technical Diagrams** - Generate flowcharts, architecture diagrams, and visualizations

## Prerequisites

- Go 1.22 or higher
- A Google Gemini API key

## Installation

### From Source

```bash
# Clone the repository
git clone <repository-url>
cd nanobanana

# Build the binary
go build -o nanobanana

# Optionally, install it to your $GOPATH/bin
go install
```

### Configuration

Set your Gemini API key as an environment variable. Nanobanana checks for API keys in this order:

```bash
export NANOBANANA_GEMINI_API_KEY="your-api-key-here"
# or
export NANOBANANA_GOOGLE_API_KEY="your-api-key-here"
# or
export GEMINI_API_KEY="your-api-key-here"
# or
export GOOGLE_API_KEY="your-api-key-here"
```

To get an API key, visit the [Google AI Studio](https://makersuite.google.com/app/apikey).

## Usage

### Generate Command

Create images from text descriptions:

```bash
# Basic generation
nanobanana generate "watercolor painting of a fox in snowy forest"

# Generate multiple variations
nanobanana generate "mountain landscape" --count=3

# Specify output directory and style
nanobanana generate "cyberpunk city" --output=./images --style="neon, futuristic"

# Generate with specific aspect ratio
nanobanana generate "wide cinematic landscape" --aspect-ratio="21:9"
nanobanana generate "phone wallpaper" --aspect-ratio="9:16"
nanobanana generate "social media post" --aspect-ratio="1:1"
```

**Flags:**
- `-c, --count` - Number of images to generate (default: 1)
- `-o, --output` - Output directory (default: current directory)
- `-s, --style` - Additional style guidance
- `-a, --aspect-ratio` - Aspect ratio (1:1, 16:9, 9:16, 4:3, 3:4, 3:2, 2:3, 21:9, 5:4, 4:5)

**Supported Aspect Ratios:**
- **Square:** 1:1 (1024x1024)
- **Landscape:** 16:9 (1344x768), 4:3, 3:2, 21:9
- **Portrait:** 9:16 (768x1344), 3:4, 2:3
- **Flexible:** 5:4, 4:5

**Note:** Use the `--aspect-ratio` flag to control the dimensions of the generated image. This is more reliable than mentioning dimensions in the prompt. If not specified, the model defaults to 1:1 (square) images.

### Edit Command

Modify existing images with natural language:

```bash
# Basic edit
nanobanana edit photo.png "make it black and white"

# Add elements
nanobanana edit landscape.png "add a rainbow in the sky"

# Change background
nanobanana edit portrait.png "change background to beach" --output=edited.png
```

**Flags:**
- `-o, --output` - Output path for edited image (default: input_edited.png)

### Restore Command

Enhance and repair photos:

```bash
# Restore an old photo
nanobanana restore old_photo.png

# Specify output path
nanobanana restore damaged.jpg --output=restored.png
```

**Flags:**
- `-o, --output` - Output path for restored image

### Icon Command

Generate icons in multiple sizes:

```bash
# Generate app icon
nanobanana icon "coffee cup logo"

# Specify sizes and type
nanobanana icon "rocket ship" --sizes="64,128,256,512" --type="app-icon"

# Generate UI element
nanobanana icon "hamburger menu" --type="ui-element"
```

**Flags:**
- `--sizes` - Comma-separated list of sizes (default: "64,128,256")
- `--type` - Icon type: app-icon, favicon, ui-element (default: "app-icon")
- `-o, --output` - Output directory

### Pattern Command

Create seamless patterns and textures:

```bash
# Generate geometric pattern
nanobanana pattern "geometric triangles"

# Specify style
nanobanana pattern "floral" --style="watercolor"

# Create minimal pattern
nanobanana pattern "hexagons" --style="minimal, modern"
```

**Flags:**
- `--type` - Pattern type: seamless, tiled, texture (default: "seamless")
- `-s, --style` - Pattern style
- `-o, --output` - Output directory

### Story Command

Generate sequential images for visual narratives:

```bash
# Create a story sequence
nanobanana story "a seed growing into a tree" --frames=4

# Specify visual style
nanobanana story "day to night transition in a city" --frames=6 --style="cinematic"

# Character transformation
nanobanana story "character transformation" --frames=3
```

**Flags:**
- `-f, --frames` - Number of frames/scenes (default: 3, min: 2, max: 10)
- `-s, --style` - Visual style for the story
- `-o, --output` - Output directory

### Diagram Command

Generate technical diagrams and flowcharts:

```bash
# Create a flowchart
nanobanana diagram "CI/CD pipeline with testing stages"

# Architecture diagram
nanobanana diagram "microservices architecture" --type="architecture"

# Sequence diagram
nanobanana diagram "user authentication flow" --type="flowchart"
```

**Flags:**
- `--type` - Diagram type: flowchart, architecture, sequence, entity-relationship (default: "diagram")
- `-o, --output` - Output directory

## Project Structure

```
nanobanana/
├── main.go                 # Application entry point
├── cmd/                    # Command implementations
│   ├── root.go            # Root command and CLI setup
│   ├── generate.go        # Text-to-image generation
│   ├── edit.go            # Image editing
│   ├── restore.go         # Photo restoration
│   ├── icon.go            # Icon generation
│   ├── pattern.go         # Pattern creation
│   ├── story.go           # Sequential image generation
│   └── diagram.go         # Diagram generation
├── pkg/
│   ├── gemini/            # Gemini API client
│   │   └── client.go
│   └── filehandler/       # File handling utilities
│       └── filehandler.go
├── go.mod                 # Go module definition
└── README.md             # This file
```

## How It Works

Nanobanana directly interacts with Google's Gemini API using the `gemini-2.5-flash-image` model:

1. **API Client**: The `pkg/gemini` package handles authentication and communication with the Gemini API
2. **Request Formation**: Text prompts are sent as JSON requests to the API endpoint
3. **Response Processing**: Images are received as base64-encoded data and decoded
4. **File Management**: The `pkg/filehandler` package handles intelligent filename generation and file I/O
5. **Commands**: Each Cobra command provides a specialized interface for different use cases

## Error Handling

The tool provides clear error messages for common issues:

- **Invalid API Key**: Ensure your API key is correctly set
- **API Quota Exceeded**: Check your API usage limits
- **Safety Concerns**: Some prompts may be rejected due to content policies
- **Network Errors**: Check your internet connection

## Development

### Building

```bash
go build -o nanobanana
```

### Testing

```bash
go test ./...
```

### Adding New Commands

1. Create a new file in `cmd/` (e.g., `cmd/newcommand.go`)
2. Implement the command using Cobra's structure
3. Register it in the `init()` function
4. Add documentation to this README

## License

[Add your license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Inspired by the [nanobanana Gemini CLI extension](https://github.com/gemini-cli-extensions/nanobanana)
- Built with [Cobra](https://github.com/spf13/cobra)
- Powered by Google's Gemini 2.5 Flash Image model
