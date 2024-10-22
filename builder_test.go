package structformat_test

import (
	"fmt"
	"strings"
	"testing"
)

const expectedBuilderTestResult = `
00:    N.S.SI.NI.N.GSE: C0 (D0)
01:      SE.N.N.S.LN.P: (B1, C1) D1 (Label: E1) F1
02:      S.P.P.LNI.P.N: (Label: D2)
03:     LNI.N.P.S.SI.S: (Label: A3) E3
04:      PI.SI.S.P.S.N: A4: B4
05:    SI.SI.N.P.LNE.P: A5 B5
06:      LN.PI.N.S.N.S: B6
07:      N.P.LN.P.SI.S: E7
08:     LN.S.S.S.P.GPE: (Label: A8) B8 C8 D8 E8
09:  N.NI.GNI.S.LNI.SI: (B9): (C9): (Label: E9) F9
10:     LN.P.SI.S.S.PI: C10 F10
11:    S.PE.SI.SE.N.PI: C11 F11
12:      S.S.LN.LN.N.S: A12 B12 (Label: C12, Label: D12, E12) F12
13:      P.S.PI.GS.S.N: C13
14: NI.GLN.LN.GNE.NI.S: (A14, E14)
15:      N.N.N.LN.N.LN: (A15, B15, C15, Label: D15, E15, Label: F15)
16:       N.S.S.P.LN.P: (A16) B16 C16 D16: (Label: E16) F16
17:   GS.S.GN.S.PI.LNE: E17
18:      N.NI.GN.N.S.N: (B18)
19:    P.N.SE.NE.SI.NI: E19 (F19)
20:      PI.P.N.N.NE.S: A20
21:     PE.N.LNI.N.S.S: (Label: C21)
22:    GNE.GN.N.PE.S.S: (B22): (C22) E22 F22
23:     S.SE.GLN.N.N.S: A23: (Label: C23): (D23, E23) F23
24:       NI.S.N.P.S.S: (A24)
25:      NI.N.S.S.SE.P: (A25)
26:    S.GSI.GN.N.PE.S: B26
27:    SI.S.S.GS.LN.NI: A27 (F27)
28:      N.LN.N.NI.N.S: (D28)
29:    S.PE.LNI.S.GN.N: (Label: C29)
30:    SI.S.P.P.LN.GLN: A30
31:      N.S.S.PI.NI.S: D31: (E31)
32:       P.N.LN.S.N.S: A32: (B32, Label: C32) D32 (E32) F32
33:     N.P.S.GLN.PI.N: E33
34:    NI.PI.S.GSI.P.N: (A34) B34: D34
35:       S.N.N.GN.S.P: A35 (B35, C35): (D35): E35 F35
36:    NE.NE.GS.P.N.GP: C36: D36: (E36): F36
37:       S.S.NI.S.S.S: (C37)
38:       NI.S.S.S.N.P: (A38)
39:       P.N.LN.P.N.S: A39: (B39, Label: C39) D39: (E39) F39
40:    LN.P.NI.SE.SI.S: (C40) E40
41:    S.S.NI.LN.S.GSE: (C41)
42:    GS.SE.S.GN.SI.S: E42
43:      S.P.S.SE.SI.N: E43
44:      S.N.LNI.P.N.S: (Label: C44)
45:     S.S.GN.S.SI.LN: E45
46:       S.GS.N.S.P.P: A46: B46: (C46) D46 E46: F46
47:       N.N.S.N.P.NE: (A47, B47) C47 (D47) E47
48:       P.LN.S.S.S.N: A48: (Label: B48) C48 D48 E48 (F48)
49:    N.S.N.LN.LNI.PI: (Label: E49) F49
`

func TestBuilder(t *testing.T) {
	const records = 50
	const seed1, seed2 = 1, 2
	formats := newFormatGenerator(seed1, seed2)
	var results []string
	for i := 0; i < records; i++ {
		record := makeRecord(i)
		format := formats.Generate()
		result := record.Format(format)
		results = append(results, fmt.Sprintf("%02d: %18s: %s", i, format, result))
	}
	output := strings.Join(results, "\n")
	if output != strings.Trim(expectedBuilderTestResult, "\n") {
		t.Errorf("Unexpected Results:\n%s", output)
	}
}
