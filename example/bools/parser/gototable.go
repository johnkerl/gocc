/*
 */
package parser

const numNTSymbols = 6

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // BoolExpr
		2,  // BoolExpr1
		3,  // Val
		7,  // CompareExpr
		8,  // SubStringExpr

	},
	gotoRow{ // S1

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S2

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S3

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S4

		-1, // S'
		13, // BoolExpr
		14, // BoolExpr1
		15, // Val
		19, // CompareExpr
		20, // SubStringExpr

	},
	gotoRow{ // S5

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S6

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S7

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S8

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S9

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S10

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S11

		-1, // S'
		26, // BoolExpr
		27, // BoolExpr1
		3,  // Val
		7,  // CompareExpr
		8,  // SubStringExpr

	},
	gotoRow{ // S12

		-1, // S'
		26, // BoolExpr
		28, // BoolExpr1
		3,  // Val
		7,  // CompareExpr
		8,  // SubStringExpr

	},
	gotoRow{ // S13

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S14

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S15

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S16

		-1, // S'
		32, // BoolExpr
		14, // BoolExpr1
		15, // Val
		19, // CompareExpr
		20, // SubStringExpr

	},
	gotoRow{ // S17

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S18

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S19

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S20

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S21

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S22

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S23

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S24

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S25

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S26

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S27

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S28

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S29

		-1, // S'
		39, // BoolExpr
		40, // BoolExpr1
		15, // Val
		19, // CompareExpr
		20, // SubStringExpr

	},
	gotoRow{ // S30

		-1, // S'
		39, // BoolExpr
		41, // BoolExpr1
		15, // Val
		19, // CompareExpr
		20, // SubStringExpr

	},
	gotoRow{ // S31

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S32

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S33

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S34

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S35

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S36

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S37

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S38

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S39

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S40

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S41

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S42

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S43

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S44

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S45

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
}
