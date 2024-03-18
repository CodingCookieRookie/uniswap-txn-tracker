package api

import (
	"testing"

	"github.com/CodingCookieRookie/uniswap-txn-tracker/api/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHistoricalTxnTimes(t *testing.T) {
	tests := []struct {
		name      string
		startTime string
		endTime   string
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "valid times",
			startTime: "2023-01-01 00:00:00",
			endTime:   "2023-01-02 00:00:00",
			wantErr:   false,
		},
		{
			name:      "invalid start time format",
			startTime: "01-01-2023 00:00:00",
			endTime:   "2023-01-02 00:00:00",
			wantErr:   true,
			errMsg:    "start time must be in correct format, eg. 2006-01-02 15:04:05",
		},
		{
			name:      "invalid end time format",
			startTime: "2023-01-01 00:00:00",
			endTime:   "01-02-2023",
			wantErr:   true,
			errMsg:    "end time must be in correct format, eg. 2006-01-02 15:04:05",
		},
		{
			name:      "start time after end time",
			startTime: "2023-01-02 00:00:00",
			endTime:   "2023-01-01 00:00:00",
			wantErr:   true,
			errMsg:    "start time must be before or equal to end time",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.HistoricalTxnReq{
				StartTime: tt.startTime,
				EndTime:   tt.endTime,
			}
			_, _, err := parseHistoricalTxnTimes(req)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			}
		})
	}
}
