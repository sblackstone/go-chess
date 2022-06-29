#!/usr/bin/env python3.7

import chess.pgn
import json


pgn = open("./lichess_db_standard_rated_2013-01.pgn")

game = chess.pgn.read_game(pgn)

result = {
    "description": "Imported Test Cases",
    "testCases": []
}

max = 0

while game:
    game = chess.pgn.read_game(pgn)
    board = game.board()
    expected = []
    for move in board.legal_moves:
        board.push(move)
        expected.append({
            "move": move.uci(),
            "fen": board.fen(en_passant="fen")
        })
        board.pop()
    result["testCases"].append({
        "start": { "fen": board.fen(en_passant="fen") },
        "expected": expected
    })
    max += 1
    if max > 24999:
        break

print(json.dumps(result))