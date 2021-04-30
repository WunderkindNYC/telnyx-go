
/*
 * Telnyx API
 *
 * SIP trunking, SMS, MMS, Call Control and Telephony Data Services.
 *
 * API version: 2.0.0
 * Contact: support@telnyx.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package telnyx

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DebuggingApiService service
/*
DebuggingApiService List call events
Filters call events by given filter parameters. Events are ordered by &#x60;event_timestamp&#x60;. If filter for &#x60;call_leg_id&#x60; or &#x60;call_session_id&#x60; is not present, it only filters events from the last 24 hours.  **Note**: Only one &#x60;filter[event_timestamp]&#x60; can be passed. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DebuggingApiCallControlDebuggingEventListOpts - Optional Parameters:
     * @param "FilterCallLegId" (optional.Interface of string) -  The unique identifier of an individual call leg.
     * @param "FilterCallSessionId" (optional.Interface of string) -  The unique identifier of the call control session. A session may include multiple call leg events.
     * @param "FilterStatus" (optional.String) -  Event status
     * @param "FilterType" (optional.String) -  Event type
     * @param "FilterEventTimestampGt" (optional.String) -  Event timestamp: greater than
     * @param "FilterEventTimestampGte" (optional.String) -  Event timestamp: greater than or equal
     * @param "FilterEventTimestampLt" (optional.String) -  Event timestamp: lower than
     * @param "FilterEventTimestampLte" (optional.String) -  Event timestamp: lower than or equal
     * @param "FilterEventTimestampEq" (optional.String) -  Event timestamp: equal
@return InlineResponse2007
*/

type DebuggingApiCallControlDebuggingEventListOpts struct {
    FilterCallLegId optional.Interface
    FilterCallSessionId optional.Interface
    FilterStatus optional.String
    FilterType optional.String
    FilterEventTimestampGt optional.String
    FilterEventTimestampGte optional.String
    FilterEventTimestampLt optional.String
    FilterEventTimestampLte optional.String
    FilterEventTimestampEq optional.String
}

func (a *DebuggingApiService) CallControlDebuggingEventList(ctx context.Context, localVarOptionals *DebuggingApiCallControlDebuggingEventListOpts) (InlineResponse2007, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2007
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/call_events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.FilterCallLegId.IsSet() {
		localVarQueryParams.Add("filter[call_leg_id]", parameterToString(localVarOptionals.FilterCallLegId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCallSessionId.IsSet() {
		localVarQueryParams.Add("filter[call_session_id]", parameterToString(localVarOptionals.FilterCallSessionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterStatus.IsSet() {
		localVarQueryParams.Add("filter[status]", parameterToString(localVarOptionals.FilterStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterType.IsSet() {
		localVarQueryParams.Add("filter[type]", parameterToString(localVarOptionals.FilterType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterEventTimestampGt.IsSet() {
		localVarQueryParams.Add("filter[event_timestamp][gt]", parameterToString(localVarOptionals.FilterEventTimestampGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterEventTimestampGte.IsSet() {
		localVarQueryParams.Add("filter[event_timestamp][gte]", parameterToString(localVarOptionals.FilterEventTimestampGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterEventTimestampLt.IsSet() {
		localVarQueryParams.Add("filter[event_timestamp][lt]", parameterToString(localVarOptionals.FilterEventTimestampLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterEventTimestampLte.IsSet() {
		localVarQueryParams.Add("filter[event_timestamp][lte]", parameterToString(localVarOptionals.FilterEventTimestampLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterEventTimestampEq.IsSet() {
		localVarQueryParams.Add("filter[event_timestamp][eq]", parameterToString(localVarOptionals.FilterEventTimestampEq.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2007
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v InlineResponseDefault1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}