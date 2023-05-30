// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package manifest_digest

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

func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(in *jlexer.Lexer, out *OciManifestResponse) {
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
		case "digest":
			out.Digest = string(in.String())
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
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(out *jwriter.Writer, in OciManifestResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"digest\":"
		out.RawString(prefix[1:])
		out.String(string(in.Digest))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OciManifestResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OciManifestResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OciManifestResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OciManifestResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesOciManifestDigest(l, v)
}