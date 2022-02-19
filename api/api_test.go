package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSearchByNameValidServer expects one specific server back from Speedtest API.
func TestSearchByNameValidServer(t *testing.T) {
	expect := &ServerResponse{
		ID:      "4302",
		Name:    "Milan",
		Sponsor: "Vodafone IT",
		Host:    "speedtest.vodafone.it.prod.hosts.ooklaserver.net:8080",
	}

	servers, err := SearchByName("vodafone it")
	assert.Nil(t, err, "error with Speedtest request")

	assert.Equal(t, 1, len(servers), "should get back one server")

	s := servers[0]

	assert.Equal(t, expect.ID, s.ID)
	assert.Equal(t, expect.Name, s.Name)
	assert.Equal(t, expect.Sponsor, s.Sponsor)
	assert.Equal(t, expect.Host, s.Host)
}

// TestSearchByNameEmptyResponse expects zero servers back from Speedtest API,
// as the query shouldn't match existing ones.
func TestSearchByNameEmptyResponse(t *testing.T) {
	servers, err := SearchByName("nonexistentservername")
	assert.Nil(t, err, "error with Speedtest request")

	assert.Equal(t, 0, len(servers), "should get back no servers")
}
