package parser

import (
	"code.google.com/p/gocc/example/calc/token"
	"code.google.com/p/gocc/example/calc/util"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Calc	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Calc : Expr	<<  >>`,
		Id:         "Calc",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "+" Term	<< X[0].(int64) + X[2].(int64), nil >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) + X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term "*" Factor	<< X[0].(int64) * X[2].(int64), nil >>`,
		Id:         "Term",
		NTType:     3,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0].(int64) * X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id:         "Factor",
		NTType:     4,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int64	<< util.IntValue(X[0].(*token.Token).Lit) >>`,
		Id:         "Factor",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return util.IntValue(X[0].(*token.Token).Lit)
		},
	},
}
