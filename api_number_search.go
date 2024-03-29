
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

type NumberSearchApiService service
/*
NumberSearchApiService List available phone numbers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *NumberSearchApiListAvailablePhoneNumbersOpts - Optional Parameters:
     * @param "FilterPhoneNumberStartsWith" (optional.String) -  Filter numbers starting with a pattern (meant to be used after &#x60;national_destination_code&#x60; filter has been set).
     * @param "FilterPhoneNumberEndsWith" (optional.String) -  Filter numbers ending with a pattern.
     * @param "FilterPhoneNumberContains" (optional.String) -  Filter numbers containing a pattern.
     * @param "FilterLocality" (optional.String) -  Filter phone numbers by city.
     * @param "FilterAdministrativeArea" (optional.String) -  Filter phone numbers by US state/CA province.
     * @param "FilterCountryCode" (optional.String) -  Filter phone numbers by ISO alpha-2 country code.
     * @param "FilterNationalDestinationCode" (optional.String) -  Filter by the national destination code of the number. This filter is only applicable to North American numbers.
     * @param "FilterRateCenter" (optional.String) -  Filter phone numbers by NANP rate center. This filter is only applicable to North American numbers.
     * @param "FilterNumberType" (optional.String) -  Filter phone numbers by number type.
     * @param "FilterFeatures" (optional.Interface of []string) -  Filter if the phone number should be used for voice, fax, mms, sms, emergency.
     * @param "FilterLimit" (optional.Int32) -  Limits the number of results.
     * @param "FilterBestEffort" (optional.Bool) -  Filter to determine if best effort results should be included.
     * @param "FilterQuickship" (optional.Bool) -  Filter to exclude phone numbers that need additional time after to purchase to receive phone calls.
     * @param "FilterReservable" (optional.Bool) -  Filter to exclude phone numbers that cannot be reserved before purchase.
     * @param "FilterExcludeHeldNumbers" (optional.Bool) -  Filter to exclude phone numbers that are currently on hold for your account.
@return ListAvailablePhoneNumbersResponse
*/

type NumberSearchApiListAvailablePhoneNumbersOpts struct {
    FilterPhoneNumberStartsWith optional.String
    FilterPhoneNumberEndsWith optional.String
    FilterPhoneNumberContains optional.String
    FilterLocality optional.String
    FilterAdministrativeArea optional.String
    FilterCountryCode optional.String
    FilterNationalDestinationCode optional.String
    FilterRateCenter optional.String
    FilterNumberType optional.String
    FilterFeatures optional.Interface
    FilterLimit optional.Int32
    FilterBestEffort optional.Bool
    FilterQuickship optional.Bool
    FilterReservable optional.Bool
    FilterExcludeHeldNumbers optional.Bool
}

func (a *NumberSearchApiService) ListAvailablePhoneNumbers(ctx context.Context, localVarOptionals *NumberSearchApiListAvailablePhoneNumbersOpts) (ListAvailablePhoneNumbersResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ListAvailablePhoneNumbersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/available_phone_numbers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.FilterPhoneNumberStartsWith.IsSet() {
		localVarQueryParams.Add("filter[phone_number][starts_with]", parameterToString(localVarOptionals.FilterPhoneNumberStartsWith.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterPhoneNumberEndsWith.IsSet() {
		localVarQueryParams.Add("filter[phone_number][ends_with]", parameterToString(localVarOptionals.FilterPhoneNumberEndsWith.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterPhoneNumberContains.IsSet() {
		localVarQueryParams.Add("filter[phone_number][contains]", parameterToString(localVarOptionals.FilterPhoneNumberContains.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterLocality.IsSet() {
		localVarQueryParams.Add("filter[locality]", parameterToString(localVarOptionals.FilterLocality.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterAdministrativeArea.IsSet() {
		localVarQueryParams.Add("filter[administrative_area]", parameterToString(localVarOptionals.FilterAdministrativeArea.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterCountryCode.IsSet() {
		localVarQueryParams.Add("filter[country_code]", parameterToString(localVarOptionals.FilterCountryCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterNationalDestinationCode.IsSet() {
		localVarQueryParams.Add("filter[national_destination_code]", parameterToString(localVarOptionals.FilterNationalDestinationCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterRateCenter.IsSet() {
		localVarQueryParams.Add("filter[rate_center]", parameterToString(localVarOptionals.FilterRateCenter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterNumberType.IsSet() {
		localVarQueryParams.Add("filter[number_type]", parameterToString(localVarOptionals.FilterNumberType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterFeatures.IsSet() {
		localVarQueryParams.Add("filter[features]", parameterToString(localVarOptionals.FilterFeatures.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.FilterLimit.IsSet() {
		localVarQueryParams.Add("filter[limit]", parameterToString(localVarOptionals.FilterLimit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterBestEffort.IsSet() {
		localVarQueryParams.Add("filter[best_effort]", parameterToString(localVarOptionals.FilterBestEffort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterQuickship.IsSet() {
		localVarQueryParams.Add("filter[quickship]", parameterToString(localVarOptionals.FilterQuickship.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterReservable.IsSet() {
		localVarQueryParams.Add("filter[reservable]", parameterToString(localVarOptionals.FilterReservable.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilterExcludeHeldNumbers.IsSet() {
		localVarQueryParams.Add("filter[exclude_held_numbers]", parameterToString(localVarOptionals.FilterExcludeHeldNumbers.Value(), ""))
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
			var v ListAvailablePhoneNumbersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v Errors
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
