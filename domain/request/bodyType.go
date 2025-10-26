package request

type BodyType int

const (
	BodyTypeNull = iota
	BodyTypePlainText
	BodyTypeJSON
	BodyTypeJavascript
	BodyTypeHTML
	// TODO: add support to pdf/files
	// BodyTypeFile
)

func (bt BodyType) String() string {
	return [...]string{
		"",                 // BodyTypeNull = iota
		"text/plain",       // BodyTypePlainText
		"application/json", // BodyTypeJSON
		"text/javascript",  // BodyTypeJavascript
		"text/html",        // BodyTypeHTML
		// "", // BodyTypeFile
	}[bt]
}
