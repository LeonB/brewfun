package brewfun

type HopSubstitutionChart []HopSubstitutionChartEntry

type HopSubstitutionChartEntry struct {
	Hop         Hop            `json:"hop"`
	Substitutes HopSubstitutes `json:"substitutes"`
	Source      string         `json:"source"`
}

type HopSubstitutes []HopSubstitute

type HopSubstitute struct {
	Hop   Hop     `json:"hop"`
	Match float64 `json:"match"`
}

type Hop struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
