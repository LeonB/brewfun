package brewfun

type HopSubstitutes []HopSubstitute

type HopSubstitute struct {
	HopA   Hop     `json:"hopA"`
	HopB   Hop     `json:"hopB"`
	Match  float64 `json:"match"`
	Source string  `json:"source"`
}

type Hop struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
