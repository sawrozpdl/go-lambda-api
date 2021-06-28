package utils

import "unicode/utf8"

func Short( s string, i int) string {
    if len( s ) < i {
        return s
    }
    if utf8.ValidString( s[:i] ) {
        return s[:i]
    }

    return s[:i+1] 
}
