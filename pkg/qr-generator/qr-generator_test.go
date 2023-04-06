package qrgenerator

import (
	"bytes"
	"image/png"
	"testing"
)

func TestCreateQRCodeGeneratesPNG(t *testing.T) {
	buf := new(bytes.Buffer)
	_ = CreateQRCode(buf, "453526", 1)

	if len(buf.Bytes()) == 0 {
		t.Errorf("generated file has no data")
	}

	_, err := png.Decode(buf)
	if err != nil {
		t.Errorf("generated file is not a PNG: %s", err)
	}
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(bytes.Buffer)
	err := CreateQRCode(w, "", 1)

	if err == nil {
		t.Errorf("error not propagated correctly, got %v", err)
	}
}

func TestVersionDetermineSize(t *testing.T) {
	table := []struct {
		version        int
		expectedWidth  int
		expectedHeight int
	}{
		{1, 21, 21},
		{2, 25, 25},
		{6, 41, 41},
		{7, 45, 45},
		{14, 73, 73},
		{40, 177, 177},
	}

	for _, testCase := range table {
		size := Version(testCase.version).CalcSizeByVersion()

		if size != testCase.expectedWidth {
			t.Errorf("version 1 QR expected %d width. got:%d", testCase.expectedWidth, size)
		}

		if size != testCase.expectedHeight {
			t.Errorf("version 1 QR expected %d height. got:%d", testCase.expectedHeight, size)
		}
	}
}
