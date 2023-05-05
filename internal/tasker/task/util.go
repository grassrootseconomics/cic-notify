package task

import (
	"fmt"
	"math"
	"strings"

	"github.com/golang-module/carbon/v2"
	"github.com/grassrootseconomics/w3-celo-patch"
)

const (
	chainTokenDecimals = 6
)

var (
	funcBalanceOf = w3.MustNewFunc("balanceOf(address)", "uint256")
)

// formatDate takes a unix timestamp and timezone and returns a formatted string
// 1649735755981 + Europe/Moscow
// 2021-01-27 13:14:15
func formatDate(unixTimestamp uint64, timeZone string) string {
	return carbon.CreateFromTimestamp(int64(unixTimestamp)).SetTimezone(timeZone).Format("Y-m-d H:i")
}

// formatIdentifier takes the first name and last name and returns a formatted name string.
func formatIdentifier(firstName string, lastName string, identifier string, blockchainAddress string) string {
	if firstName != "" || lastName != "" || identifier != "" {
		return strings.ToUpper(strings.TrimSpace(fmt.Sprintf("%s %s %s", firstName, lastName, identifier)))
	} else {
		return blockchainAddress
	}
}

// formatShortHash takes a full txHash and returns the last 8 chars of the Ethereum transaction hash (hex).
// 0x1562767d2a01098da599cdea23ff798838a530a17e6072838c425d48.
// Should return 7837424A.
func formatShortHash(txHash string) string {
	return strings.ToUpper(txHash[58:])
}

// truncateVoucherValue takes an input and formats it to a human readable value.
// We truncate instead of rounding off extra decimal places realised from the computation.
// 6219000 -> 6.21
func truncateVoucherValue(inputValue uint64) string {
	return fmt.Sprintf("%.2f", float64(int(float64(inputValue)/math.Pow10(chainTokenDecimals)*100))/100)
}
