package helpers

import (
	"bytes"
	"fmt"
)

func generateGopher(text string) string {
	var out bytes.Buffer

    out.WriteString("          ,_---~~~~~----._         \n")
    out.WriteString("   _,,_,*^____      _____``*g*\\\"*, \n")
    out.WriteString("  / __/ /'     ^.  /      \\ ^@q   f \n")
    out.WriteString(" [  @f | @))    |  | @))   l  0 _/  \n")
    out.WriteString("  \\`/   \\~____ / __ \\_____/    \\   \n")
    out.WriteString("   |           _l__l_           I   \n")
    out.WriteString("   }          [______]           I  \n")
    out.WriteString(fmt.Sprintf("   ]           | | |             |   (%s)\n", text))
    out.WriteString("   ]            ~ ~             |  \n")
    out.WriteString("   |                            |   \n")
    out.WriteString("    |                           |   ")

    return out.String()
}

func GetTextAlert(text string) string {
    return generateGopher(text)
}
