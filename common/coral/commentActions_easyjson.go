// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package coral

import (
	json "encoding/json"
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

func easyjsonF96a437cDecodeGithubComCoralprojectCoralImporterCommonCoral(in *jlexer.Lexer, out *CommentAction) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "tenantID":
			out.TenantID = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "siteID":
			out.SiteID = string(in.String())
		case "actionType":
			out.ActionType = string(in.String())
		case "commentID":
			out.CommentID = string(in.String())
		case "commentRevisionID":
			out.CommentRevisionID = string(in.String())
		case "reason":
			out.Reason = string(in.String())
		case "additionalDetails":
			out.AdditionalDetails = string(in.String())
		case "storyID":
			out.StoryID = string(in.String())
		case "userID":
			if in.IsNull() {
				in.Skip()
				out.UserID = nil
			} else {
				if out.UserID == nil {
					out.UserID = new(string)
				}
				*out.UserID = string(in.String())
			}
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Metadata = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Metadata)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "importedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ImportedAt).UnmarshalJSON(data))
			}
		case "extra":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Extra = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 interface{}
					if m, ok := v2.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v2.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v2 = in.Interface()
					}
					(out.Extra)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonF96a437cEncodeGithubComCoralprojectCoralImporterCommonCoral(out *jwriter.Writer, in CommentAction) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tenantID\":"
		out.RawString(prefix[1:])
		out.String(string(in.TenantID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"siteID\":"
		out.RawString(prefix)
		out.String(string(in.SiteID))
	}
	{
		const prefix string = ",\"actionType\":"
		out.RawString(prefix)
		out.String(string(in.ActionType))
	}
	{
		const prefix string = ",\"commentID\":"
		out.RawString(prefix)
		out.String(string(in.CommentID))
	}
	{
		const prefix string = ",\"commentRevisionID\":"
		out.RawString(prefix)
		out.String(string(in.CommentRevisionID))
	}
	if in.Reason != "" {
		const prefix string = ",\"reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	if in.AdditionalDetails != "" {
		const prefix string = ",\"additionalDetails\":"
		out.RawString(prefix)
		out.String(string(in.AdditionalDetails))
	}
	{
		const prefix string = ",\"storyID\":"
		out.RawString(prefix)
		out.String(string(in.StoryID))
	}
	{
		const prefix string = ",\"userID\":"
		out.RawString(prefix)
		if in.UserID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.UserID))
		}
	}
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		if in.Metadata == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.Metadata {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				if m, ok := v3Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v3Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v3Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"importedAt\":"
		out.RawString(prefix)
		out.Raw((in.ImportedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"extra\":"
		out.RawString(prefix)
		if in.Extra == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Extra {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				if m, ok := v4Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v4Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v4Value))
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentAction) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF96a437cEncodeGithubComCoralprojectCoralImporterCommonCoral(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentAction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF96a437cEncodeGithubComCoralprojectCoralImporterCommonCoral(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentAction) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF96a437cDecodeGithubComCoralprojectCoralImporterCommonCoral(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentAction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF96a437cDecodeGithubComCoralprojectCoralImporterCommonCoral(l, v)
}
