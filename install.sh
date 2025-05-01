#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if curl is installed
if ! command -v curl &> /dev/null; then
    echo -e "${RED}Error: curl is not installed. Please install curl first.${NC}"
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed. Please install Go first.${NC}"
    exit 1
fi

# Create temporary directory
TEMP_DIR=$(mktemp -d)
cd $TEMP_DIR

echo -e "${YELLOW}Downloading podlog...${NC}"

# Clone the repository
git clone https://github.com/Grim-R3ap3r/podlog.git
cd podlog

# Build the project
echo -e "${YELLOW}Building podlog...${NC}"
go build -o podlog

# Install to /usr/local/bin
echo -e "${YELLOW}Installing podlog to /usr/local/bin...${NC}"
sudo cp podlog /usr/local/bin/

# Clean up
cd ..
rm -rf $TEMP_DIR

echo -e "${GREEN}podlog has been successfully installed!${NC}"
echo -e "${YELLOW}You can now use the 'podlog' command.${NC}"
echo -e "${YELLOW}Example: podlog -n your-namespace -p your-pod${NC}" 