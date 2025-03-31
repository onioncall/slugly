package slugly

import (
    "errors"
    "sort"
    "strings"
)

var ShortPauseCharacterLength int = 10
var LongPauseCharacterLength int = 20
var maxPauseCharacterLength int = 3

var pauseCharacters = []PauseCharacter {
    {"...", LongPauseCharacterLength},
    {"?", ShortPauseCharacterLength},
    {".", ShortPauseCharacterLength},
    {"!", ShortPauseCharacterLength},
}

func addCustomPauseCharacters(chars []PauseCharacter) error {
    for _, char := range chars {
        if len(char.Character) > 5 {
            return errors.New("A Custom Pause Character can not be more than five Characters (confusing, I know)")
        } 
        pauseCharacters = append(pauseCharacters, char)
    }
    return nil
}

func removeCustomPauseCharacters(chars []PauseCharacter) {
    // Todo
}

func sortPauseCharacters() {
    sort.Slice(pauseCharacters, func(i, j int) bool {
        return len(pauseCharacters[i].Character) > len(pauseCharacters[j].Character)
    })
}

func setMaxPauseCharacterLength() {
    for _, char := range pauseCharacters {
        if len(char.Character) > maxPauseCharacterLength {
            maxPauseCharacterLength = len(char.Character)
        }
    }
}

func getPauseMultiplier(text string, currentPos int) int {
    currentChar := string(text[currentPos])
    var charMapBackward string
    var charMapFull string
    for i := range maxPauseCharacterLength {
        if currentPos - i > 0 {
            charMapBackward = string(text[currentPos - i]) + charMapBackward
            charMapFull = string(text[currentPos - i]) + charMapFull
        }
        if currentPos + i < len(text) && i > 0 {
            charMapFull = charMapFull + string(text[currentPos + i])
        }        
    }
    for i := maxPauseCharacterLength; i > 0; i-- {
        for _, char := range pauseCharacters {
            if charMapBackward == char.Character {
                return char.PauseMultiplier
            } else if currentChar != char.Character && strings.Contains(charMapFull, char.Character) {
                break
            } else if currentChar == char.Character {
                return char.PauseMultiplier
            }
        }
    }
    return 1
}
