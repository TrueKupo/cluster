package blockchain

import (
	"github.com/tyler-smith/go-bip39/wordlists"
)

var (
	lang_wordlist map[int32][]string = map[int32][]string{
		0: wordlists.English,
		1: wordlists.French,
		2: wordlists.Italian,
		3: wordlists.Spanish,
		4: wordlists.Czech,
		5: wordlists.Japanese,
		6: wordlists.Korean,
		7: wordlists.ChineseSimplified,
		8: wordlists.ChineseTraditional,
	}
)

func WordList(lang_idx int32) []string {
	if wl, ok := lang_wordlist[lang_idx]; ok {
		return wl
	} else {
		return wordlists.English
	}
}
