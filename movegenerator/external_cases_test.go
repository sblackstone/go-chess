package movegenerator

import (
  "testing"
  "github.com/sblackstone/go-chess/fen"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "sort"
  "reflect"
)



type StartingPosition struct {
    Description string `json:"description"`
    Fen string `json:"fen"`
}

type ExpectedPosition struct {
  Move string `json:"move"`
  Fen string `json:"fen"`
}

type TestCase struct {
  Start StartingPosition `json:"start"`
  Expected []ExpectedPosition `json":expected"`
}

type TestCaseFile struct {
  Description string `json:"description"`
  TestCases []TestCase `json:"testCases"`
}


func (t *TestCase) ExpectedFens() []string {
  var result []string
  for i := range(t.Expected) {
    result = append(result, t.Expected[i].Fen)
  }
  sort.Strings(result)
  return result
}

func TestFromFile(t *testing.T) {
  byteValue, err := ioutil.ReadFile("../chessmovegen-test-suites/src/main/resources/testcases/famous.json")
  if err != nil {
    t.Errorf("%v", err)
    return
  }

  //fmt.Printf("byteValue = %s\n", byteValue)
  var testCaseFile TestCaseFile;
  json.Unmarshal(byteValue, &testCaseFile)

  fmt.Printf("Processing %v\n", testCaseFile.Description)

  for i := range(testCaseFile.TestCases) {
    tc := testCaseFile.TestCases[i]
    b, err := fen.FromFEN(tc.Start.Fen)
    if err != nil {
      t.Errorf("%v", err)
      return
    }

    successors := GenLegalSuccessors(b)

    var fens []string
    for i := range(successors) {
      f, err := fen.ToFEN(successors[i])
      if err != nil {
        t.Errorf("%v", err)
        return
      }
      fens = append(fens, f)
    }

    sort.Strings(fens)
    expected := tc.ExpectedFens()

    if !reflect.DeepEqual(fens, expected) {

      t.Errorf("Initial %v", tc.Start.Fen)

      for j := range(expected) {
        seen := false
        for k := range(fens) {
          if expected[j] == fens[k] {
            seen = true
            break
          }
        }
        if !seen {
          t.Errorf("Missing    %v", expected[j])
        }
      }


      for j := range(fens) {
        seen := false
        for k := range(expected) {
          if fens[j] == expected[k] {
            seen = true
            break
          }
        }
        if !seen {
          t.Errorf("Unexpected %v", fens[j])
        }
      }
  	}
  }
}

// func TestLetsTryIt(t *testing.T) {
//   b, err := fen.FromFEN("8/ppp3p1/8/8/3p4/8/1ppp2K1/brk2Q1n b - - 12 7")
//   if err != nil {
//     t.Errorf("Err: %v", err)
//   }
//
//   successors := GenLegalSuccessors(b)
//
//   for i := range(successors) {
//     fenStr, err := fen.ToFEN(successors[i])
//     if err != nil {
//       t.Errorf("%v", err)
//     }
//     fmt.Println(fenStr)
//   }
//
//   t.Errorf("BLARG")
//
// }


