package merchandiser_utils

import (
	"bytes"

	merchandiser_goods "merchandiser/goods"
)

func ProductDescription(description string, attributes []merchandiser_goods.Attribute) string {
	descatts := []string{description, "\n-{}-\n"}

	for _, attattribute := range attributes {
		if attattribute.Title != "Код товара" {
			descatts = append(descatts, attattribute.Title+" - "+attattribute.Value+"\n")
		}
	}

	var buf bytes.Buffer
	for _, descatt := range descatts {
		buf.WriteString(descatt)
	}

	return buf.String()
}
