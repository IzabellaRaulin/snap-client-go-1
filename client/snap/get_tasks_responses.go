package snap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intelsdi-x/snap-client-go/models"
)

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*GetTasksOK handles this case with default header values.

TasksResp returns a list of created tasks.
*/
type GetTasksOK struct {
	Payload GetTasksOKBody
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTasksOKBody get tasks o k body
swagger:model GetTasksOKBody
*/
type GetTasksOKBody struct {

	// tasks
	// Required: true
	Tasks models.Tasks `json:"tasks"`
}

// Validate validates this get tasks o k body
func (o *GetTasksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTasks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTasksOKBody) validateTasks(formats strfmt.Registry) error {

	if err := validate.Required("getTasksOK"+"."+"tasks", "body", o.Tasks); err != nil {
		return err
	}

	return nil
}
