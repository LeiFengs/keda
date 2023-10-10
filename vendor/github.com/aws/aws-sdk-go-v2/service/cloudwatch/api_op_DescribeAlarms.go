// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Retrieves the specified alarms. You can filter the results by specifying a
// prefix for the alarm name, the alarm state, or a prefix for any action. To use
// this operation and return information about composite alarms, you must be signed
// on with the cloudwatch:DescribeAlarms permission that is scoped to * . You can't
// return information about composite alarms if your cloudwatch:DescribeAlarms
// permission has a narrower scope.
func (c *Client) DescribeAlarms(ctx context.Context, params *DescribeAlarmsInput, optFns ...func(*Options)) (*DescribeAlarmsOutput, error) {
	if params == nil {
		params = &DescribeAlarmsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlarms", params, optFns, c.addOperationDescribeAlarmsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlarmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmsInput struct {

	// Use this parameter to filter the results of the operation to only those alarms
	// that use a certain alarm action. For example, you could specify the ARN of an
	// SNS topic to find all alarms that send notifications to that topic.
	ActionPrefix *string

	// An alarm name prefix. If you specify this parameter, you receive information
	// about all alarms that have names that start with this prefix. If this parameter
	// is specified, you cannot specify AlarmNames .
	AlarmNamePrefix *string

	// The names of the alarms to retrieve information about.
	AlarmNames []string

	// Use this parameter to specify whether you want the operation to return metric
	// alarms or composite alarms. If you omit this parameter, only metric alarms are
	// returned, even if composite alarms exist in the account. For example, if you
	// omit this parameter or specify MetricAlarms , the operation returns only a list
	// of metric alarms. It does not return any composite alarms, even if composite
	// alarms exist in the account. If you specify CompositeAlarms , the operation
	// returns only a list of composite alarms, and does not return any metric alarms.
	AlarmTypes []types.AlarmType

	// If you use this parameter and specify the name of a composite alarm, the
	// operation returns information about the "children" alarms of the alarm you
	// specify. These are the metric alarms and composite alarms referenced in the
	// AlarmRule field of the composite alarm that you specify in ChildrenOfAlarmName .
	// Information about the composite alarm that you name in ChildrenOfAlarmName is
	// not returned. If you specify ChildrenOfAlarmName , you cannot specify any other
	// parameters in the request except for MaxRecords and NextToken . If you do so,
	// you receive a validation error. Only the Alarm Name , ARN , StateValue
	// (OK/ALARM/INSUFFICIENT_DATA), and StateUpdatedTimestamp information are
	// returned by this operation when you use this parameter. To get complete
	// information about these alarms, perform another DescribeAlarms operation and
	// specify the parent alarm names in the AlarmNames parameter.
	ChildrenOfAlarmName *string

	// The maximum number of alarm descriptions to retrieve.
	MaxRecords *int32

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string

	// If you use this parameter and specify the name of a metric or composite alarm,
	// the operation returns information about the "parent" alarms of the alarm you
	// specify. These are the composite alarms that have AlarmRule parameters that
	// reference the alarm named in ParentsOfAlarmName . Information about the alarm
	// that you specify in ParentsOfAlarmName is not returned. If you specify
	// ParentsOfAlarmName , you cannot specify any other parameters in the request
	// except for MaxRecords and NextToken . If you do so, you receive a validation
	// error. Only the Alarm Name and ARN are returned by this operation when you use
	// this parameter. To get complete information about these alarms, perform another
	// DescribeAlarms operation and specify the parent alarm names in the AlarmNames
	// parameter.
	ParentsOfAlarmName *string

	// Specify this parameter to receive information only about alarms that are
	// currently in the state that you specify.
	StateValue types.StateValue

	noSmithyDocumentSerde
}

type DescribeAlarmsOutput struct {

	// The information about any composite alarms returned by the operation.
	CompositeAlarms []types.CompositeAlarm

	// The information about any metric alarms returned by the operation.
	MetricAlarms []types.MetricAlarm

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAlarmsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAlarms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAlarms{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDescribeAlarmsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarms(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeAlarmsAPIClient is a client that implements the DescribeAlarms
// operation.
type DescribeAlarmsAPIClient interface {
	DescribeAlarms(context.Context, *DescribeAlarmsInput, ...func(*Options)) (*DescribeAlarmsOutput, error)
}

var _ DescribeAlarmsAPIClient = (*Client)(nil)

// DescribeAlarmsPaginatorOptions is the paginator options for DescribeAlarms
type DescribeAlarmsPaginatorOptions struct {
	// The maximum number of alarm descriptions to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAlarmsPaginator is a paginator for DescribeAlarms
type DescribeAlarmsPaginator struct {
	options   DescribeAlarmsPaginatorOptions
	client    DescribeAlarmsAPIClient
	params    *DescribeAlarmsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAlarmsPaginator returns a new DescribeAlarmsPaginator
func NewDescribeAlarmsPaginator(client DescribeAlarmsAPIClient, params *DescribeAlarmsInput, optFns ...func(*DescribeAlarmsPaginatorOptions)) *DescribeAlarmsPaginator {
	if params == nil {
		params = &DescribeAlarmsInput{}
	}

	options := DescribeAlarmsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAlarmsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAlarmsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAlarms page.
func (p *DescribeAlarmsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAlarmsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeAlarms(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// AlarmExistsWaiterOptions are waiter options for AlarmExistsWaiter
type AlarmExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AlarmExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AlarmExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeAlarmsInput, *DescribeAlarmsOutput, error) (bool, error)
}

// AlarmExistsWaiter defines the waiters for AlarmExists
type AlarmExistsWaiter struct {
	client DescribeAlarmsAPIClient

	options AlarmExistsWaiterOptions
}

// NewAlarmExistsWaiter constructs a AlarmExistsWaiter.
func NewAlarmExistsWaiter(client DescribeAlarmsAPIClient, optFns ...func(*AlarmExistsWaiterOptions)) *AlarmExistsWaiter {
	options := AlarmExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = alarmExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AlarmExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AlarmExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *AlarmExistsWaiter) Wait(ctx context.Context, params *DescribeAlarmsInput, maxWaitDur time.Duration, optFns ...func(*AlarmExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AlarmExists waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *AlarmExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeAlarmsInput, maxWaitDur time.Duration, optFns ...func(*AlarmExistsWaiterOptions)) (*DescribeAlarmsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeAlarms(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for AlarmExists waiter")
}

func alarmExistsStateRetryable(ctx context.Context, input *DescribeAlarmsInput, output *DescribeAlarmsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(MetricAlarms[]) > `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	return true, nil
}

// CompositeAlarmExistsWaiterOptions are waiter options for
// CompositeAlarmExistsWaiter
type CompositeAlarmExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CompositeAlarmExistsWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, CompositeAlarmExistsWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeAlarmsInput, *DescribeAlarmsOutput, error) (bool, error)
}

// CompositeAlarmExistsWaiter defines the waiters for CompositeAlarmExists
type CompositeAlarmExistsWaiter struct {
	client DescribeAlarmsAPIClient

	options CompositeAlarmExistsWaiterOptions
}

// NewCompositeAlarmExistsWaiter constructs a CompositeAlarmExistsWaiter.
func NewCompositeAlarmExistsWaiter(client DescribeAlarmsAPIClient, optFns ...func(*CompositeAlarmExistsWaiterOptions)) *CompositeAlarmExistsWaiter {
	options := CompositeAlarmExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = compositeAlarmExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CompositeAlarmExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CompositeAlarmExists waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *CompositeAlarmExistsWaiter) Wait(ctx context.Context, params *DescribeAlarmsInput, maxWaitDur time.Duration, optFns ...func(*CompositeAlarmExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for CompositeAlarmExists waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *CompositeAlarmExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeAlarmsInput, maxWaitDur time.Duration, optFns ...func(*CompositeAlarmExistsWaiterOptions)) (*DescribeAlarmsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeAlarms(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for CompositeAlarmExists waiter")
}

func compositeAlarmExistsStateRetryable(ctx context.Context, input *DescribeAlarmsInput, output *DescribeAlarmsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(CompositeAlarms[]) > `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeAlarms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAlarms",
	}
}

type opDescribeAlarmsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeAlarmsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeAlarmsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "monitoring"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "monitoring"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("monitoring")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeAlarmsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeAlarmsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}