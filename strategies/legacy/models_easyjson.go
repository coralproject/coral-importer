// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package legacy

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	coral "gitlab.com/coralproject/coral-importer/common/coral"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy(in *jlexer.Lexer, out *UserToken) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "active":
			out.Active = bool(in.Bool())
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy(out *jwriter.Writer, in UserToken) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"active\":"
		out.RawString(prefix)
		out.Bool(bool(in.Active))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserToken) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserToken) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserToken) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserToken) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(in *jlexer.Lexer, out *UserProfile) {
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
		case "id":
			out.ID = string(in.String())
		case "provider":
			out.Provider = string(in.String())
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(out *jwriter.Writer, in UserProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"provider\":"
		out.RawString(prefix)
		out.String(string(in.Provider))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy1(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(in *jlexer.Lexer, out *UserNotifications) {
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
		case "settings":
			if in.IsNull() {
				in.Skip()
				out.Settings = nil
			} else {
				if out.Settings == nil {
					out.Settings = new(UserNotificationSettings)
				}
				(*out.Settings).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(out *jwriter.Writer, in UserNotifications) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"settings\":"
		out.RawString(prefix[1:])
		if in.Settings == nil {
			out.RawString("null")
		} else {
			(*in.Settings).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserNotifications) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserNotifications) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserNotifications) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserNotifications) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy2(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(in *jlexer.Lexer, out *UserNotificationSettings) {
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
		case "onFeatured":
			if in.IsNull() {
				in.Skip()
				out.OnFeatured = nil
			} else {
				if out.OnFeatured == nil {
					out.OnFeatured = new(bool)
				}
				*out.OnFeatured = bool(in.Bool())
			}
		case "onModeration":
			if in.IsNull() {
				in.Skip()
				out.OnModeration = nil
			} else {
				if out.OnModeration == nil {
					out.OnModeration = new(bool)
				}
				*out.OnModeration = bool(in.Bool())
			}
		case "onReply":
			if in.IsNull() {
				in.Skip()
				out.OnReply = nil
			} else {
				if out.OnReply == nil {
					out.OnReply = new(bool)
				}
				*out.OnReply = bool(in.Bool())
			}
		case "onStaffReply":
			if in.IsNull() {
				in.Skip()
				out.OnStaffReply = nil
			} else {
				if out.OnStaffReply == nil {
					out.OnStaffReply = new(bool)
				}
				*out.OnStaffReply = bool(in.Bool())
			}
		case "digestFrequency":
			if m, ok := out.DigestFrequency.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.DigestFrequency.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.DigestFrequency = in.Interface()
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(out *jwriter.Writer, in UserNotificationSettings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"onFeatured\":"
		out.RawString(prefix[1:])
		if in.OnFeatured == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.OnFeatured))
		}
	}
	{
		const prefix string = ",\"onModeration\":"
		out.RawString(prefix)
		if in.OnModeration == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.OnModeration))
		}
	}
	{
		const prefix string = ",\"onReply\":"
		out.RawString(prefix)
		if in.OnReply == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.OnReply))
		}
	}
	{
		const prefix string = ",\"onStaffReply\":"
		out.RawString(prefix)
		if in.OnStaffReply == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.OnStaffReply))
		}
	}
	{
		const prefix string = ",\"digestFrequency\":"
		out.RawString(prefix)
		if m, ok := in.DigestFrequency.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.DigestFrequency.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.DigestFrequency))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserNotificationSettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserNotificationSettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserNotificationSettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserNotificationSettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy3(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(in *jlexer.Lexer, out *UserMetadata) {
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
		case "notifications":
			if in.IsNull() {
				in.Skip()
				out.Notifications = nil
			} else {
				if out.Notifications == nil {
					out.Notifications = new(UserNotifications)
				}
				(*out.Notifications).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(out *jwriter.Writer, in UserMetadata) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"notifications\":"
		out.RawString(prefix[1:])
		if in.Notifications == nil {
			out.RawString("null")
		} else {
			(*in.Notifications).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserMetadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserMetadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserMetadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserMetadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy4(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "role":
			out.Role = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "ignoresUsers":
			if in.IsNull() {
				in.Skip()
				out.IgnoredUsers = nil
			} else {
				in.Delim('[')
				if out.IgnoredUsers == nil {
					if !in.IsDelim(']') {
						out.IgnoredUsers = make([]string, 0, 4)
					} else {
						out.IgnoredUsers = []string{}
					}
				} else {
					out.IgnoredUsers = (out.IgnoredUsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.IgnoredUsers = append(out.IgnoredUsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "profiles":
			if in.IsNull() {
				in.Skip()
				out.Profiles = nil
			} else {
				in.Delim('[')
				if out.Profiles == nil {
					if !in.IsDelim(']') {
						out.Profiles = make([]UserProfile, 0, 2)
					} else {
						out.Profiles = []UserProfile{}
					}
				} else {
					out.Profiles = (out.Profiles)[:0]
				}
				for !in.IsDelim(']') {
					var v2 UserProfile
					(v2).UnmarshalEasyJSON(in)
					out.Profiles = append(out.Profiles, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tokens":
			if in.IsNull() {
				in.Skip()
				out.Tokens = nil
			} else {
				in.Delim('[')
				if out.Tokens == nil {
					if !in.IsDelim(']') {
						out.Tokens = make([]UserToken, 0, 1)
					} else {
						out.Tokens = []UserToken{}
					}
				} else {
					out.Tokens = (out.Tokens)[:0]
				}
				for !in.IsDelim(']') {
					var v3 UserToken
					(v3).UnmarshalEasyJSON(in)
					out.Tokens = append(out.Tokens, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(UserMetadata)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"role\":"
		out.RawString(prefix)
		out.String(string(in.Role))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"ignoresUsers\":"
		out.RawString(prefix)
		if in.IgnoredUsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.IgnoredUsers {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"profiles\":"
		out.RawString(prefix)
		if in.Profiles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Profiles {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tokens\":"
		out.RawString(prefix)
		if in.Tokens == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Tokens {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil {
			out.RawString("null")
		} else {
			(*in.Metadata).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy5(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(in *jlexer.Lexer, out *CommentTag) {
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
		case "assigned_by":
			if in.IsNull() {
				in.Skip()
				out.AssignedBy = nil
			} else {
				if out.AssignedBy == nil {
					out.AssignedBy = new(string)
				}
				*out.AssignedBy = string(in.String())
			}
		case "tag":
			easyjsonD2b7633eDecode(in, &out.Tag)
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(out *jwriter.Writer, in CommentTag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"assigned_by\":"
		out.RawString(prefix[1:])
		if in.AssignedBy == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.AssignedBy))
		}
	}
	{
		const prefix string = ",\"tag\":"
		out.RawString(prefix)
		easyjsonD2b7633eEncode(out, in.Tag)
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentTag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentTag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentTag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentTag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy6(l, v)
}
func easyjsonD2b7633eDecode(in *jlexer.Lexer, out *struct {
	Name string `json:"name"`
}) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjsonD2b7633eEncode(out *jwriter.Writer, in struct {
	Name string `json:"name"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(in *jlexer.Lexer, out *CommentBodyHistory) {
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
		case "body":
			out.Body = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(out *jwriter.Writer, in CommentBodyHistory) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix[1:])
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentBodyHistory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentBodyHistory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentBodyHistory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentBodyHistory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy7(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(in *jlexer.Lexer, out *Comment) {
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
		case "id":
			out.ID = string(in.String())
		case "asset_id":
			out.AssetID = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "body_history":
			if in.IsNull() {
				in.Skip()
				out.BodyHistory = nil
			} else {
				in.Delim('[')
				if out.BodyHistory == nil {
					if !in.IsDelim(']') {
						out.BodyHistory = make([]CommentBodyHistory, 0, 1)
					} else {
						out.BodyHistory = []CommentBodyHistory{}
					}
				} else {
					out.BodyHistory = (out.BodyHistory)[:0]
				}
				for !in.IsDelim(']') {
					var v10 CommentBodyHistory
					(v10).UnmarshalEasyJSON(in)
					out.BodyHistory = append(out.BodyHistory, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]CommentTag, 0, 1)
					} else {
						out.Tags = []CommentTag{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v11 CommentTag
					(v11).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "parent_id":
			if in.IsNull() {
				in.Skip()
				out.ParentID = nil
			} else {
				if out.ParentID == nil {
					out.ParentID = new(string)
				}
				*out.ParentID = string(in.String())
			}
		case "author_id":
			out.AuthorID = string(in.String())
		case "deleted_at":
			if in.IsNull() {
				in.Skip()
				out.DeletedAt = nil
			} else {
				if out.DeletedAt == nil {
					out.DeletedAt = new(coral.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DeletedAt).UnmarshalJSON(data))
				}
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(out *jwriter.Writer, in Comment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"asset_id\":"
		out.RawString(prefix)
		out.String(string(in.AssetID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"body_history\":"
		out.RawString(prefix)
		if in.BodyHistory == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.BodyHistory {
				if v12 > 0 {
					out.RawByte(',')
				}
				(v13).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Tags {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"parent_id\":"
		out.RawString(prefix)
		if in.ParentID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ParentID))
		}
	}
	{
		const prefix string = ",\"author_id\":"
		out.RawString(prefix)
		out.String(string(in.AuthorID))
	}
	{
		const prefix string = ",\"deleted_at\":"
		out.RawString(prefix)
		if in.DeletedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.DeletedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Comment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Comment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Comment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Comment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy8(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(in *jlexer.Lexer, out *Asset) {
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
		case "id":
			out.ID = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "closedAt":
			if in.IsNull() {
				in.Skip()
				out.ClosedAt = nil
			} else {
				if out.ClosedAt == nil {
					out.ClosedAt = new(coral.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ClosedAt).UnmarshalJSON(data))
				}
			}
		case "closedMessage":
			if in.IsNull() {
				in.Skip()
				out.ClosedMessage = nil
			} else {
				if out.ClosedMessage == nil {
					out.ClosedMessage = new(string)
				}
				*out.ClosedMessage = string(in.String())
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "scraped":
			if in.IsNull() {
				in.Skip()
				out.Scraped = nil
			} else {
				if out.Scraped == nil {
					out.Scraped = new(coral.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Scraped).UnmarshalJSON(data))
				}
			}
		case "metadata":
			if m, ok := out.Metadata.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Metadata.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Metadata = in.Interface()
			}
		case "settings":
			if m, ok := out.Settings.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Settings.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Settings = in.Interface()
			}
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(string)
				}
				*out.Author = string(in.String())
			}
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
			}
		case "image":
			if in.IsNull() {
				in.Skip()
				out.Image = nil
			} else {
				if out.Image == nil {
					out.Image = new(string)
				}
				*out.Image = string(in.String())
			}
		case "section":
			if in.IsNull() {
				in.Skip()
				out.Section = nil
			} else {
				if out.Section == nil {
					out.Section = new(string)
				}
				*out.Section = string(in.String())
			}
		case "modified_date":
			if in.IsNull() {
				in.Skip()
				out.ModifiedDate = nil
			} else {
				if out.ModifiedDate == nil {
					out.ModifiedDate = new(coral.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ModifiedDate).UnmarshalJSON(data))
				}
			}
		case "publication_date":
			if in.IsNull() {
				in.Skip()
				out.PublicationDate = nil
			} else {
				if out.PublicationDate == nil {
					out.PublicationDate = new(coral.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.PublicationDate).UnmarshalJSON(data))
				}
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(out *jwriter.Writer, in Asset) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"closedAt\":"
		out.RawString(prefix)
		if in.ClosedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.ClosedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"closedMessage\":"
		out.RawString(prefix)
		if in.ClosedMessage == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ClosedMessage))
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"scraped\":"
		out.RawString(prefix)
		if in.Scraped == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Scraped).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if m, ok := in.Metadata.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Metadata.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Metadata))
		}
	}
	{
		const prefix string = ",\"settings\":"
		out.RawString(prefix)
		if m, ok := in.Settings.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Settings.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Settings))
		}
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		if in.Title == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Title))
		}
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		if in.Author == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Author))
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		if in.Image == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Image))
		}
	}
	{
		const prefix string = ",\"section\":"
		out.RawString(prefix)
		if in.Section == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Section))
		}
	}
	{
		const prefix string = ",\"modified_date\":"
		out.RawString(prefix)
		if in.ModifiedDate == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.ModifiedDate).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"publication_date\":"
		out.RawString(prefix)
		if in.PublicationDate == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.PublicationDate).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Asset) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Asset) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Asset) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Asset) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy9(l, v)
}
func easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(in *jlexer.Lexer, out *Action) {
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
		case "id":
			out.ID = string(in.String())
		case "action_type":
			out.ActionType = string(in.String())
		case "group_id":
			out.GroupID = string(in.String())
		case "item_id":
			out.ItemID = string(in.String())
		case "item_type":
			out.ItemType = string(in.String())
		case "user_id":
			if in.IsNull() {
				in.Skip()
				out.UserID = nil
			} else {
				if out.UserID == nil {
					out.UserID = new(string)
				}
				*out.UserID = string(in.String())
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]interface{})
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v16 interface{}
					if m, ok := v16.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v16.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v16 = in.Interface()
					}
					(out.Metadata)[key] = v16
					in.WantComma()
				}
				in.Delim('}')
			}
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
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
func easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(out *jwriter.Writer, in Action) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"action_type\":"
		out.RawString(prefix)
		out.String(string(in.ActionType))
	}
	{
		const prefix string = ",\"group_id\":"
		out.RawString(prefix)
		out.String(string(in.GroupID))
	}
	{
		const prefix string = ",\"item_id\":"
		out.RawString(prefix)
		out.String(string(in.ItemID))
	}
	{
		const prefix string = ",\"item_type\":"
		out.RawString(prefix)
		out.String(string(in.ItemType))
	}
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix)
		if in.UserID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.UserID))
		}
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v17First := true
			for v17Name, v17Value := range in.Metadata {
				if v17First {
					v17First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v17Name))
				out.RawByte(':')
				if m, ok := v17Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v17Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v17Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Action) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Action) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Action) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Action) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGitlabComCoralprojectCoralImporterStrategiesLegacy10(l, v)
}
