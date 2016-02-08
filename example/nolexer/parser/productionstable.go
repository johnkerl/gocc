package parser

import (
	"fmt"

	"github.com/goccmack/gocc/example/nolexer/token"
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
		String: `S' : Hello	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Hello : Saying name	<< func () (Attrib, error) {
							fmt.Println(string(X[1].(*token.Token).Lit)); 
                       		return nil, nil} () >>`,
		Id:         "Hello",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return func() (Attrib, error) {
				fmt.Println(string(X[1].(*token.Token).Lit))
				return nil, nil
			}()
		},
	},
	ProdTabEntry{
		String: `Saying : "hello"	<< func () (Attrib, error) {
							fmt.Print("hello "); 
                       		return nil, nil} () >>`,
		Id:         "Saying",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return func() (Attrib, error) {
				fmt.Print("hello ")
				return nil, nil
			}()
		},
	},
	ProdTabEntry{
		String: `Saying : "hiya"	<< func () (Attrib, error) {
							fmt.Print("hiya "); 
                            return nil, nil} () >>`,
		Id:         "Saying",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return func() (Attrib, error) {
				fmt.Print("hiya ")
				return nil, nil
			}()
		},
	},
}
