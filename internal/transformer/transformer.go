package transformer

type SbayIdTransformer struct {
	encoder *UrlEncoder
}

func NewSbayIdTransformer(alphabet string) *SbayIdTransformer {
	return &SbayIdTransformer{
		encoder: NewURLEncoder(&Options{
			Alphabet: alphabet,
		}),
	}
}

func (s *SbayIdTransformer) String2ID(value string) uint64 {
	// i, err := strconv.Atoi(value)
	// if err == nil {
	// 	return uint64(i)
	// }
	return s.encoder.DecodeURL(value)
}

func (s *SbayIdTransformer) ID2String(value uint64) string {
	return s.encoder.EncodeURL(value)
}
