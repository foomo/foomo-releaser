package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateLink(t *testing.T) {
	res := calculateJiraLink("hotfix/123")
	assert.Equal(t, "", res)

	res = calculateJiraLink("feature/ECOMDEV-123-test")
	assert.Equal(t, "https://jira.globuswiki.com/browse/ECOMDEV-123", res)

}
