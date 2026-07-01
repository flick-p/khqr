# khqr

A Go SDK for generating and decoding **KHQR** codes — the EMVCo-based QR payment
standard used by Cambodia's National Bank of Cambodia (Bakong) for interoperable
QR payments.

It supports:

- Generating individual (P2P) and merchant KHQR strings, static or dynamic
  (amount-bearing), in USD or KHR.
- Decoding a KHQR string back into its structured fields.
- Validating a decoded KHQR: CRC integrity, required fields, and dynamic QR
  expiry.

## Installation

The module is currently developed under the import path declared in `go.mod`
(`module khqr`). If you consume this repository from another Go module, either
add a `replace` directive pointing at this repository, or update the `module`
directive in `go.mod` to your intended public import path (e.g.
`github.com/flick-p/khqr`) before publishing.

```
go get github.com/flick-p/khqr
```

## Quick start

### Generate an individual (P2P) static QR

A static QR has no fixed amount — the payer enters the amount at scan time.

```go
package main

import (
	"fmt"
	"log"

	"khqr"
	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

func main() {
	result, err := khqr.GenerateIndividual(models.IndividualInfo{
		BakongAccountID:    "receivekhqr@dvpy",
		Currency:           constants.CurrencyUSD,
		MerchantName:       "John Doe",
		MerchantCity:       "Phnom Penh",
		AccountInformation: util.Ptr("012345678"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("QR:", result.QR)
	fmt.Println("MD5:", result.MD5)
}
```

### Generate a dynamic QR with a fixed amount

Setting `Amount` produces a dynamic QR. Add `ExpirationTimestamp` (Unix millis)
to make it expire.

```go
result, err := khqr.GenerateIndividual(models.IndividualInfo{
	BakongAccountID:     "receivekhqr@dvpy",
	Currency:            constants.CurrencyUSD,
	MerchantName:        "John Doe",
	MerchantCity:        "Phnom Penh",
	Amount:              util.Ptr("10.50"),
	ExpirationTimestamp: util.Ptr(time.Now().Add(time.Hour).UnixMilli()),
})
```

### Generate a merchant QR

Merchant QRs require a `MerchantID` and an `AcquiringBank`.

```go
result, err := khqr.GenerateMerchant(models.MerchantInfo{
	MerchantID: "M123456789",
	IndividualInfo: models.IndividualInfo{
		BakongAccountID: "merchant@dvpy",
		Currency:        constants.CurrencyUSD,
		MerchantName:    "Merchant Co",
		MerchantCity:    "Phnom Penh",
		AcquiringBank:   util.Ptr("Acme Bank"),
	},
})
```

### Decode a KHQR string

```go
decoded := khqr.DecodeKHQR(qrString)
fmt.Println(decoded.MerchantName, decoded.BakongAccountID)
```

`DecodeKHQR` always returns a `models.DecodedKHQR`, even for malformed input —
fields it could not parse are left at their zero value.

### Decode and validate a KHQR string

Use this when scanning an untrusted or externally-sourced QR. It checks the
CRC16 checksum, required fields, and — for dynamic QRs — that an amount is
present and the code has not expired.

```go
decoded, err := khqr.DecodeKHQRValidation(qrString)
if err != nil {
	// qrString was tampered with, incomplete, or expired.
	log.Fatal(err)
}
```

## API reference

### `khqr` package

| Function | Description |
|---|---|
| `GenerateIndividual(data models.IndividualInfo) (*models.KHQRData, error)` | Builds a static or dynamic individual (P2P) KHQR string. |
| `GenerateMerchant(data models.MerchantInfo) (*models.KHQRData, error)` | Builds a static or dynamic merchant KHQR string. Requires `MerchantID` and `AcquiringBank`. |
| `DecodeKHQR(khqrString string) models.DecodedKHQR` | Parses a KHQR string into its component fields. Never returns an error. |
| `DecodeKHQRValidation(khqrString string) (*models.DecodedKHQR, error)` | Decodes and validates CRC, required fields, and dynamic-QR expiry. |

`models.KHQRData` contains the generated `QR` string and its `MD5` hash.

### Key `models.IndividualInfo` fields

| Field | Type | Notes |
|---|---|---|
| `BakongAccountID` | `string` | Required. Must contain `@` (e.g. `name@bank`). |
| `Currency` | `int` | Required. `constants.CurrencyUSD` (840) or `constants.CurrencyKHR` (116). |
| `MerchantName` | `string` | Required. Max 25 characters. |
| `MerchantCity` | `string` | Defaults to `"Phnom Penh"` if empty. |
| `Amount` | `*string` | Omit or leave nil/`"0"` for a static QR. USD amounts are formatted to 2 decimals; KHR amounts must be whole numbers. |
| `AccountInformation` | `*string` | Individual account reference (max 32 chars). |
| `AcquiringBank` | `*string` | Required when generating a merchant QR via `GenerateMerchant`. |
| `MerchantCategoryCode` | `*string` | Defaults to `"5999"`. Must be numeric, 0–9999. |
| `ExpirationTimestamp` | `*int64` | Unix millis. Required for the QR to be treated as expirable; must be in the future. |
| `BillNumber`, `StoreLabel`, `TerminalLabel`, `MobileNumber`, `PurposeOfTransaction` | `*string` | Optional additional data fields, each max 25 chars. |
| `LanguagePreference`, `MerchantNameAlternateLanguage`, `MerchantCityAlternateLanguage` | `*string` | Optional alternate-language fields. If `LanguagePreference` is set, `MerchantNameAlternateLanguage` is required. |

`models.MerchantInfo` embeds `IndividualInfo` and adds a required `MerchantID string`.

## Error handling

Generation errors are returned as `*constants.ErrorCode`, which implements
`error` and carries a stable numeric `Code` alongside a human-readable
`Message`. See `constants/khqr_error_code.go` for the full list (e.g.
`ErrBakongAccountIDInvalid`, `ErrUnsupportedCurrency`, `ErrKHQRExpired`).

```go
result, err := khqr.GenerateIndividual(data)
if err != nil {
	if code, ok := err.(*constants.ErrorCode); ok {
		fmt.Println(code.Code, code.Message)
	}
}
```

`DecodeKHQRValidation` wraps these in a plain `error` built from the same
messages via `errors.New`.

## Project layout

| Package | Responsibility |
|---|---|
| `khqr` | Public entry points: generate, decode, decode+validate. |
| `generator` | Builds each EMVCo tag-length-value field and assembles/validates the final QR string. |
| `decoder` | Parses a raw KHQR string back into `models.DecodedKHQR`. |
| `models` | Shared data types (`IndividualInfo`, `MerchantInfo`, `DecodedKHQR`, `KHQRData`) and the TLV helper. |
| `constants` | EMVCo tags, field lengths, currency codes, and error definitions. |
| `util` | CRC16, MD5, and TLV string-cutting helpers used by both the generator and decoder. |

## Testing

```
go test ./...
```

Tests cover the public generate/decode/validate API in `khqr_test.go`, plus
focused unit tests for CRC16/MD5/TLV parsing (`util`), individual field
validators (`generator`), and tag decoding (`decoder`).
