

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
		offset := cursor+1+int(encoded[cursor])
		bytArr := encoded[cursor+1:offset]
		words = append(words, bytArr)
		cursor = offset-1
	}

	return words
}
