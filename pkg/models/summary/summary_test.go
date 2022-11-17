package summary

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromRaw(t *testing.T) {
	cases := map[string]struct {
		pld      []byte
		expected *Summary
		err      func(*testing.T, error)
	}{
		"rest retrieve summary": {
			pld: []byte(`[
				null,
				null,
				null,
				null,
				[
					[0.001,0.001,0.001,null,null,-0.0002],
					[0.002,0.003,0.004,null,null,0.00075]
				],
				null,
				null,
				null,
				null,
				{leo_lev:0,leo_amount_avg:0}
			]`),
			expected: &Summary{
				MakerFee:         0.001,
				DerivRebate:      -0.0002,
				TakerFeeToCrypto: 0.002,
				TakerFeeToStable: 0.003,
				TakerFeeToFiat:   0.004,
				DerivTakerFee:    0.00075,
			},
			err: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			re := regexp.MustCompile(`,([\n\s]+)?{.+}`)
			str := re.ReplaceAllString(string(v.pld), "")
			var raw []interface{}
			err := json.Unmarshal([]byte(str), &raw)
			require.Nil(t, err)

			got, err := FromRaw(raw)
			v.err(t, err)
			assert.Equal(t, v.expected, got)
		})
	}
}
