package rest

import (
	"einvite/common/services"
	"einvite/framework"
	"fmt"
	"time"
)

type EventController struct {
	eventService services.EventService
}

func (this EventController) Test(ctx framework.WebContext) framework.WebResult {

	//note: time.sleep frees the thread and allows other goroutines to execute
	//sending 1000 concurrent requests with AB all complete within 5seconds

	time.Sleep(1 * time.Second)

	fmt.Println(ctx.Header())

	fmt.Println("Test invoked")

	return ctx.Text("test operation")
}

func (this EventController) CreateEvent(ctx framework.WebContext) framework.WebResult {

	return ctx.Unauthorized("Thou shall not create an event")
}

func NewEventController(eventService services.EventService) *EventController {

	return &EventController{eventService}
}
