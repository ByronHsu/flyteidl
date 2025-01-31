/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// SignalCondition represents a dependency on an signal.
type CoreSignalCondition struct {
	// A unique identifier for the requested signal.
	SignalId string `json:"signal_id,omitempty"`
	// A type denoting the required value type for this signal.
	Type_ *CoreLiteralType `json:"type,omitempty"`
	// The variable name for the signal value in this nodes outputs.
	OutputVariableName string `json:"output_variable_name,omitempty"`
}
