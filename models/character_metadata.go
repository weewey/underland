package models

type Creator struct {
	Address string `json:"address"`
	Share   int    `json:"share"`
}

type Files struct {
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

type Properties struct {
	Files    []Files   `json:"files"`
	Creators []Creator `json:"creators"`
}

type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type CharacterMetaData struct {
	Name        string        `json:"name"`
	Symbol      string        `json:"symbol"`
	Image       string        `json:"image"`
	SocialIndex int           `json:"social_index"`
	Attributes  []*Attributes `json:"attributes"`
	Properties  *Properties   `json:"properties"`
}

const (
	SOCIAL       = "SOCIAL"
	SOCIAL_INDEX = "SOCIAL_INDEX"
	CONTRIBUTION = "CONTRIBUTION"
	BEHAVIOUR    = "BEHAVIOUR"
	RESPECT      = "RESPECT"
	INFLUENCE    = "INFLUENCE"
)
