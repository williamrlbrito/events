package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type TestEventHandler struct {
	ID int
}

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
	suite.handlerOne = TestEventHandler{ID: 1}
	suite.handlerTwo = TestEventHandler{ID: 2}
	suite.handlerThree = TestEventHandler{ID: 3}
	suite.eventDispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherRegister() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))
	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))
	assert.Equal(suite.T(), &suite.handlerOne, suite.eventDispatcher.handlers[suite.eventOne.GetName()][0])
	assert.Equal(suite.T(), &suite.handlerTwo, suite.eventDispatcher.handlers[suite.eventOne.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherRegisterWithSameHandler() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherClear() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventTwo.GetName(), &suite.handlerThree)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventTwo.GetName()]))

	err = suite.eventDispatcher.Clear()
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherHas() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventTwo.GetName(), &suite.handlerThree)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.eventTwo.GetName()]))

	suite.True(suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerOne))
	suite.True(suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerTwo))
	suite.False(suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerThree))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcherDispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.eventOne)
	suite.eventDispatcher.Register(suite.eventOne.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.eventOne)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
