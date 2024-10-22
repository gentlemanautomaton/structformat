package structformat_test

import (
	"fmt"
	"strings"
	"testing"
)

const expectedBuilderTestResult = `
00:  N.LNE.NI.S.NI.S: (C0, E0)
01:  NE.SE.LNE.S.N.N: D1 (E1, F1)
02:     LN.P.S.P.S.S: (Label: A2) B2: C2 D2: E2 F2
03:    N.P.P.LNI.N.P: (Label: D3)
04:   N.LNI.P.N.P.LN: (Label: B4)
05: SI.SI.S.NI.LNI.S: A5 B5 (D5, Label: E5)
06:    N.P.S.N.LNE.S: (A6) B6: C6 (D6) F6
07:    SI.N.N.N.SE.S: A7
08:     N.S.P.S.N.PE: (A8) B8 C8: D8 (E8)
09:    LN.P.LN.P.N.P: (Label: A9) B9: (Label: C9) D9: (E9) F9
10:    S.S.P.PI.SI.S: D10: E10
11:    LN.S.PE.N.N.N: (Label: A11) B11 (D11, E11, F11)
12:   LN.S.N.N.SE.NI: (F12)
13:  SI.N.SI.LNI.S.S: A13 C13 (Label: D13)
14:   N.PE.S.S.LN.NI: (F14)
15:    N.P.SI.SI.S.S: C15 D15
16:   N.NI.SE.S.S.NE: (B16)
17:  LN.S.N.NI.N.LNE: (D17)
18:    LN.SE.S.S.P.N: (Label: A18) C18 D18 E18: (F18)
19:   S.PI.S.S.NI.LN: B19: (E19)
20:   PE.S.SI.SE.N.P: C20
21:    P.NI.N.NI.P.N: (B21, D21)
22:     S.S.S.N.S.LN: A22 B22 C22 (D22) E22 (Label: F22)
23:     S.N.S.S.SE.S: A23 (B23) C23 D23 F23
24:  LN.S.S.PI.PI.PI: D24: E24: F24
25:   NI.S.N.P.LNE.N: (A25)
26:   N.LNI.LN.S.P.N: (Label: B26)
27:  P.PI.S.NE.LNI.S: B27: (Label: E27)
28:    PI.S.SI.P.S.P: A28: C28
29:    S.P.N.SI.S.NE: D29
30:    S.P.S.SI.N.PI: D30 F30
31:   S.SE.PE.NE.S.N: A31 E31 (F31)
32:     N.S.S.S.PI.N: E32
33:     P.SE.S.S.P.N: A33: C33 D33 E33: (F33)
34:   N.SE.N.SI.LN.S: D34
35:     N.P.SI.P.S.S: C35
36:   PI.N.NI.S.S.SI: A36: (C36) F36
37:   S.P.SI.PE.SI.S: C37 E37
38:     S.S.N.PE.S.N: A38 B38 (C38) E38 (F38)
39:   S.S.SI.SE.S.NI: C39 (F39)
40:    SI.NI.N.N.S.N: A40 (B40)
41:    N.LN.N.P.NE.S: (A41, Label: B41, C41) D41: F41
42:    PE.NI.S.S.S.N: (B42)
43:     N.P.SI.S.N.P: C43
44:     N.S.S.S.S.NI: (F44)
45:  S.PE.N.SI.NI.LN: D45 (E45)
46:    P.N.N.LN.SE.P: A46: (B46, C46, Label: D46) F46
47:    N.SI.P.N.N.LN: B47
48:    S.S.NI.P.N.NI: (C48, F48)
49:     P.S.S.SI.S.P: D49
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
		results = append(results, fmt.Sprintf("%02d: %16s: %s", i, format, result))
	}
	output := strings.Join(results, "\n")
	if output != strings.Trim(expectedBuilderTestResult, "\n") {
		t.Errorf("Unexpected Results:\n%s", output)
	}
}
