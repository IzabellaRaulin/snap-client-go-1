package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/intelsdi-x/snap-client-go/models"
)

// DeletePluginConfigItemReader is a Reader for the DeletePluginConfigItem structure.
type DeletePluginConfigItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginConfigItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePluginConfigItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeletePluginConfigItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeletePluginConfigItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePluginConfigItemOK creates a DeletePluginConfigItemOK with default headers values
func NewDeletePluginConfigItemOK() *DeletePluginConfigItemOK {
	return &DeletePluginConfigItemOK{}
}

/*DeletePluginConfigItemOK handles this case with default header values.

PluginConfigResponse represents the response of a plugin config items.
*/
type DeletePluginConfigItemOK struct {
	Payload models.ConfigDataNode
}

func (o *DeletePluginConfigItemOK) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{ptype}/{pname}/{pversion}/config][%d] deletePluginConfigItemOK  %+v", 200, o.Payload)
}

func (o *DeletePluginConfigItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginConfigItemBadRequest creates a DeletePluginConfigItemBadRequest with default headers values
func NewDeletePluginConfigItemBadRequest() *DeletePluginConfigItemBadRequest {
	return &DeletePluginConfigItemBadRequest{}
}

/*DeletePluginConfigItemBadRequest handles this case with default header values.

ErrorResponse represents the Snap error response type.

It includes an error message and a map of fields.
*/
type DeletePluginConfigItemBadRequest struct {
	Payload *models.Error
}

func (o *DeletePluginConfigItemBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{ptype}/{pname}/{pversion}/config][%d] deletePluginConfigItemBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePluginConfigItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginConfigItemUnauthorized creates a DeletePluginConfigItemUnauthorized with default headers values
func NewDeletePluginConfigItemUnauthorized() *DeletePluginConfigItemUnauthorized {
	return &DeletePluginConfigItemUnauthorized{}
}

/*DeletePluginConfigItemUnauthorized handles this case with default header values.

UnauthResponse returns Unauthorized error struct message.
*/
type DeletePluginConfigItemUnauthorized struct {
	Payload *models.UnauthError
}

func (o *DeletePluginConfigItemUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{ptype}/{pname}/{pversion}/config][%d] deletePluginConfigItemUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePluginConfigItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
