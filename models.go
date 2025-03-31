package slugly

import "time"

type PrintSpeed time.Duration

const (
    SlowPrintSpeed   PrintSpeed = PrintSpeed(100 * time.Millisecond)
    MediumPrintSpeed PrintSpeed = PrintSpeed(50 * time.Millisecond)
    FastPrintSpeed   PrintSpeed = PrintSpeed(25 * time.Millisecond)
)

type PauseCharacter struct {
    Character      string
    PauseMultiplier int 
}

type SlugConfig struct {
    Speed                       PrintSpeed
    CustomPauseCharacters       []PauseCharacter
    RemoveCustomPauseCharacters []PauseCharacter
    Writer                      SlugWriter
}
