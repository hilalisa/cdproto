package audits

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// GetEncodedResponseEncoding the encoding to use.
type GetEncodedResponseEncoding string

// String returns the GetEncodedResponseEncoding as string value.
func (t GetEncodedResponseEncoding) String() string {
	return string(t)
}

// GetEncodedResponseEncoding values.
const (
	GetEncodedResponseEncodingWebp GetEncodedResponseEncoding = "webp"
	GetEncodedResponseEncodingJpeg GetEncodedResponseEncoding = "jpeg"
	GetEncodedResponseEncodingPng  GetEncodedResponseEncoding = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t GetEncodedResponseEncoding) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t GetEncodedResponseEncoding) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch GetEncodedResponseEncoding(in.String()) {
	case GetEncodedResponseEncodingWebp:
		*t = GetEncodedResponseEncodingWebp
	case GetEncodedResponseEncodingJpeg:
		*t = GetEncodedResponseEncodingJpeg
	case GetEncodedResponseEncodingPng:
		*t = GetEncodedResponseEncodingPng

	default:
		in.AddError(errors.New("unknown GetEncodedResponseEncoding value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
