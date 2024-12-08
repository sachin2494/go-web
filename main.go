package main

import (
	"fmt"
	"net/http"
)

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Home - My Go Web Page</title>
		<link rel="stylesheet" href="/static/styles.css"> <!-- Link to the shared styles -->
	</head>
	<body>
		<!-- Semi-transparent overlay to improve content visibility -->
		<div class="overlay"></div>

		<header>
			<h1>Welcome to My Go Web Page</h1>
		</header>
		<nav>
			<a href="/">Home</a>
			<a href="/about">About</a>
			<a href="/contact">Contact</a>
			<a href="/docs">docs</a>
		</nav>
		<main>
			<h2>Welcome to my Page</h2>
			<p>This is the page which we provide Tech releted stuff</p>
			<p>Coming soon......!</p>
		</main>
		<footer>
			<p>&copy; 2024 Course Hub</p>
		</footer>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, htmlContent)
}

// About page handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>About - My Go Web Page</title>
		<link rel="stylesheet" href="/static/styles.css"> <!-- Link to the shared styles -->
	</head>
	<body>
		<!-- Semi-transparent overlay to improve content visibility -->
		<div class="overlay"></div>

		<header>
			<h1>About Us</h1>
		</header>
		<nav>
			<a href="/">Home</a>
			<a href="/about">About</a>
			<a href="/contact">Contact</a>
		</nav>
		<main>
			<h2>About Course Hub</h2>
			<p>This is the About page, featuring a static background image.</p>
		</main>
		<footer>
			<p>&copy; 2024 Course Hub</p>
		</footer>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, htmlContent)
}

// Contact page handler
func contactHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Contact - My Go Web Page</title>
		<link rel="stylesheet" href="/static/styles.css"> <!-- Link to the shared styles -->
	</head>
	<body>
		<!-- Semi-transparent overlay to improve content visibility -->
		<div class="overlay"></div>

		<header>
			<h1>Contact Us</h1>
		</header>
		<nav>
			<a href="/">Home</a>
			<a href="/about">About</a>
			<a href="/contact">Contact</a>
		</nav>
		<main>
			<h2>Contact Information</h2>
			<p>Feel free to contact us for more information.</p>
		</main>
		<footer>
			<p>&copy; 2024 Course Hub</p>
		</footer>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, htmlContent)
}

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Registering the handlers for each route
	http.HandleFunc("/", homeHandler)           // Home page route
	http.HandleFunc("/about", aboutHandler)     // About page route
	http.HandleFunc("/contact", contactHandler) // Contact page route

	// Starting the web server on port 7080
	fmt.Println("Starting server at http://localhost:7080...")
	if err := http.ListenAndServe(":7080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
