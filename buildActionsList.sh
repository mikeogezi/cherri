#!/bin/sh

# Define the input and output files
input_file="main.go"
output_file="main_without_main.go"

# Remove the main function and create a temp file
sed '/func main()/,/^}/d' "$input_file" > "$output_file"

# Run go with all .go files in the current directory, excluding main.go and *_test.go files
go run $(find . -maxdepth 1 -name "*.go" ! -name "main.go" ! -name "*_test.go")

# Remove the temporary file
rm "$output_file"
