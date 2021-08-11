package sms

type PhoneNumberParser interface {
	Parse(number string, allowAlpha bool) PhoneNumberMetadata
}

type TextSegmentParser interface {
	ExtractSegments(text string) TextMetadata
}
