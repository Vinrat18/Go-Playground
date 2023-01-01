package schema

type AssetDerived map[string]string

type AssetLink struct {
	Id     string
	Source string
	Type   string
}

type AssetRelationships map[string][]string

type AssetData map[string]string

type AssetAttribute struct {
	key   string
	value string
}

type Asset struct {
	Id            string             `json:"id,omitempty" validate:"required"`
	Source        string             `json:"source,omitempty" validate:"required"`
	Type          string             `json:"type,omitempty"`
	Derived       AssetDerived       `json:"derived,omitempty"`
	Links         []AssetLink        `json:"links,omitempty"`
	Relationships AssetRelationships `json:"relationships,omitempty"`
	Data          AssetData          `json:"data,omitempty"`
	Attributes    []AssetAttribute   `json:"attributes,omitempty"`
}
