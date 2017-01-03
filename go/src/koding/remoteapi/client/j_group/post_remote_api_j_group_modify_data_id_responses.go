package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJGroupModifyDataIDReader is a Reader for the PostRemoteAPIJGroupModifyDataID structure.
type PostRemoteAPIJGroupModifyDataIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupModifyDataIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupModifyDataIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupModifyDataIDOK creates a PostRemoteAPIJGroupModifyDataIDOK with default headers values
func NewPostRemoteAPIJGroupModifyDataIDOK() *PostRemoteAPIJGroupModifyDataIDOK {
	return &PostRemoteAPIJGroupModifyDataIDOK{}
}

/*PostRemoteAPIJGroupModifyDataIDOK handles this case with default header values.

includes error if something went wrong
*/
type PostRemoteAPIJGroupModifyDataIDOK struct {
	Payload *models.Error
}

func (o *PostRemoteAPIJGroupModifyDataIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.modifyData/{id}][%d] postRemoteApiJGroupModifyDataIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupModifyDataIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupModifyDataIDBody post remote API j group modify data ID body
swagger:model PostRemoteAPIJGroupModifyDataIDBody
*/
type PostRemoteAPIJGroupModifyDataIDBody struct {

	// github organization token
	// Required: true
	GithubOrganizationToken *string `json:"github.organizationToken"`
}
