#!/usr/bin/env bash

pgn-extract --stopafter 1000 --nocomments --notags -C --quiet --noduplicates -WFEN $1 | sed '/^$/d' > `echo $1 | sed s/.pgn/.fens/`
