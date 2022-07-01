#!/usr/bin/env python3.7

import chess.pgn
import json


pgn = open("./lichess_db_standard_rated_2013-01.pgn")

game = chess.pgn.read_game(pgn)

result = {
    "description": "Imported Test Cases",
    "testCases": []
}

count = 0

while game:
    game = chess.pgn.read_game(pgn)
    board = game.board()
    for game_move in game.mainline_moves():
        board.push(game_move)
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

        # print(result)
        # exit()


    count += 1
    if count > 1000:
        break

print(json.dumps(result))
