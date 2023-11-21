// This file creates an alias for the types borrowed from `consensys/gnark`.
// In the future these types can be replaced with custom structures if they
// change across different `consensys/gnark` versions.
package extractor

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/schema"
)

type ExtractorSchema = schema.Schema
type ExtractorField = schema.Field
type ExtractorCircuit = frontend.Circuit
type ExtractorVariable = frontend.Variable
