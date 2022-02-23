package chessgo

type Board struct {
	pawns      [2]bitboard
	rooks      [2]bitboard
	bishops    [2]bitboard
	knights    [2]bitboard
	queens     [2]bitboard
	kings      [2]bitboard
	attackFrom [64]bitboard
	attackTo   [64]bitboard
	rot90R     bitboard
	rot45R     bitboard
	rot45L     bitboard
}

func (board *Board) move(startpos int, endpos int) {

}
