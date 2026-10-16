

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var buffer bytes.Buffer

	for _, word := range strs{
		buffer.WriteByte(byte(len(word)))
		buffer.WriteString(word)
	}

	return buffer.String()
}

func (s *Solution) Decode(encoded string) []string {
	words := []string{}

	for cursor := 0 ; cursor<len(encoded); cursor++{
		var buffer bytes.Buffer
		offset := cursor+1+int(encoded[cursor])
		buffer.WriteString(encoded[cursor+1:offset])
		words = append(words, buffer.String())
		cursor = offset-1
	}

	return words
}