/*
{
  "description": "Some common positions.",
  "testCases": [
    {
      "start": {
        "description": "Standard starting position.",
        "fen": "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
      },
      "expected": [
        {
          "move": "Na3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/N7/PPPPPPPP/R1BQKBNR b KQkq - 1 1"
        },
        {
          "move": "Nc3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/2N5/PPPPPPPP/R1BQKBNR b KQkq - 1 1"
        },
        {
          "move": "Nf3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R b KQkq - 1 1"
        },
        {
          "move": "Nh3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/7N/PPPPPPPP/RNBQKB1R b KQkq - 1 1"
        },
        {
          "move": "a3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "a4",
          "fen": "rnbqkbnr/pppppppp/8/8/P7/8/1PPPPPPP/RNBQKBNR b KQkq a3 0 1"
        },
        {
          "move": "b3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/1P6/P1PPPPPP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "b4",
          "fen": "rnbqkbnr/pppppppp/8/8/1P6/8/P1PPPPPP/RNBQKBNR b KQkq b3 0 1"
        },
        {
          "move": "c3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/2P5/PP1PPPPP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "c4",
          "fen": "rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b KQkq c3 0 1"
        },
        {
          "move": "d3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/3P4/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "d4",
          "fen": "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1"
        },
        {
          "move": "e3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "e4",
          "fen": "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
        },
        {
          "move": "f3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/5P2/PPPPP1PP/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "f4",
          "fen": "rnbqkbnr/pppppppp/8/8/5P2/8/PPPPP1PP/RNBQKBNR b KQkq f3 0 1"
        },
        {
          "move": "g3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/6P1/PPPPPP1P/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "g4",
          "fen": "rnbqkbnr/pppppppp/8/8/6P1/8/PPPPPP1P/RNBQKBNR b KQkq g3 0 1"
        },
        {
          "move": "h3",
          "fen": "rnbqkbnr/pppppppp/8/8/8/7P/PPPPPPP1/RNBQKBNR b KQkq - 0 1"
        },
        {
          "move": "h4",
          "fen": "rnbqkbnr/pppppppp/8/8/7P/8/PPPPPPP1/RNBQKBNR b KQkq h3 0 1"
        }
      ]
    },
    {
      "start": {
        "fen": "8/ppp3p1/8/8/3p4/5Q2/1ppp2K1/brk4n w - - 11 7"
      },
      "expected": [
        {
          "move": "Kf1",
          "fen": "8/ppp3p1/8/8/3p4/5Q2/1ppp4/brk2K1n b - - 12 7"
        },
        {
          "move": "Kg1",
          "fen": "8/ppp3p1/8/8/3p4/5Q2/1ppp4/brk3Kn b - - 12 7"
        },
        {
          "move": "Kh1",
          "fen": "8/ppp3p1/8/8/3p4/5Q2/1ppp4/brk4K b - - 0 7"
        },
        {
          "move": "Kh2",
          "fen": "8/ppp3p1/8/8/3p4/5Q2/1ppp3K/brk4n b - - 12 7"
        },
        {
          "move": "Kh3",
          "fen": "8/ppp3p1/8/8/3p4/5Q1K/1ppp4/brk4n b - - 12 7"
        },
        {
          "move": "Qe2",
          "fen": "8/ppp3p1/8/8/3p4/8/1pppQ1K1/brk4n b - - 12 7"
        },
        {
          "move": "Qd1+",
          "fen": "8/ppp3p1/8/8/3p4/8/1ppp2K1/brkQ3n b - - 12 7"
        },
        {
          "move": "Qf2",
          "fen": "8/ppp3p1/8/8/3p4/8/1ppp1QK1/brk4n b - - 12 7"
        },
        {
          "move": "Qf1+",
          "fen": "8/ppp3p1/8/8/3p4/8/1ppp2K1/brk2Q1n b - - 12 7"
        },
        {
          "move": "Qe3",
          "fen": "8/ppp3p1/8/8/3p4/4Q3/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qd3",
          "fen": "8/ppp3p1/8/8/3p4/3Q4/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qc3",
          "fen": "8/ppp3p1/8/8/3p4/2Q5/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qb3",
          "fen": "8/ppp3p1/8/8/3p4/1Q6/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qa3",
          "fen": "8/ppp3p1/8/8/3p4/Q7/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qg3",
          "fen": "8/ppp3p1/8/8/3p4/6Q1/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qh3",
          "fen": "8/ppp3p1/8/8/3p4/7Q/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qe4",
          "fen": "8/ppp3p1/8/8/3pQ3/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qd5",
          "fen": "8/ppp3p1/8/3Q4/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qc6",
          "fen": "8/ppp3p1/2Q5/8/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qxb7",
          "fen": "8/pQp3p1/8/8/3p4/8/1ppp2K1/brk4n b - - 0 7"
        },
        {
          "move": "Qf4",
          "fen": "8/ppp3p1/8/8/3p1Q2/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qf5",
          "fen": "8/ppp3p1/8/5Q2/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qf6",
          "fen": "8/ppp3p1/5Q2/8/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qf7",
          "fen": "8/ppp2Qp1/8/8/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qf8",
          "fen": "5Q2/ppp3p1/8/8/3p4/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qg4",
          "fen": "8/ppp3p1/8/8/3p2Q1/8/1ppp2K1/brk4n b - - 12 7"
        },
        {
          "move": "Qh5",
          "fen": "8/ppp3p1/8/7Q/3p4/8/1ppp2K1/brk4n b - - 12 7"
        }
      ]
    }
  ]
}
*/
