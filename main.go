package main

import (
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
)

func main() {
	// nvim = neovim.attach('socket', os.environ['NVIM_LISTEN_ADDRESS'])
	// nvim.vars['__autocd_cwd'] = os.getcwd()
	// nvim.command('execute "lcd" fnameescape(g:__autocd_cwd)')
	// del nvim.vars['__autocd_cwd']

	n, err := nvim.Dial(os.Getenv("NVIM_LISTEN_ADDRESS"))
	if err != nil {
		log.Fatalf("could not connect to neovim: %s", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %s", err)
	}
	err = n.SetVar("__autocd_cwd", cwd)
	if err != nil {
		log.Fatalf("could not set variable: %s", err)
	}
	err = n.Command("execute \"lcd\" fnameescape(g:__autocd_cwd)")
	if err != nil {
		log.Fatalf("could not execute command: %s", err)
	}
	err = n.DeleteVar("__autocd_cwd")
	if err != nil {
		log.Fatalf("could not delete variable: %s", err)
	}
	err = n.Close()
	if err != nil {
		log.Fatalf("could not close connection to neovim: %s", err)
	}
}
