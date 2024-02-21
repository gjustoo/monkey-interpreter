package main

import (
	"fmt"
	"interpreter/src/monkey/repl"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	fmt.Println(MONKEY_FACE)

	repl.Start(os.Stdin, os.Stdout)
}

const MONKEY_FACE = `
       .="=.
     _/.-.-.\_     _
    ( ( o o ) )    ))
     |/  "  \|    //
      \'---'/    //
      /"""""\\  ((
     / /_,_\ \\  \\
     \_\\_'__/ \  ))
     /\  / ~\  |//
    /   /    \  /
,--\,--'\/\    /
  
__   __  _______  __    _  ___   _  _______  __   __  ___      _______  __    _  _______ 
|  |_|  ||       ||  |  | ||   | | ||       ||  | |  ||   |    |   _   ||  |  | ||       |
|       ||   _   ||   |_| ||   |_| ||    ___||  |_|  ||   |    |  |_|  ||   |_| ||    ___|
|       ||  | |  ||       ||      _||   |___ |       ||   |    |       ||       ||   | __ 
|       ||  |_|  ||  _    ||     |_ |    ___||_     _||   |___ |       ||  _    ||   ||  |
| ||_|| ||       || | |   ||    _  ||   |___   |   |  |       ||   _   || | |   ||   |_| |
|_|   |_||_______||_|  |__||___| |_||_______|  |___|  |_______||__| |__||_|  |__||_______|



`
