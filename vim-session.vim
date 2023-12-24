let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd /mnt/c/Users/Lenovo/Documents/Programming/util
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 /mnt/c/Users/Lenovo/Documents/Programming/util
badd +5 io/argparse/Parser.go
badd +221 io/argparse/Parser_test.go
badd +2 io/argparse/Errors.go
badd +3 io/argparse/Arg.go
badd +1 io/argparse/Common.go
badd +22 io/argparse/String.go
badd +39 test/Test.go
badd +16 io/argparse/argTypes/IpV4Addr.go
badd +14 io/argparse/argTypes/String.go
badd +17 io/argparse/argTypes/Flag.go
badd +6 io/argparse/argTypes/Errors.go
badd +6 dataStruct/Tree.go
badd +5 io/argparse/argTypes/Common.go
badd +0 io/argparse/Compiler_test.go
badd +1 io/argparse/Compiler.go
badd +29 algo/parser/Parser.go
badd +53 algo/iter/Producer.go
badd +39 dataStruct/types/Types.go
badd +12 dataStruct/Common.go
badd +108 algo/iter/Producer_test.go
argglobal
%argdel
$argadd /mnt/c/Users/Lenovo/Documents/Programming/util
set stal=2
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabnew +setlocal\ bufhidden=wipe
tabrewind
edit io/argparse/Common.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd _ | wincmd |
split
wincmd _ | wincmd |
split
2wincmd k
wincmd w
wincmd w
wincmd w
wincmd _ | wincmd |
split
1wincmd k
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe '1resize ' . ((&lines * 15 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 15 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe '3resize ' . ((&lines * 14 + 24) / 49)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
exe '4resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 4resize ' . ((&columns * 105 + 105) / 211)
exe '5resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 5resize ' . ((&columns * 105 + 105) / 211)
argglobal
balt io/argparse/Errors.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 19 - ((2 * winheight(0) + 7) / 15)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 19
normal! 0
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Errors.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Errors.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Errors.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Errors.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Common.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 7) / 15)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 012|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler_test.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler_test.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler_test.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler_test.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser_test.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 1 - ((0 * winheight(0) + 7) / 14)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 1
normal! 02|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 5 - ((4 * winheight(0) + 11) / 23)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 5
normal! 0
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Compiler.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 94 - ((15 * winheight(0) + 11) / 22)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 94
normal! 0
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
exe '1resize ' . ((&lines * 15 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 15 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe '3resize ' . ((&lines * 14 + 24) / 49)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
exe '4resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 4resize ' . ((&columns * 105 + 105) / 211)
exe '5resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 5resize ' . ((&columns * 105 + 105) / 211)
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/algo/parser/Parser.go
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 45 - ((41 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 45
normal! 09|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/Common.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
1wincmd h
wincmd _ | wincmd |
split
1wincmd k
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe '1resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 22 - ((21 * winheight(0) + 11) / 23)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 22
normal! 0
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/Common.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 39 - ((18 * winheight(0) + 11) / 22)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 39
normal! 012|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/types/Types.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 53 - ((6 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 53
normal! 0
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
exe '1resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 105 + 105) / 211)
exe '2resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 105 + 105) / 211)
exe 'vert 3resize ' . ((&columns * 105 + 105) / 211)
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer_test.go
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/algo/iter/Producer.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 108 - ((38 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 108
normal! 016|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/dataStruct/Tree.go
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 5 - ((4 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 5
normal! 05|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Common.go
let s:save_splitbelow = &splitbelow
let s:save_splitright = &splitright
set splitbelow splitright
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
wincmd _ | wincmd |
vsplit
3wincmd h
wincmd _ | wincmd |
split
1wincmd k
wincmd w
wincmd w
wincmd w
wincmd w
let &splitbelow = s:save_splitbelow
let &splitright = s:save_splitright
wincmd t
let s:save_winminheight = &winminheight
let s:save_winminwidth = &winminwidth
set winminheight=0
set winheight=1
set winminwidth=0
set winwidth=1
exe '1resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 52 + 105) / 211)
exe '2resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 3resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 4resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 5resize ' . ((&columns * 52 + 105) / 211)
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 5 - ((4 * winheight(0) + 11) / 23)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 5
normal! 016|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Common.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 6 - ((5 * winheight(0) + 11) / 22)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 6
normal! 070|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Flag.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Flag.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Flag.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Flag.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Errors.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 17 - ((16 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 17
normal! 011|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/String.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/String.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/String.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/String.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/Flag.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 14 - ((13 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 14
normal! 012|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
argglobal
if bufexists(fnamemodify("/mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/IpV4Addr.go", ":p")) | buffer /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/IpV4Addr.go | else | edit /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/IpV4Addr.go | endif
if &buftype ==# 'terminal'
  silent file /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/IpV4Addr.go
endif
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/argTypes/String.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 16 - ((15 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 16
normal! 012|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
wincmd w
exe '1resize ' . ((&lines * 23 + 24) / 49)
exe 'vert 1resize ' . ((&columns * 52 + 105) / 211)
exe '2resize ' . ((&lines * 22 + 24) / 49)
exe 'vert 2resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 3resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 4resize ' . ((&columns * 52 + 105) / 211)
exe 'vert 5resize ' . ((&columns * 52 + 105) / 211)
tabnext
edit /mnt/c/Users/Lenovo/Documents/Programming/util/test/Test.go
argglobal
balt /mnt/c/Users/Lenovo/Documents/Programming/util/io/argparse/Parser.go
setlocal fdm=manual
setlocal fde=0
setlocal fmr={{{,}}}
setlocal fdi=#
setlocal fdl=0
setlocal fml=1
setlocal fdn=20
setlocal fen
silent! normal! zE
let &fdl = &fdl
let s:l = 45 - ((44 * winheight(0) + 23) / 46)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 45
normal! 012|
lcd /mnt/c/Users/Lenovo/Documents/Programming/util
tabnext 4
set stal=1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
set hlsearch
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
