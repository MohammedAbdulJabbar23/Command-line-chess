# Commandline-chess

Chess Server with Stockfish

You can try it online by:

    telnet 16.171.159.45 8080

A simple TCP-based chess server written in Go that integrates with the Stockfish chess engine. This server allows players to connect, choose their side, make moves, and play against Stockfish.
Features

    Play chess against the Stockfish engine
    Choose to play as white or black
    Real-time move handling
    ANSI color-coded board representation in terminal
    

Requirements

    Go (version 1.18 or later)
    Stockfish chess engine
    

Installation

Clone the Repository:

bash

    git clone https://github.com/yourusername/chess-server-go.git
    
    
    cd chess-server-go

Install Dependencies:

Ensure you have Go installed and use the following command to fetch dependencies:

bash

    go mod tidy

    Download Stockfish:
        Download the Stockfish binary from the official Stockfish website.
        Place the Stockfish binary in the same directory as your Go code or adjust the path in the code to where you have it.

Usage

  Run the Server:

  Start the chess server with:
  bash

    go run main.go
    

The server will start listening on port 8080.

Connect to the Server:

You can use a tool like telnet to connect to the server:

bash

    telnet localhost 8080

    Play Chess:
        When prompted, choose your side by typing w for white or b for black.
        Make your moves using standard algebraic notation (e.g., e2e4).
        The board will be displayed with ANSI color codes, making it easier to visualize the game.

Code Overview

    main.go: The main file containing the server logic, including connection handling, game management, and interaction with Stockfish.
    boardWithColor function: Formats and colors the board for terminal output.
    handlePlayerMove and handleAIMove functions: Handle moves from the player and the Stockfish engine, respectively.

Contributing

Contributions are welcome! If you have suggestions or improvements, please open an issue or submit a pull request.
License

This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

    Stockfish Chess Engine
    Notnil Chess Library
