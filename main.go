package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	game := chess.NewGame()
	eng, err := uci.New("stockfish")
	if err != nil {
		panic(err)
	}
	defer eng.Close()

	chooseSide(conn)

	wOrb := getPlayerSide(conn)
  isFlipped := false;
  if wOrb == "b" {
    isFlipped = true;
  }
	turn := wOrb == "w"

	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}

	for game.Outcome() == chess.NoOutcome {
		if turn {
			handlePlayerMove(conn, game)
			turn = false
		} else {
			handleAIMove(eng, game)
			turn = true
		}
		sendGameState(conn, game, isFlipped) 
	}
}

func chooseSide(conn net.Conn) {
	conn.Write([]byte("Select a side...\n"))
	conn.Write([]byte("Write 'w' for white and 'b' for black: "))
}

func getPlayerSide(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	wOrb, _ := reader.ReadString('\n')
	return strings.TrimSpace(wOrb)
}

func handlePlayerMove(conn net.Conn, game *chess.Game) {
	fmt.Println("Your turn: ")
	conn.Write([]byte("Your turn: \n"))

	reader := bufio.NewReader(conn)
	move, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("Client disconnected")
		} else {
			fmt.Println("Error reading move:", err)
		}
		return
	}

	move = strings.TrimSpace(move)
	if err := game.MoveStr(move); err != nil {
		conn.Write([]byte("Invalid move. Try again.\n"))
		fmt.Println("Invalid move:", err)
	}
}


func handleAIMove(eng *uci.Engine, game *chess.Game) {
	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second / 100}

	if err := eng.Run(cmdPos, cmdGo); err != nil {
		panic(err)
	}

	move := eng.SearchResults().BestMove
	if err := game.Move(move); err != nil {
		panic(err)
	}
}

func sendGameState(conn net.Conn, game *chess.Game, blackToMove bool) {
	boardStr := game.Position().Board().Draw()
	if blackToMove {
		boardStr = flipBoard(boardStr)
	}
	conn.Write([]byte(boardStr))
	conn.Write([]byte("\n"))
	conn.Write([]byte(game.String()))
	conn.Write([]byte("\n"))
}

func flipBoard(boardStr string) string {
	lines := strings.Split(boardStr, "\n")
	flippedLines := make([]string, len(lines))

	for i, line := range lines {
		flippedLines[len(lines)-1-i] = line
	}

	return strings.Join(flippedLines, "\n")
}
