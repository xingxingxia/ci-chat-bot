// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns summary information about stack sets that are associated with the user.
//   - [Self-managed permissions] If you set the CallAs parameter to SELF while
//     signed in to your Amazon Web Services account, ListStackSets returns all
//     self-managed stack sets in your Amazon Web Services account.
//   - [Service-managed permissions] If you set the CallAs parameter to SELF while
//     signed in to the organization's management account, ListStackSets returns all
//     stack sets in the management account.
//   - [Service-managed permissions] If you set the CallAs parameter to
//     DELEGATED_ADMIN while signed in to your member account, ListStackSets returns
//     all stack sets with service-managed permissions in the management account.
func (c *Client) ListStackSets(ctx context.Context, params *ListStackSetsInput, optFns ...func(*Options)) (*ListStackSetsOutput, error) {
	if params == nil {
		params = &ListStackSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackSets", params, optFns, c.addOperationListStackSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStackSetsInput struct {

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the management account or as a delegated administrator in a
	// member account. By default, SELF is specified. Use SELF for stack sets with
	// self-managed permissions.
	//   - If you are signed in to the management account, specify SELF .
	//   - If you are signed in to a delegated administrator account, specify
	//   DELEGATED_ADMIN . Your Amazon Web Services account must be registered as a
	//   delegated administrator in the management account. For more information, see
	//   Register a delegated administrator (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	//   in the CloudFormation User Guide.
	CallAs types.CallAs

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int32

	// If the previous paginated request didn't return all the remaining results, the
	// response object's NextToken parameter value is set to a token. To retrieve the
	// next set of results, call ListStackSets again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null .
	NextToken *string

	// The status of the stack sets that you want to get summary information about.
	Status types.StackSetStatus

	noSmithyDocumentSerde
}

type ListStackSetsOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListStackInstances again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null .
	NextToken *string

	// A list of StackSetSummary structures that contain information about the user's
	// stack sets.
	Summaries []types.StackSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStackSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStackSets"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackSets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListStackSetsAPIClient is a client that implements the ListStackSets operation.
type ListStackSetsAPIClient interface {
	ListStackSets(context.Context, *ListStackSetsInput, ...func(*Options)) (*ListStackSetsOutput, error)
}

var _ ListStackSetsAPIClient = (*Client)(nil)

// ListStackSetsPaginatorOptions is the paginator options for ListStackSets
type ListStackSetsPaginatorOptions struct {
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStackSetsPaginator is a paginator for ListStackSets
type ListStackSetsPaginator struct {
	options   ListStackSetsPaginatorOptions
	client    ListStackSetsAPIClient
	params    *ListStackSetsInput
	nextToken *string
	firstPage bool
}

// NewListStackSetsPaginator returns a new ListStackSetsPaginator
func NewListStackSetsPaginator(client ListStackSetsAPIClient, params *ListStackSetsInput, optFns ...func(*ListStackSetsPaginatorOptions)) *ListStackSetsPaginator {
	if params == nil {
		params = &ListStackSetsInput{}
	}

	options := ListStackSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStackSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStackSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStackSets page.
func (p *ListStackSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStackSetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListStackSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStackSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStackSets",
	}
}