package fingerprint

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaries(t *testing.T) {
	binaries := []string{
		// pe
		"calc.exe",
		"calc.packed.exe",
		"helloworld.exe",
		"avcodec.dll",
		// macho
		"main",
		"main.fat",
		"sqlite3.fat.packed",
		// elf
		"hello.linux.stripped",
		"hello.linux.packed.stripped",
		"libxml2.so",
	}
	for _, binary := range binaries {
		t.Run(binary, func(t *testing.T) {
			f, err := os.Open("./fixtures/" + binary)
			require.NoError(t, err)
			defer f.Close()
			fileInfo, err := f.Stat()
			require.NoError(t, err)

			fixture, err := os.Open("./fixtures/" + binary + ".fingerprint")
			require.NoError(t, err)
			defer fixture.Close()
			expected, err := ioutil.ReadAll(fixture)
			require.NoError(t, err)

			info, err := Parse(f, int(fileInfo.Size()))
			require.NoError(t, err)

			data, err := json.Marshal(info)
			require.NoError(t, err)
			require.JSONEq(t, string(expected), string(data))
		})
	}
}
