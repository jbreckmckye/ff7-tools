package ff7

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const (
	// is the length of a .X file (executable module) header
	moduleHeaderLen int64 = 8
)

func ExtractX(log Logger, src string, dst string) error {
  var err error
	
	stats, err := os.Stat(src)
	if err != nil {
		return err
	}

	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

  input.Seek(moduleHeaderLen, io.SeekStart)
	bodyLen := stats.Size() - moduleHeaderLen

	log(
		"# Extracting data",
		fmt.Sprintf("- Compressed body: `%d bytes`", bodyLen),
		fmt.Sprintf("- Decompressing to `%s`", dst),
	)

	unzip, err := gzip.NewReader(input)
	if err != nil {
		return err
	}
	defer unzip.Close()

	output, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer output.Close()

	written, err := io.Copy(output, unzip)
	if err != nil {
		return err
	}

	log(
		"# Complete!",
		fmt.Sprintf("Wrote decompressed object code - size `%d bytes`", written),
		"",
		"The object code is plain MIPS machine code. To disassemble it (turn back into asm), use a tool like `objdump`",
	)

	return nil
}
