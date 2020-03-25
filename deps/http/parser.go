package http

import (
	"log"

	"gopkg.in/yaml.v2"

	"codezero/deps"
)

func ParseSpec(spec deps.Spec) (r Handler, err error) {
	err = yaml.Unmarshal([]byte(spec), &r)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	return
}