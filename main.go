package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// List of image files located in the "images" directory
// These are the paths to the images we will randomly serve.
var imageFiles = []string{
	"images/toefat.jpg",
	"images/toefat1.jpg",
	"images/toefat2.jpg",
	"images/toefat3.jpg",
	"images/toefat4.jpg",
	"images/toefat5.jpg",
	"images/toefat6.jpg",
	"images/toefat7.jpg",
	"images/toefat8.jpg",
	"images/toefat9.jpg",
	"images/toefat10.jpg",
	"images/toefat11jpg",
	"images/toefat12.jpg",
	"images/toefat13.jpg",
	"images/toefat14.jpg",
	"images/toefat15.jpg",
	"images/toefat16.jpg",
	"images/toefat17.jpg",
	"images/toefat18.jpg",
	"images/toefat19.jpg",
	"images/toefat20.jpg",
	"images/toefat21.jpg",
	"images/toefat22.jpg",
}

// getRandomImage selects a random image from the imageFiles slice.
// It uses the math/rand package to generate a random index based on the current Unix timestamp.
func getRandomImage() string {
	// Seed the random number generator to ensure different random selections each time
	rand.Seed(time.Now().UnixNano())

	// Return a random image path from the imageFiles slice
	return imageFiles[rand.Intn(len(imageFiles))]
}

// serveRandomImage handles HTTP requests and serves a random image to the user.
// It uses the getRandomImage function to select the image to serve.
func serveRandomImage(w http.ResponseWriter, r *http.Request) {
	// Get a random image file path from the image list
	imagePath := getRandomImage()

	// Serve the selected image file to the user
	// http.ServeFile sends the contents of the specified file as the HTTP response.
	http.ServeFile(w, r, imagePath)
}

func main() {
	// Serve static files from the "images" directory.
	// The "/images/" URL path will map to the "images/" folder on the server.
	// http.StripPrefix removes the "/images" prefix when looking for the files.
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	// Handle the root route ("/") and respond by serving a random image.
	// When a user visits the homepage, serveRandomImage will be called to return a random image.
	http.HandleFunc("/", serveRandomImage)

	// Specify the port on which to run the server.
	port := "8080"
	fmt.Printf("Server started on http://localhost:%s\n", port)

	// Start the HTTP server on the specified port.
	// If the server fails to start, we log an error and exit.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

