package framework

type HTTPMethod string

type WebServer interface {
	ListenAndServe()

	Get(route string, fn WebHandler)
	Post(route string, fn WebHandler)
	Put(route string, fn WebHandler)
	Delete(route string, fn WebHandler)

	//Websocket(route string)
}

type WebHandler func(WebContext) WebResult

type WebResponse struct {
	Error  *FrameworkError "error"
	Result interface{}     "result"
}

func NewWebResponse(result interface{}, err error) *WebResponse {

	var appErr *FrameworkError

	if err != nil {
		switch err.(type) {
		case *FrameworkError:
			appErr = err.(*FrameworkError)
		default:
			appErr = ToError(Error_Generic, err)
		}
	}

	return &WebResponse{Error: appErr, Result: result}
}
