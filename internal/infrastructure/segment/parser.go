package segment

import (
	"github.com/cipma/smsbin/app/sms"
	"strings"
)

var unicodeToGSM7 = map[rune]byte{
	'@': 0x00, '£': 0x01, '$': 0x02, '¥': 0x03, 'è': 0x04, 'é': 0x05, 'ù': 0x06, 'ì': 0x07,
	'ò': 0x08, 'Ç': 0x09, '\n': 0x0a, 'Ø': 0x0b, 'ø': 0x0c, '\r': 0x0d, 'Å': 0x0e, 'å': 0x0f,
	'Δ': 0x10, '_': 0x11, 'Φ': 0x12, 'Γ': 0x13, 'Λ': 0x14, 'Ω': 0x15, 'Π': 0x16, 'Ψ': 0x17,
	'Σ': 0x18, 'Θ': 0x19, 'Ξ': 0x1a /* 0x1B */, 'Æ': 0x1c, 'æ': 0x1d, 'ß': 0x1e, 'É': 0x1f,
	' ': 0x20, '!': 0x21, '"': 0x22, '#': 0x23, '¤': 0x24, '%': 0x25, '&': 0x26, '\'': 0x27,
	'(': 0x28, ')': 0x29, '*': 0x2a, '+': 0x2b, ',': 0x2c, '-': 0x2d, '.': 0x2e, '/': 0x2f,
	'0': 0x30, '1': 0x31, '2': 0x32, '3': 0x33, '4': 0x34, '5': 0x35, '6': 0x36, '7': 0x37,
	'8': 0x38, '9': 0x39, ':': 0x3a, ';': 0x3b, '<': 0x3c, '=': 0x3d, '>': 0x3e, '?': 0x3f,
	'¡': 0x40, 'A': 0x41, 'B': 0x42, 'C': 0x43, 'D': 0x44, 'E': 0x45, 'F': 0x46, 'G': 0x47,
	'H': 0x48, 'I': 0x49, 'J': 0x4a, 'K': 0x4b, 'L': 0x4c, 'M': 0x4d, 'N': 0x4e, 'O': 0x4f,
	'P': 0x50, 'Q': 0x51, 'R': 0x52, 'S': 0x53, 'T': 0x54, 'U': 0x55, 'V': 0x56, 'W': 0x57,
	'X': 0x58, 'Y': 0x59, 'Z': 0x5a, 'Ä': 0x5b, 'Ö': 0x5c, 'Ñ': 0x5d, 'Ü': 0x5e, '§': 0x5f,
	'¿': 0x60, 'a': 0x61, 'b': 0x62, 'c': 0x63, 'd': 0x64, 'e': 0x65, 'f': 0x66, 'g': 0x67,
	'h': 0x68, 'i': 0x69, 'j': 0x6a, 'k': 0x6b, 'l': 0x6c, 'm': 0x6d, 'n': 0x6e, 'o': 0x6f,
	'p': 0x70, 'q': 0x71, 'r': 0x72, 's': 0x73, 't': 0x74, 'u': 0x75, 'v': 0x76, 'w': 0x77,
	'x': 0x78, 'y': 0x79, 'z': 0x7a, 'ä': 0x7b, 'ö': 0x7c, 'ñ': 0x7d, 'ü': 0x7e, 'à': 0x7f,
}
var unicodeToExtendedGSM7 = map[rune]byte{
	'\f': 0x0A, '^': 0x14, '{': 0x28, '}': 0x29, '\\': 0x2F, '[': 0x3C, '~': 0x3D, ']': 0x3E, '|': 0x40, '€': 0x65,
}

const (
	maxGSMSize        int = 160
	maxGSMSegmentSize int = 153
	maxUCSSize        int = 70
	maxUCSSegmentSize int = 67
)

type Parser struct{}

func (p Parser) ExtractSegments(text string) sms.TextMetadata {
	isGSM := true
	totalSize := 0

	for _, r := range text {
		charSize, isGSMChar := calculateCharSize(r)
		totalSize += charSize
		isGSM = isGSM && isGSMChar
	}

	segmentSize := calculateSegmentSize(totalSize, isGSM)

	metadata := sms.TextMetadata{
		Segments:  make([]sms.Segment, 0),
		TotalSize: totalSize,
	}

	size := 0
	sgmt := sms.Segment{}
	sb := strings.Builder{}

	for _, r := range text {
		charSize, _ := calculateCharSize(r)

		if charSize+size > segmentSize {
			sgmt.Data = sb.String()
			sgmt.Size = size
			metadata.Segments = append(metadata.Segments, sgmt)

			size = 0
			sgmt = sms.Segment{}
			sb.Reset()
		}

		size += charSize
		sb.WriteRune(r)
	}

	if sb.Len() > 0 {
		sgmt.Data = sb.String()
		sgmt.Size = size
		metadata.Segments = append(metadata.Segments, sgmt)
	}

	return metadata
}

func calculateCharSize(r rune) (int, bool) {
	if _, ok := unicodeToGSM7[r]; ok {
		return 1, true
	} else if _, ok := unicodeToExtendedGSM7[r]; ok {
		return 2, true
	} else {
		return 1, false
	}
}

func calculateSegmentSize(totalSize int, isGSM bool) int {
	singleSegmentSize := maxGSMSize
	segmentSize := maxGSMSegmentSize

	if !isGSM {
		singleSegmentSize = maxUCSSize
		segmentSize = maxUCSSegmentSize
	}

	if totalSize <= singleSegmentSize {
		segmentSize = singleSegmentSize
	}

	return segmentSize
}
