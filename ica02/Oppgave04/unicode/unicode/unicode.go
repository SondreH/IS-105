package unicode

import (
  "fmt"
)
// Check which language and return either japanese, icelandic or alien
func Translate(expression string, language string) string {
  if language == "is" {
    return "\xE2\x80\x9C\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\xE2\x80\x9D\n"
  }
  if language == "jp" {
    return "\xE2\x80\x9C\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\xE2\x80\x9D\n"
  }
	return "alien language"
}

// Print the unicode symbols of this UTF-8 code
func UnicodeCodePointDemo() {
  fmt.Println("\xf0\x9F\x98\x80")
  fmt.Println("\xf0\x9F\x98\x97")
  fmt.Println("\xf0\x9F\x98")
  fmt.Println("\xf0\x9F")
  fmt.Println("\xf0")
}
