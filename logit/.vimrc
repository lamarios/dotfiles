set number relativenumber
set showcmd
set incsearch
set hlsearch
syntax on

" Automatically exit insert mode
au InsertEnter * let updaterestore=&updatetime | set updatetime=4000
au InsertLeave * let &updatetime=updaterestore
au CursorHoldI * stopinsert
