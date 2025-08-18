package shortcuts

import "strings"

type KeyBoardKey string

const (
	F1          KeyBoardKey = "f1"
	F2          KeyBoardKey = "f2"
	F3          KeyBoardKey = "f3"
	F4          KeyBoardKey = "f4"
	F5          KeyBoardKey = "f5"
	F6          KeyBoardKey = "f6"
	F7          KeyBoardKey = "f7"
	F8          KeyBoardKey = "f8"
	F9          KeyBoardKey = "f9"
	F10         KeyBoardKey = "f10"
	F11         KeyBoardKey = "f11"
	F12         KeyBoardKey = "f12"
	A           KeyBoardKey = "a"
	B           KeyBoardKey = "b"
	C           KeyBoardKey = "c"
	D           KeyBoardKey = "d"
	E           KeyBoardKey = "e"
	F           KeyBoardKey = "f" //reserved for full screen
	G           KeyBoardKey = "g"
	H           KeyBoardKey = "h"
	I           KeyBoardKey = "i"
	J           KeyBoardKey = "j"
	K           KeyBoardKey = "k"
	L           KeyBoardKey = "l"
	M           KeyBoardKey = "m"
	N           KeyBoardKey = "n"
	O           KeyBoardKey = "o"
	P           KeyBoardKey = "p"
	Q           KeyBoardKey = "q"
	R           KeyBoardKey = "r"
	S           KeyBoardKey = "s"
	T           KeyBoardKey = "t"
	U           KeyBoardKey = "u"
	V           KeyBoardKey = "v"
	W           KeyBoardKey = "w"
	X           KeyBoardKey = "x"
	Y           KeyBoardKey = "y"
	Z           KeyBoardKey = "z"
	NumberZero  KeyBoardKey = "0"
	NumberOne   KeyBoardKey = "1"
	NumberTwo   KeyBoardKey = "2"
	NumberThree KeyBoardKey = "3"
	NumberFour  KeyBoardKey = "4"
	NumberFive  KeyBoardKey = "5"
	NumberSix   KeyBoardKey = "6"
	NumberSeven KeyBoardKey = "7"
	NumberEight KeyBoardKey = "8"
	NumberNine  KeyBoardKey = "9"
	NumberTen   KeyBoardKey = "10"
	SemiColon   KeyBoardKey = ";" //does not work as hot key with karabiner
)

var AllowedKeyBoardKeys = []KeyBoardKey{
	F1, F2, F3, F4, F5, F6, F7, F8, F9, F10, F11, F12,
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z,
	NumberZero, NumberOne, NumberTwo, NumberThree, NumberFour, NumberFive, NumberSix, NumberSeven, NumberEight, NumberNine, NumberTen,
	SemiColon,
}

func IsValidKeyBoardKey(key string) bool {
	key = strings.ToLower(key)
	for _, k := range AllowedKeyBoardKeys {
		if string(k) == key {
			return true
		}
	}
	return false
}
