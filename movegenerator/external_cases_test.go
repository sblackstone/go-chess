package movegenerator

import (
  "testing"
  "github.com/sblackstone/go-chess/fen"
  //"fmt"
  "encoding/json"
  "io/ioutil"
  "errors"
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
  for _, expect := range(t.Expected) {
    result = append(result, expect.Fen)
  }
  sort.Strings(result)
  return result
}

func TestLichessBank(t *testing.T) {
  testFromFile(t, "../test-cases/generated.json")
}


func TestCastling(t *testing.T) {
  testFromFile(t, "../test-cases/chessmovegen/castling.json")
}

func TestFamous(t *testing.T) {
  testFromFile(t, "../test-cases/chessmovegen/famous.json")
}

func TestPawns(t *testing.T) {
  testFromFile(t, "../test-cases/chessmovegen/pawns.json")
}

func TestStandard(t *testing.T) {
  testFromFile(t, "../test-cases/chessmovegen/standard.json")
}

func TestTaxing(t *testing.T) {
  testFromFile(t, "../test-cases/chessmovegen/taxing.json")
}



func testFromFile(t *testing.T, fileName string ) error {
  byteValue, err := ioutil.ReadFile(fileName)
  if err != nil {
    return err
  }

  var testCaseFile TestCaseFile;
  json.Unmarshal(byteValue, &testCaseFile)

  for _, tc := range(testCaseFile.TestCases) {
    b, err := fen.FromFEN(tc.Start.Fen)
    if err != nil {
      return err
    }

    successors := GenLegalSuccessors(b)

    var fens []string
    for _, succ := range(successors) {
      f, err := fen.ToFEN(succ)
      if err != nil {
        return err
      }
      fens = append(fens, f)
    }

    sort.Strings(fens)
    expected := tc.ExpectedFens()

    if !reflect.DeepEqual(fens, expected) {

      t.Errorf("Initial %v", tc.Start.Fen)

      for _, expect := range(expected) {
        seen := false
        for _, fen := range(fens) {
          if expect == fen {
            seen = true
            break
          }
        }
        if !seen {
          t.Errorf("Missing    %v", expect)
        }
      }


      for _, fen := range(fens) {
        seen := false
        for _, expect := range(expected) {
          if fen == expect {
            seen = true
            break
          }
        }
        if !seen {
          t.Errorf("Unexpected %v", fen)
        }
      }
      return errors.New("False")
    }
  }
  return nil
}
