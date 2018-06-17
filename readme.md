# Build a realtime photo stream using Go and Pusher Channels
This is a demo on how to build a realtime photo feed using Go and Pusher.

[View tutorial](https://pusher.com/tutorials/live-comments-go-vuejs)

### Prerequisites
- An IDE of your choice e.g. [Visual Studio Code](https://code.visualstudio.com/).
- [Go](https://golang.org/doc/install) installed on your computer.
- Basic knowledge of GoLang.
- SQLite (v3.x) [installed on your machine](http://www.sqlitetutorial.net/download-install-sqlite/).
- Basic knowledge of JavaScript (ES6 syntax), Vuejs and jQuery.
- Basic knowledge of using a CLI tool or terminal.
- Pusher application. Create one [here](http://pusher.com).

## Getting Started
To get started with the project, make sure you have all the prequiisites above.

1. Clone the project to your machine.
2. Update the Pusher keys in the `models/models.go` and `public/index.html` file.
3. Run the command: `$ go run main.go`.
4. Visit http://localhost:9000 to see application in action.


## Built With
* [Go](https://golang.org/doc/install) - Modern programming language.
* [Pusher](https://pusher.com) - build realtime applications easily.
* [Vue.js](http://vuejs.org) - JavaScript framework
* [Echo](https://echo.labstack.com/) - Go web framework.
