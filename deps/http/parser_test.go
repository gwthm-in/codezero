package http

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"codezero/deps"
)

func TestParseDependency(t *testing.T) {
	spec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz.yml")
	require.NoError(t, err, "yamlFile.Get error")

	actualHandler, err := parseDependency(deps.Spec(spec))
	require.NoError(t, err, "http.parseDependency error")

	expectedHandler := Handler{
		Deps: Dependencies{
			"service_xyz": &Dependency{
				Sits: Situations{
					"response_2xx": &Situation{
						Req: &Request{
							Method:  Get,
							Path:    "/v1/ping",
							Query:   Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
							Headers: Header{"Accept-Encoding": []string{"gzip", "compress"}},
						},
						Res: &Response{
							Body: fmt.Sprintf("%s\n", `{"ping": "pong"}`),
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedHandler.Deps, actualHandler)
}

func TestParseScenario(t *testing.T) {
	actualScenario, err := parseScenario(`
service_xyz:
  response_2xx:
    port: 8010
`)
	require.NoError(t, err)

	expectedScenario := Scenario{"service_xyz": map[SituationName]Spec{"response_2xx": {Port: 8010}}}
	assert.Equal(t, expectedScenario, actualScenario)
}
