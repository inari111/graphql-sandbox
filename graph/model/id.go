package model

import "encoding/base64"

const apiVersion = "01"

type GlobalIDGenerator interface {
	Generate(objName string, id string) string
}

func NewGlobalIDGenerator() GlobalIDGenerator {
	return &globalIDGenerator{}
}

type globalIDGenerator struct{}

func (g *globalIDGenerator) Generate(objName string, id string) string {
	return base64.URLEncoding.EncodeToString([]byte(apiVersion + ":" + objName + ":" + id))
}
