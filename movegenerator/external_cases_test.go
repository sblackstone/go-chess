package movegenerator

import (
  "testing"
  "github.com/sblackstone/go-chess/fen"
  "fmt"
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
  for i := range(t.Expected) {
    result = append(result, t.Expected[i].Fen)
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

  //fmt.Printf("byteValue = %s\n", byteValue)
  var testCaseFile TestCaseFile;
  json.Unmarshal(byteValue, &testCaseFile)

  fmt.Printf("Processing %v with %v entries\n", testCaseFile.Description, len(testCaseFile.TestCases))

  for i := range(testCaseFile.TestCases) {
    fmt.Printf("Processing #%v\n", i)
    tc := testCaseFile.TestCases[i]
    b, err := fen.FromFEN(tc.Start.Fen)
    if err != nil {
      return err
    }

    successors := GenLegalSuccessors(b)

    var fens []string
    for i := range(successors) {
      f, err := fen.ToFEN(successors[i])
      if err != nil {
        return err
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
      return errors.New("False")
    }
  }
  //t.Errorf("BLARG")
  return nil
}
