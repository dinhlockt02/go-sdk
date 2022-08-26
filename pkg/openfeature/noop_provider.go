package openfeature

// NoopProvider implements the FeatureProvider interface and provides functions for evaluating flags
type NoopProvider struct{}

// Metadata returns the metadata of the provider
func (e NoopProvider) Metadata() Metadata {
	return Metadata{Name: "NoopProvider"}
}

// BooleanEvaluation returns a boolean flag.
func (e NoopProvider) BooleanEvaluation(flag string, defaultValue bool, evalCtx EvaluationContext) BoolResolutionDetail {
	return BoolResolutionDetail{
		Value: defaultValue,
		ResolutionDetail: ResolutionDetail{
			Variant: "default-variant",
			Reason:  DEFAULT,
		},
	}
}

// StringEvaluation returns a string flag.
func (e NoopProvider) StringEvaluation(flag string, defaultValue string, evalCtx EvaluationContext) StringResolutionDetail {
	return StringResolutionDetail{
		Value: defaultValue,
		ResolutionDetail: ResolutionDetail{
			Variant: "default-variant",
			Reason:  DEFAULT,
		},
	}
}

// FloatEvaluation returns a float flag.
func (e NoopProvider) FloatEvaluation(flag string, defaultValue float64, evalCtx EvaluationContext) FloatResolutionDetail {
	return FloatResolutionDetail{
		Value: defaultValue,
		ResolutionDetail: ResolutionDetail{
			Variant: "default-variant",
			Reason:  DEFAULT,
		},
	}
}

// IntEvaluation returns an int flag.
func (e NoopProvider) IntEvaluation(flag string, defaultValue int64, evalCtx EvaluationContext) IntResolutionDetail {
	return IntResolutionDetail{
		Value: defaultValue,
		ResolutionDetail: ResolutionDetail{
			Variant: "default-variant",
			Reason:  DEFAULT,
		},
	}
}

// ObjectEvaluation returns an object flag
func (e NoopProvider) ObjectEvaluation(flag string, defaultValue interface{}, evalCtx EvaluationContext) ResolutionDetail {
	return ResolutionDetail{
		Value:   defaultValue,
		Variant: "default-variant",
		Reason:  DEFAULT,
	}
}

// Hooks returns hooks
func (e NoopProvider) Hooks() []Hook {
	return []Hook{}
}
