/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Weight struct {

	// bhive_id to identify the data source. Might become empty in a future iteration as it's redundant with the query parameter.
	BhiveId string `json:"bhive_id,omitempty"`

	// actual measurement of weight
	Weight float32 `json:"weight,omitempty"`

	// timestamp of the measurement on one second precision
	Epoch int64 `json:"epoch,omitempty"`
}

// AssertWeightRequired checks if the required fields are not zero-ed
func AssertWeightRequired(obj Weight) error {
	return nil
}

// AssertRecurseWeightRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Weight (e.g. [][]Weight), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseWeightRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aWeight, ok := obj.(Weight)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertWeightRequired(aWeight)
	})
}