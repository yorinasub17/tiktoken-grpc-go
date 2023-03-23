package tiktokengrpc

// TiktokenEncodingName is an enum for valid encoding names accepted by the tiktoken service.
type TiktokenEncodingName string

const (
	CL100kBase   TiktokenEncodingName = "cl100k_base"
	P50kBase     TiktokenEncodingName = "p50k_base"
	R50kBase     TiktokenEncodingName = "r50k_base"
	GPT2Encoding TiktokenEncodingName = "gpt2"
)

// IsValid returns whether the given tiktoken encoding name fits one of the enum values, since go enums can be casted
// with bogus values easily.
func (n TiktokenEncodingName) IsValid() bool {
	switch n {
	case CL100kBase, P50kBase, R50kBase, GPT2Encoding:
		return true
	}
	return false
}
