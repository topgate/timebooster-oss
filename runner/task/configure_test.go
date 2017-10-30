package task

import (
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// configがパースできる
func TestParseConfigure(t *testing.T) {
	buf, _ := ioutil.ReadFile("../examples/timebooster.yaml")

	assert.NotNil(t, buf)

	cfg := Configure{}
	yaml.Unmarshal(buf, &cfg)
	assert.NotEqual(t, len(cfg.Env.Variable), 0)
	assert.NotEqual(t, len(cfg.Env.Cache), 0)
	assert.NotEqual(t, len(cfg.Task.Exec), 0)
	assert.NotEqual(t, cfg.Task.Exec[0].Dockerfile, "")
	assert.NotEqual(t, cfg.Task.Exec[0].DockerImage, "")
	assert.NotEqual(t, len(cfg.Task.Exec[0].Cmd), 0)
}
