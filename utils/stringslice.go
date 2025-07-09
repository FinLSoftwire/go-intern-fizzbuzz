package utils

type StringSlice []string

func (currentWordsSlice *StringSlice) InsertBeforeFirstB(newWord string) {
	newWordsSlice := make([]string, 0)
	for wordIndex, currentWord := range *currentWordsSlice {
		if currentWord[0] == 'B' {
			*currentWordsSlice = append(append(newWordsSlice, newWord), (*currentWordsSlice)[wordIndex:]...)
			return
		} else {
			newWordsSlice = append(newWordsSlice, currentWord)
		}
	}
	*currentWordsSlice = append(newWordsSlice, newWord)
}

func (currentWordsSlice *StringSlice) AppendInPlace(newWord string) {
	*currentWordsSlice = append(*currentWordsSlice, newWord)
}

func (currentWordsSlice *StringSlice) SetToSingleElement(word string) {
	*currentWordsSlice = StringSlice{word}
}
