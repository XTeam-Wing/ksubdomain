package result

type Result struct {
	Subdomain string   `json:"subdomain"`
	Answers   []Answer `json:"answers"`
}

type Answer struct {
	Type  string   `json:"type"`
	Value []string `json:"value"`
}
