#!/bin/bash
# Example usage of nanobanana CLI
# Make sure to set your API key first:
# export GEMINI_API_KEY="your-api-key"

# Build the project if needed
if [ ! -f "./nanobanana" ]; then
    echo "Building nanobanana..."
    go build -o nanobanana
fi

# Create output directory
mkdir -p ./examples

echo "=== Example 1: Generate a simple image ==="
./nanobanana generate "watercolor painting of a fox in snowy forest" --output=./examples

echo ""
echo "=== Example 2: Generate multiple variations ==="
./nanobanana generate "mountain landscape at sunset" --count=3 --output=./examples

echo ""
echo "=== Example 3: Generate with style ==="
./nanobanana generate "cyberpunk city" --style="neon, futuristic, rainy" --output=./examples

echo ""
echo "=== Example 3b: Generate with aspect ratio ==="
./nanobanana generate "ultra wide cinematic landscape" --aspect-ratio="21:9" --output=./examples
./nanobanana generate "phone wallpaper nature scene" --aspect-ratio="9:16" --output=./examples

echo ""
echo "=== Example 4: Create an icon ==="
./nanobanana icon "coffee cup logo" --sizes="64,128,256" --output=./examples

echo ""
echo "=== Example 5: Generate a pattern ==="
./nanobanana pattern "geometric triangles" --style="minimal, modern" --output=./examples

echo ""
echo "=== Example 6: Create a diagram ==="
./nanobanana diagram "microservices architecture with API gateway" --type="architecture" --output=./examples

echo ""
echo "=== Example 7: Generate a story sequence ==="
./nanobanana story "sunrise to sunset over a lake" --frames=4 --style="peaceful, naturalistic" --output=./examples

echo ""
echo "All examples completed! Check the ./examples directory for generated images."
