# goneocd

use within neovim terminals to change the working directory

add goneocd somewhere on your PATH and the following to .zshrc
```sh
# change dir inside neovim
neovim_autocd() {
    [[ $NVIM_LISTEN_ADDRESS ]] && goneocd
}
chpwd_functions+=( neovim_autocd )
```
