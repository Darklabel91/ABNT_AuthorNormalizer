package Abnt

import (
	"errors"
	"strings"
)

//DataABNT main struct containing authors name, abnt format and short version of abnt.
type DataABNT struct {
	AuthorName string `json:"AuthorName,omitempty"`
	ABNT       string `json:"abnt,omitempty"`
	ABNTShort  string `json:"abnt_short,omitempty"`
}

//InitialABNT is Test struct that assemble the first part of abnt format with the last name and possible jr name
type InitialABNT struct {
	ABNTFristPart string
	LastName      string
	JRName        string
}

//TransformABNT receives Test name and returns the same name on ABNT format
func TransformABNT(authorName string) (DataABNT, error) {
	words, err := splitName(authorName)
	if err != nil {
		return DataABNT{}, err
	}

	initAbnt, err := returnInitialName(words)
	if err != nil {
		return DataABNT{}, err
	}

	middleName, err := returnMiddleName(words, initAbnt)
	if err != nil {
		return DataABNT{}, err
	}

	shortMiddleName, err := returnShortMiddleName(words, initAbnt)
	if err != nil {
		return DataABNT{}, err
	}

	abnt := strings.TrimSpace(initAbnt.ABNTFristPart + middleName)
	abntShort := strings.TrimSpace(initAbnt.ABNTFristPart + shortMiddleName)

	return DataABNT{
		AuthorName: authorName,
		ABNT:       abnt,
		ABNTShort:  abntShort,
	}, nil

}

//assemble the first part of ABNT format containing the last name and eventually the jrName
func returnInitialName(words []string) (InitialABNT, error) {
	if len(words) < 1 {
		return InitialABNT{}, errors.New("the name contains only one word: the minimum is 2")
	}

	lastWord := words[len(words)-1]

	var jrName string
	var lastName string
	if isJuniorName(lastWord) {
		lastName = words[len(words)-2]
		jrName = lastWord

		return InitialABNT{
			ABNTFristPart: strings.ToUpper(lastName) + " " + strings.ToUpper(jrName) + ", ",
			LastName:      lastName,
			JRName:        jrName,
		}, nil
	} else {
		return InitialABNT{
			ABNTFristPart: strings.ToUpper(lastWord) + ", ",
			LastName:      lastWord,
			JRName:        "",
		}, nil
	}
}

//returns all names that are not last name or jrName
func returnMiddleName(words []string, initialName InitialABNT) (string, error) {
	if len(words) < 1 {
		return "", errors.New("the name contains only one word: the minimum is 2")
	}

	var middleNames string
	for _, name := range words {
		if name != initialName.LastName && name != initialName.JRName {
			if isPreposition(name) {
				middleNames = middleNames + strings.ToLower(name) + " "
			} else {
				middleNames = middleNames + strings.ToUpper(name[0:1]) + strings.ToLower(name[1:]) + " "
			}
		}
	}

	return middleNames, nil
}

//returns initial letters of all names that are not last name or jrName
func returnShortMiddleName(words []string, initialName InitialABNT) (string, error) {
	if len(words) < 1 {
		return "", errors.New("the name contains only one word: the minimum is 2")
	}

	var middleNames string
	for _, name := range words {
		if name != initialName.LastName && name != initialName.JRName {
			if isPreposition(name) {
				middleNames = middleNames + ""
			} else {
				middleNames = middleNames + strings.ToUpper(name[0:1]) + ". "
			}
		}
	}

	return middleNames, nil
}

// return true if it is one of the twelve possible jr. names
//  júnior
//  junior
//  jr
//  jr.
//  filho
//  filha
//  neto
//  neta
//  segundo
//  segunda
//  terceiro
// 	terceira
func isJuniorName(word string) bool {
	juniorNameArray := []string{"filho", "filha", "neto", "neta", "junior", "júnior", "jr", "jr.", "segundo", "segunda", "terceiro", "terceira"}

	for _, junior := range juniorNameArray {
		if strings.ToLower(word) == junior {
			return true
		}
	}

	return false
}

//return true if it is one of 6 possible prepositions:
//  do
//  da
//  de
//  dos
//  das
//  e
func isPreposition(word string) bool {
	prepositionArray := []string{"do", "da", "de", "dos", "das", "e"}

	for _, prepArray := range prepositionArray {
		if strings.ToLower(word) == prepArray {
			return true
		}
	}

	return false
}
