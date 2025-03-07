# Tobese.org

A simple Go web server that serves random images of Tobi, a tobese cat with unique aesthetic qualities.

## Overview

Tobese.org is a lightweight web application built in Go that randomly selects and displays images of Tobi the cat from a collection of metadata-cleaned photographs. Each time the page is refreshed, visitors are treated to a different image from the collection.

## Features

- **Random Image Selection**: Uses Go's `math/rand` package to randomly select an image on each page load
- **Caching Prevention**: Implements HTTP headers to prevent browser caching, ensuring a new image is displayed with each visit
- **Simple Architecture**: Lightweight codebase with minimal dependencies
- **Static File Serving**: Serves static images from the local filesystem

## How It Works

1. The application maintains a list of image paths stored in the `images/` directory
2. When a user visits the root URL (`/`), the server:
   - Seeds the random number generator with the current timestamp
   - Selects a random image path from the predefined list
   - Sets cache control headers to prevent image caching
   - Serves the selected image directly to the user

## Directory Structure

```
Tobese/
│
├── main.go         # The main application code
├── images/         # Directory containing Tobi's images
│   ├── toefat.jpg
│   ├── toefat1.jpg
│   ├── toefat2.jpg
│   └── ...         # Additional images
└── README.md       # This documentation file
```

## Installation and Setup

### Prerequisites

- Go (version 1.15 or higher recommended)
- Git (for cloning the repository)

### Installation Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/toxicodryas/Tobese.git
   cd Tobese
   ```

2. Make sure your images are placed in the `images/` directory
   - All image filenames should match the patterns defined in the `imageFiles` slice
   - Images should be cleaned of metadata for privacy/security

3. Build the application:
   ```bash
   go build
   ```

4. Run the server:
   ```bash
   ./Tobese
   ```

5. The server will start on port 8080 by default. You can now access the website at http://localhost:8080

## Deployment

For deploying to a production environment:

1. Build the application for your target platform:
   ```bash
   go build
   ```

2. Copy the executable and the `images/` directory to your server

3. Consider setting up a reverse proxy (like Nginx or Caddy) to handle HTTPS and other production concerns

4. You might want to set up the application as a service using systemd or another init system

## Customization

### Adding New Images

1. Add your new images to the `images/` directory
2. Update the `imageFiles` slice in `main.go` to include the paths to your new images
3. Rebuild and restart the application

### Changing the Port

Modify the `port` variable in the `main.go` file to change the default port (8080).

## License

[Add your license information here]

## Acknowledgments

- Tobi, for being a photogenic and slightly chunky feline companion
- Go, for providing a simple and efficient way to serve random cat pictures to the internet
