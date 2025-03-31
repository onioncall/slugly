package slugly

import "time"

// SlugPrint prints text at a modulated rate
func (c *SlugConfig) SlugPrint(text string) {
    if len(c.CustomPauseCharacters) > 0 {
        addCustomPauseCharacters(c.CustomPauseCharacters)
    }
    if len(c.RemoveCustomPauseCharacters) > 0 {
        removeCustomPauseCharacters(c.RemoveCustomPauseCharacters)
    }
    if len(c.CustomPauseCharacters)+len(c.RemoveCustomPauseCharacters) > 0 {
        setMaxPauseCharacterLength()
    }
    sortPauseCharacters()
    
    for i := range len(text) {
        pm := 1
        if i < len(text)-1 {
            pm = getPauseMultiplier(text, i)
        }
        char := string(text[i])
        c.Writer.WriteChar(char)
        time.Sleep(time.Duration(c.Speed * PrintSpeed(pm)))
    }
    
    c.Writer.WriteChar("\n")
    c.Writer.Flush()
}
