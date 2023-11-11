package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name     string
	Playload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Playload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (h *TestEventHandler) Handle(event EventInterface) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	eventOne        TestEvent
	eventTwo        TestEvent
	handlerOne      TestEventHandler
	handlerTwo      TestEventHandler
	handlerThree    TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventOne = TestEvent{Name: "eventOne", Playload: "eventOne"}
	suite.eventTwo = TestEvent{Name: "eventTwo", Playload: "eventTwo"}
	suite.handlerOne = TestEventHandler{}
	suite.handlerTwo = TestEventHandler{}
	suite.handlerThree = TestEventHandler{}
	suite.eventDispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherRegister() {
	assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
