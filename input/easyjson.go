// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package input

import (
	json "encoding/json"
	cdp "github.com/knq/chromedp/cdp"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput(in *jlexer.Lexer, out *TouchPoint) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "state":
			(out.State).UnmarshalEasyJSON(in)
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "radiusX":
			out.RadiusX = int64(in.Int64())
		case "radiusY":
			out.RadiusY = int64(in.Int64())
		case "rotationAngle":
			out.RotationAngle = float64(in.Float64())
		case "force":
			out.Force = float64(in.Float64())
		case "id":
			out.ID = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput(out *jwriter.Writer, in TouchPoint) {
	out.RawByte('{')
	first := true
	_ = first
	if in.State != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"state\":")
		(in.State).MarshalEasyJSON(out)
	}
	if in.X != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"x\":")
		out.Int64(int64(in.X))
	}
	if in.Y != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"y\":")
		out.Int64(int64(in.Y))
	}
	if in.RadiusX != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"radiusX\":")
		out.Int64(int64(in.RadiusX))
	}
	if in.RadiusY != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"radiusY\":")
		out.Int64(int64(in.RadiusY))
	}
	if in.RotationAngle != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"rotationAngle\":")
		out.Float64(float64(in.RotationAngle))
	}
	if in.Force != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"force\":")
		out.Float64(float64(in.Force))
	}
	if in.ID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.Float64(float64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TouchPoint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TouchPoint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TouchPoint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TouchPoint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput1(in *jlexer.Lexer, out *SynthesizeTapGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "duration":
			out.Duration = int64(in.Int64())
		case "tapCount":
			out.TapCount = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput1(out *jwriter.Writer, in SynthesizeTapGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int64(int64(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int64(int64(in.Y))
	if in.Duration != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"duration\":")
		out.Int64(int64(in.Duration))
	}
	if in.TapCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"tapCount\":")
		out.Int64(int64(in.TapCount))
	}
	if in.GestureSourceType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"gestureSourceType\":")
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizeTapGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizeTapGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizeTapGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizeTapGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput2(in *jlexer.Lexer, out *SynthesizeScrollGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "xDistance":
			out.XDistance = int64(in.Int64())
		case "yDistance":
			out.YDistance = int64(in.Int64())
		case "xOverscroll":
			out.XOverscroll = int64(in.Int64())
		case "yOverscroll":
			out.YOverscroll = int64(in.Int64())
		case "preventFling":
			out.PreventFling = bool(in.Bool())
		case "speed":
			out.Speed = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		case "repeatCount":
			out.RepeatCount = int64(in.Int64())
		case "repeatDelayMs":
			out.RepeatDelayMs = int64(in.Int64())
		case "interactionMarkerName":
			out.InteractionMarkerName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput2(out *jwriter.Writer, in SynthesizeScrollGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int64(int64(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int64(int64(in.Y))
	if in.XDistance != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"xDistance\":")
		out.Int64(int64(in.XDistance))
	}
	if in.YDistance != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"yDistance\":")
		out.Int64(int64(in.YDistance))
	}
	if in.XOverscroll != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"xOverscroll\":")
		out.Int64(int64(in.XOverscroll))
	}
	if in.YOverscroll != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"yOverscroll\":")
		out.Int64(int64(in.YOverscroll))
	}
	if in.PreventFling {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"preventFling\":")
		out.Bool(bool(in.PreventFling))
	}
	if in.Speed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"speed\":")
		out.Int64(int64(in.Speed))
	}
	if in.GestureSourceType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"gestureSourceType\":")
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	if in.RepeatCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"repeatCount\":")
		out.Int64(int64(in.RepeatCount))
	}
	if in.RepeatDelayMs != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"repeatDelayMs\":")
		out.Int64(int64(in.RepeatDelayMs))
	}
	if in.InteractionMarkerName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"interactionMarkerName\":")
		out.String(string(in.InteractionMarkerName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizeScrollGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizeScrollGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizeScrollGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizeScrollGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput3(in *jlexer.Lexer, out *SynthesizePinchGestureParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "scaleFactor":
			out.ScaleFactor = float64(in.Float64())
		case "relativeSpeed":
			out.RelativeSpeed = int64(in.Int64())
		case "gestureSourceType":
			(out.GestureSourceType).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput3(out *jwriter.Writer, in SynthesizePinchGestureParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int64(int64(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int64(int64(in.Y))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scaleFactor\":")
	out.Float64(float64(in.ScaleFactor))
	if in.RelativeSpeed != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"relativeSpeed\":")
		out.Int64(int64(in.RelativeSpeed))
	}
	if in.GestureSourceType != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"gestureSourceType\":")
		(in.GestureSourceType).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SynthesizePinchGestureParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SynthesizePinchGestureParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SynthesizePinchGestureParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SynthesizePinchGestureParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput3(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput4(in *jlexer.Lexer, out *SetIgnoreInputEventsParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ignore":
			out.Ignore = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput4(out *jwriter.Writer, in SetIgnoreInputEventsParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ignore\":")
	out.Bool(bool(in.Ignore))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SetIgnoreInputEventsParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SetIgnoreInputEventsParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SetIgnoreInputEventsParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SetIgnoreInputEventsParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput4(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput5(in *jlexer.Lexer, out *EmulateTouchFromMouseEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(cdp.Timestamp)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "button":
			(out.Button).UnmarshalEasyJSON(in)
		case "deltaX":
			out.DeltaX = float64(in.Float64())
		case "deltaY":
			out.DeltaY = float64(in.Float64())
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "clickCount":
			out.ClickCount = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput5(out *jwriter.Writer, in EmulateTouchFromMouseEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	(in.Type).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int64(int64(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int64(int64(in.Y))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"timestamp\":")
	if in.Timestamp == nil {
		out.RawString("null")
	} else {
		(*in.Timestamp).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"button\":")
	(in.Button).MarshalEasyJSON(out)
	if in.DeltaX != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"deltaX\":")
		out.Float64(float64(in.DeltaX))
	}
	if in.DeltaY != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"deltaY\":")
		out.Float64(float64(in.DeltaY))
	}
	if in.Modifiers != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modifiers\":")
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.ClickCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"clickCount\":")
		out.Int64(int64(in.ClickCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EmulateTouchFromMouseEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EmulateTouchFromMouseEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EmulateTouchFromMouseEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EmulateTouchFromMouseEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput5(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput6(in *jlexer.Lexer, out *DispatchTouchEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "touchPoints":
			if in.IsNull() {
				in.Skip()
				out.TouchPoints = nil
			} else {
				in.Delim('[')
				if out.TouchPoints == nil {
					if !in.IsDelim(']') {
						out.TouchPoints = make([]*TouchPoint, 0, 8)
					} else {
						out.TouchPoints = []*TouchPoint{}
					}
				} else {
					out.TouchPoints = (out.TouchPoints)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *TouchPoint
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(TouchPoint)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.TouchPoints = append(out.TouchPoints, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(cdp.Timestamp)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput6(out *jwriter.Writer, in DispatchTouchEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	(in.Type).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"touchPoints\":")
	if in.TouchPoints == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.TouchPoints {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	if in.Modifiers != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modifiers\":")
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timestamp\":")
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			(*in.Timestamp).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchTouchEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchTouchEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchTouchEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchTouchEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput6(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput7(in *jlexer.Lexer, out *DispatchMouseEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "x":
			out.X = int64(in.Int64())
		case "y":
			out.Y = int64(in.Int64())
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(cdp.Timestamp)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "button":
			(out.Button).UnmarshalEasyJSON(in)
		case "clickCount":
			out.ClickCount = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput7(out *jwriter.Writer, in DispatchMouseEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	(in.Type).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"x\":")
	out.Int64(int64(in.X))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"y\":")
	out.Int64(int64(in.Y))
	if in.Modifiers != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modifiers\":")
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timestamp\":")
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			(*in.Timestamp).MarshalEasyJSON(out)
		}
	}
	if in.Button != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"button\":")
		(in.Button).MarshalEasyJSON(out)
	}
	if in.ClickCount != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"clickCount\":")
		out.Int64(int64(in.ClickCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchMouseEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchMouseEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchMouseEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchMouseEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput7(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput8(in *jlexer.Lexer, out *DispatchKeyEventParams) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			(out.Type).UnmarshalEasyJSON(in)
		case "modifiers":
			(out.Modifiers).UnmarshalEasyJSON(in)
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(cdp.Timestamp)
				}
				(*out.Timestamp).UnmarshalEasyJSON(in)
			}
		case "text":
			out.Text = string(in.String())
		case "unmodifiedText":
			out.UnmodifiedText = string(in.String())
		case "keyIdentifier":
			out.KeyIdentifier = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "key":
			out.Key = string(in.String())
		case "windowsVirtualKeyCode":
			out.WindowsVirtualKeyCode = int64(in.Int64())
		case "nativeVirtualKeyCode":
			out.NativeVirtualKeyCode = int64(in.Int64())
		case "autoRepeat":
			out.AutoRepeat = bool(in.Bool())
		case "isKeypad":
			out.IsKeypad = bool(in.Bool())
		case "isSystemKey":
			out.IsSystemKey = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput8(out *jwriter.Writer, in DispatchKeyEventParams) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"type\":")
	(in.Type).MarshalEasyJSON(out)
	if in.Modifiers != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modifiers\":")
		(in.Modifiers).MarshalEasyJSON(out)
	}
	if in.Timestamp != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"timestamp\":")
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			(*in.Timestamp).MarshalEasyJSON(out)
		}
	}
	if in.Text != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"text\":")
		out.String(string(in.Text))
	}
	if in.UnmodifiedText != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"unmodifiedText\":")
		out.String(string(in.UnmodifiedText))
	}
	if in.KeyIdentifier != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"keyIdentifier\":")
		out.String(string(in.KeyIdentifier))
	}
	if in.Code != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"code\":")
		out.String(string(in.Code))
	}
	if in.Key != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"key\":")
		out.String(string(in.Key))
	}
	if in.WindowsVirtualKeyCode != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"windowsVirtualKeyCode\":")
		out.Int64(int64(in.WindowsVirtualKeyCode))
	}
	if in.NativeVirtualKeyCode != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"nativeVirtualKeyCode\":")
		out.Int64(int64(in.NativeVirtualKeyCode))
	}
	if in.AutoRepeat {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"autoRepeat\":")
		out.Bool(bool(in.AutoRepeat))
	}
	if in.IsKeypad {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isKeypad\":")
		out.Bool(bool(in.IsKeypad))
	}
	if in.IsSystemKey {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isSystemKey\":")
		out.Bool(bool(in.IsSystemKey))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DispatchKeyEventParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DispatchKeyEventParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpInput8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DispatchKeyEventParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DispatchKeyEventParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpInput8(l, v)
}
