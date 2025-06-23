# ASCII Art Web Stylize

A web application that generates ASCII art from user input using different styles ("standard", "shadow", "thinkertoy"). Built with Go and HTML/CSS.

---

## Features

- Convert text to ASCII art in three styles: **standard**, **shadow**, and **thinkertoy**
- Web interface with textarea input and style selection
- Error handling for invalid input, unsupported characters, and character limits
- Responsive and modern UI

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.23.2 or higher

### Set Up

1. **Run the application:**
    ```sh
    go run .
    ```

2. **Open your browser and visit:**
    ```
    http://localhost:8000
    ```

---

## Usage

1. Enter your text in the textarea.
2. Select an ASCII art style from the dropdown.
3. Click **Generate**.
4. The ASCII art will be displayed in the output area.

---

## Error Handling

- **404 Not Found:** Accessing an invalid route.
- **400 Bad Request:** Invalid input or unsupported characters.
- **405 Method Not Allowed:** Using an unsupported HTTP method.
- **500 Internal Server Error:** Server-side error.
- **Character Limit:** Input exceeding 1000 characters will be rejected.

---

## File Overview

- **main.go**: Application entry point, HTTP server setup.
- **handlers.go**: HTTP route handlers and form processing.
- **templates.go**: Template rendering and error data structures.
- **asciiart/utils.go**: ASCII art processing logic.
- **asciiart/*.txt**: ASCII art font definitions.
- **index.html**: Main input form.
- **indexout.html**: Output page with ASCII art result.
- **error.html**: Error display page.
- **css/style.css**: Stylesheet for the web UI.

---

## Author

mbadawy
mkhattar
faisaahmed
