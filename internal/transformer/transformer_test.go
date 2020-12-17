package transformer

import "testing"

func TestSbayTransformer(t *testing.T) {
	transfomer := NewSbayIdTransformer("mn6j2c4rv8bpygw95z7hsdaetxuk3fq")
	resultString := transfomer.ID2String(0)
	if resultString != "mmmmm" {
		t.Errorf("ID2String is error, got %s, want %s", resultString, "mmmmm")
	}

	resultID := transfomer.String2ID("mmmmm")
	if resultID != 0 {
		t.Errorf("String2ID is error, got: %d, want: %d", resultID, 0)
	}
}
