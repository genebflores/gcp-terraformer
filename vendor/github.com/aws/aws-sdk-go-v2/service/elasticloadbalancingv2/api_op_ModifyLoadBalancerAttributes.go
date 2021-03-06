// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// The load balancer attributes.
	//
	// Attributes is a required field
	Attributes []LoadBalancerAttribute `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyLoadBalancerAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyLoadBalancerAttributesInput"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancer attributes.
	Attributes []LoadBalancerAttribute `type:"list"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyLoadBalancerAttributes = "ModifyLoadBalancerAttributes"

// ModifyLoadBalancerAttributesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Modifies the specified attributes of the specified Application Load Balancer
// or Network Load Balancer.
//
// If any of the specified attributes can't be modified as requested, the call
// fails. Any existing attributes that you do not modify retain their current
// values.
//
//    // Example sending a request using ModifyLoadBalancerAttributesRequest.
//    req := client.ModifyLoadBalancerAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/ModifyLoadBalancerAttributes
func (c *Client) ModifyLoadBalancerAttributesRequest(input *ModifyLoadBalancerAttributesInput) ModifyLoadBalancerAttributesRequest {
	op := &aws.Operation{
		Name:       opModifyLoadBalancerAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyLoadBalancerAttributesInput{}
	}

	req := c.newRequest(op, input, &ModifyLoadBalancerAttributesOutput{})
	return ModifyLoadBalancerAttributesRequest{Request: req, Input: input, Copy: c.ModifyLoadBalancerAttributesRequest}
}

// ModifyLoadBalancerAttributesRequest is the request type for the
// ModifyLoadBalancerAttributes API operation.
type ModifyLoadBalancerAttributesRequest struct {
	*aws.Request
	Input *ModifyLoadBalancerAttributesInput
	Copy  func(*ModifyLoadBalancerAttributesInput) ModifyLoadBalancerAttributesRequest
}

// Send marshals and sends the ModifyLoadBalancerAttributes API request.
func (r ModifyLoadBalancerAttributesRequest) Send(ctx context.Context) (*ModifyLoadBalancerAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyLoadBalancerAttributesResponse{
		ModifyLoadBalancerAttributesOutput: r.Request.Data.(*ModifyLoadBalancerAttributesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyLoadBalancerAttributesResponse is the response type for the
// ModifyLoadBalancerAttributes API operation.
type ModifyLoadBalancerAttributesResponse struct {
	*ModifyLoadBalancerAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyLoadBalancerAttributes request.
func (r *ModifyLoadBalancerAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
