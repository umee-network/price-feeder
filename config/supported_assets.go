package config

import (
	"github.com/ojo-network/price-feeder/oracle/provider"
	"github.com/ojo-network/price-feeder/oracle/types"
)

type APIKeyRequired bool

var (
	// SupportedProviders defines a lookup table of all the supported currency API
	// providers and whether or not they require an API key to be passed in.
	SupportedProviders = map[types.ProviderName]APIKeyRequired{
		provider.ProviderKraken:    false,
		provider.ProviderBinance:   false,
		provider.ProviderBinanceUS: false,
		provider.ProviderCrescent:  false,
		provider.ProviderOsmosisV2: false,
		provider.ProviderOkx:       false,
		provider.ProviderHuobi:     false,
		provider.ProviderGate:      false,
		provider.ProviderCoinbase:  false,
		provider.ProviderBitget:    false,
		provider.ProviderMexc:      false,
		provider.ProviderCrypto:    false,
		provider.ProviderPolygon:   true,
		provider.ProviderMock:      false,
	}

	// SupportedQuotes defines a lookup table for which assets we support
	// using as quotes.
	SupportedQuotes = map[string]struct{}{
		DenomUSD: {},
		"USDC":   {},
		"USDT":   {},
		"DAI":    {},
		"BTC":    {},
		"ETH":    {},
		"ATOM":   {},
		"OSMO":   {},
	}

	// SupportedForexCurrencies defines a lookup table for all the supported
	// Forex currencies
	SupportedForexCurrencies = map[string]struct{}{
		"AED": {},
		"AFN": {},
		"ALL": {},
		"AMD": {},
		"ANG": {},
		"AOA": {},
		"ARS": {},
		"AUD": {},
		"AWG": {},
		"AZN": {},
		"BAM": {},
		"BBD": {},
		"BDT": {},
		"BGN": {},
		"BHD": {},
		"BIF": {},
		"BMD": {},
		"BND": {},
		"BOB": {},
		"BRL": {},
		"BSD": {},
		"BTN": {},
		"BWP": {},
		"BZD": {},
		"CAD": {},
		"CDF": {},
		"CHF": {},
		"CLF": {},
		"CLP": {},
		"CNH": {},
		"CNY": {},
		"COP": {},
		"CUP": {},
		"CVE": {},
		"CZK": {},
		"DJF": {},
		"DKK": {},
		"DOP": {},
		"DZD": {},
		"EGP": {},
		"ERN": {},
		"ETB": {},
		"EUR": {},
		"FJD": {},
		"FKP": {},
		"GBP": {},
		"GEL": {},
		"GHS": {},
		"GIP": {},
		"GMD": {},
		"GNF": {},
		"GTQ": {},
		"GYD": {},
		"HKD": {},
		"HNL": {},
		"HRK": {},
		"HTG": {},
		"HUF": {},
		"ICP": {},
		"IDR": {},
		"ILS": {},
		"INR": {},
		"IQD": {},
		"IRR": {},
		"ISK": {},
		"JEP": {},
		"JMD": {},
		"JOD": {},
		"JPY": {},
		"KES": {},
		"KGS": {},
		"KHR": {},
		"KMF": {},
		"KPW": {},
		"KRW": {},
		"KWD": {},
		"KYD": {},
		"KZT": {},
		"LAK": {},
		"LBP": {},
		"LKR": {},
		"LRD": {},
		"LSL": {},
		"LYD": {},
		"MAD": {},
		"MDL": {},
		"MGA": {},
		"MKD": {},
		"MMK": {},
		"MNT": {},
		"MOP": {},
		"MRO": {},
		"MRU": {},
		"MUR": {},
		"MVR": {},
		"MWK": {},
		"MXN": {},
		"MYR": {},
		"MZN": {},
		"NAD": {},
		"NGN": {},
		"NOK": {},
		"NPR": {},
		"NZD": {},
		"OMR": {},
		"PAB": {},
		"PEN": {},
		"PGK": {},
		"PHP": {},
		"PKR": {},
		"PLN": {},
		"PYG": {},
		"QAR": {},
		"RON": {},
		"RSD": {},
		"RUB": {},
		"RUR": {},
		"RWF": {},
		"SAR": {},
		"SBD": {},
		"SCR": {},
		"SDG": {},
		"SDR": {},
		"SEK": {},
		"SGD": {},
		"SHP": {},
		"SLL": {},
		"SOS": {},
		"SRD": {},
		"SYP": {},
		"SZL": {},
		"THB": {},
		"TJS": {},
		"TMT": {},
		"TND": {},
		"TOP": {},
		"TRY": {},
		"TTD": {},
		"TWD": {},
		"TZS": {},
		"UAH": {},
		"UGX": {},
		"USD": {},
		"UYU": {},
		"UZS": {},
		"VND": {},
		"VUV": {},
		"WST": {},
		"XAF": {},
		"XCD": {},
		"XDR": {},
		"XOF": {},
		"XPF": {},
		"YER": {},
		"ZAR": {},
		"ZMW": {},
		"ZWL": {},
	}
)
