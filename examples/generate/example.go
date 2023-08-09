package main

type Example struct {
	Glossary `json:"glossary"`
}

type Glossary struct {
	GlossDiv `json:"GlossDiv"`
	Title    string `json:"title"`
}

type GlossDiv struct {
	GlossList `json:"GlossList"`
	Title     string `json:"title"`
}

type GlossList struct {
	GlossEntry `json:"GlossEntry"`
}

type GlossEntry struct {
	Abbrev    string   `json:"Abbrev"`
	Acronym   string   `json:"Acronym"`
	GlossDef  GlossDef `json:"GlossDef"`
	GlossSee  string   `json:"GlossSee"`
	GlossTerm string   `json:"GlossTerm"`
	ID        string   `json:"ID"`
	SortAs    string   `json:"SortAs"`
}

type GlossDef struct {
	GlossSeeAlso []string `json:"GlossSeeAlso"`
	Para         string   `json:"para"`
}
