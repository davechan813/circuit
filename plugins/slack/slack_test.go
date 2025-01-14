package slack_test

import (
	"testing"
	"time"

	"github.com/codeamp/circuit/plugins"
	_ "github.com/codeamp/circuit/plugins/slack"
	"github.com/codeamp/circuit/test"
	"github.com/codeamp/transistor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

type TestSuite struct {
	suite.Suite
	transistor *transistor.Transistor
}

var viperConfig = []byte(`
plugins:
  slack:
    workers: 1
`)

func (suite *TestSuite) SetupSuite() {
	suite.transistor, _ = test.SetupPluginTest(viperConfig)
	go suite.transistor.Run()
}

func (suite *TestSuite) TearDownSuite() {
	suite.transistor.Stop()
	time.Sleep(1 * time.Second)
}

func getTestProjectExtensionPayload() plugins.ProjectExtension {
	payload := plugins.ProjectExtension{
		Project: plugins.Project{
			Slug:       "slack",
			Repository: "checkr/deploy-test",
		},
		Environment: "testing",
	}

	return payload
}

func getTestNotificationExtensionPayload() plugins.NotificationExtension {
	deploytestHash := "4930db36d9ef6ef4e6a986b6db2e40ec477c7bc9"

	payload := plugins.NotificationExtension{
		Release: plugins.Release{
			User: "test@checkr.com",
			Project: plugins.Project{
				Slug:       "slack",
				Repository: "checkr/deploy-test",
			},
			Git: plugins.Git{
				Url:           "https://github.com/checkr/deploy-test.git",
				Protocol:      "HTTPS",
				Branch:        "master",
				RsaPrivateKey: "",
				RsaPublicKey:  "",
				Workdir:       "/tmp/something",
			},
			HeadFeature: plugins.Feature{
				Hash:       deploytestHash,
				ParentHash: deploytestHash,
				User:       "",
				Message:    "Test",
			},
			TailFeature: plugins.Feature{
				Hash:       deploytestHash,
				ParentHash: deploytestHash,
				User:       "",
				Message:    "Test",
			},
			Environment: "testing",
		},
		Project: plugins.Project{
			Slug:       "slack",
			Repository: "checkr/deploy-test",
		},
		Environment: "testing",
	}

	return payload
}

func (suite *TestSuite) TestSlackNotifySuccessfulRelease() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var ev, re transistor.Event
	var err error

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "http://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "http://hooks.slack.com/services/token/token/valid_token",
		httpmock.NewStringResponder(200, "{}"))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("complete"), re.State)

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/invalid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "https://hooks.slack.com/services/token/token/invalid_token",
		httpmock.NewStringResponder(403, ""))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("failed"), re.State)

	ev = transistor.NewEvent(transistor.EventName("slack:notify"), transistor.GetAction("status"), getTestNotificationExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)
	ev.AddArtifact("message", "success", false)
	ev.AddArtifact("dashboard_url", "URL", false)

	ev.State = transistor.GetState("complete")
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	ev.State = transistor.GetState("failed")
	ev.AddArtifact("message", "failed", false)
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
}

func (suite *TestSuite) TestSlackNotifyFailedRelease() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var ev, re transistor.Event
	var err error

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "http://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "http://hooks.slack.com/services/token/token/valid_token",
		httpmock.NewStringResponder(200, "{}"))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("complete"), re.State)

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/invalid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "https://hooks.slack.com/services/token/token/invalid_token",
		httpmock.NewStringResponder(403, ""))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("failed"), re.State)

	ev = transistor.NewEvent(transistor.EventName("slack:notify"), transistor.GetAction("status"), getTestNotificationExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)
	ev.AddArtifact("message", "failed", false)
	ev.AddArtifact("dashboard_url", "URL", false)

	ev.State = transistor.GetState("complete")
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	ev.State = transistor.GetState("failed")
	ev.AddArtifact("message", "failed", false)
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
}

func (suite *TestSuite) TestSlackNotifyCanceledRelease() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var ev, re transistor.Event
	var err error

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "http://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "http://hooks.slack.com/services/token/token/valid_token",
		httpmock.NewStringResponder(200, "{}"))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("complete"), re.State)

	ev = transistor.NewEvent(plugins.GetEventName("slack"), transistor.GetAction("create"), getTestProjectExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/invalid_token", false)
	ev.AddArtifact("channel", "devops-test", false)

	httpmock.RegisterResponder("POST", "https://hooks.slack.com/services/token/token/invalid_token",
		httpmock.NewStringResponder(403, ""))

	suite.transistor.Events <- ev

	re, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	assert.Equal(suite.T(), transistor.GetState("failed"), re.State)

	ev = transistor.NewEvent(transistor.EventName("slack:notify"), transistor.GetAction("status"), getTestNotificationExtensionPayload())
	ev.AddArtifact("webhook_url", "https://hooks.slack.com/services/token/token/valid_token", false)
	ev.AddArtifact("channel", "devops-test", false)
	ev.AddArtifact("message", "canceled", false)
	ev.AddArtifact("dashboard_url", "URL", false)

	ev.State = transistor.GetState("complete")
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	ev.State = transistor.GetState("failed")
	ev.AddArtifact("message", "failed", false)
	suite.transistor.Events <- ev
	_, err = suite.transistor.GetTestEvent(plugins.GetEventName("slack:notify"), transistor.GetAction("status"), 100)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
}

func TestSlack(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
