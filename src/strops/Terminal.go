package strops

import "fmt"

type (
	FGColor int
	BGColor int
	Effect  int
)

const (
	FGBlack FGColor = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
	FGDefault = 39
)

const (
	FGBrightBlack FGColor = iota + 90
	FGBrightRed
	FGBrightGreen
	FGBrightYellow
	FGBrightBlue
	FGBrightMagenta
	FGBrightCyan
	FGBrightWhite
	FGBrightDefault = 99
)

const (
	BGBlack BGColor = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
	BGDefault = 49
)

const (
	BGBrightBlack BGColor = iota + 100
	BGBrightRed
	BGBrightGreen
	BGBrightYellow
	BGBrightBlue
	BGBrightMagenta
	BGBrightCyan
	BGBrightWhite
	BGBrightDefault = 109
)

const (
	BoldEffect       Effect = 1
	DimEffect               = 2
	UnderlyineEffect        = 4
	BlinkEffect             = 5
	ReverseEffect           = 7
	HideEffect              = 8
)

func NoEffectOrColor(fmtStr string, args ...any) string {
	return fmt.Sprintf(fmtStr, args...)
}

func ColorApplier(fg FGColor, bg BGColor) func(fmtStr string, args ...any) string {
	return func(fmtStr string, args ...any) string {
		fmtStr = fmt.Sprintf("\033[%d;%dm%s\033[0m", fg, bg, fmtStr)
		return fmt.Sprintf(fmtStr, args...)
	}
}

func EffectApplier(effect Effect) func(fmtStr string, args ...any) string {
	return func(fmtStr string, args ...any) string {
		fmtStr = fmt.Sprintf("\033[%dm%s\033[0m", effect, fmtStr)
		return fmt.Sprintf(fmtStr, args...)
	}
}

func ApplyColor(s string, fg FGColor, bg BGColor) string {
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", fg, bg, s)
}

func ApplyEffect(s string, effect Effect) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", effect, s)
}

func FmtApplyColor(fg FGColor, bg BGColor) string {
	return fmt.Sprintf("\033[%d;%dm%%s\033[0m", fg, bg)
}

func FmtApplyEffect(effect Effect) string {
	return fmt.Sprintf("\033[%dm%%s\033[0m", effect)
}
