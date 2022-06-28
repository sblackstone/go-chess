#!/usr/bin/env python3.7

import chess.pgn

pgn = open("./lichess_db_standard_rated_2013-01.pgn")

game = chess.pgn.read_game(pgn)

while game:
    game = chess.pgn.read_game(pgn)
    board = game.board()
    for move in board.legal_moves:        
        print(move)
