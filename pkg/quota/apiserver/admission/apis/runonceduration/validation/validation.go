package validation

import (
	"github.com/openshift/origin/pkg/quota/apiserver/admission/apis/runonceduration"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateRunOnceDurationConfig validates the RunOnceDuration plugin configuration
func ValidateRunOnceDurationConfig(config *runonceduration.RunOnceDurationConfig) field.ErrorList {
	allErrs := field.ErrorList{}
	if config == nil || config.ActiveDeadlineSecondsLimit == nil {
		return allErrs
	}
	if *config.ActiveDeadlineSecondsLimit <= 0 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("activeDeadlineSecondsOverride"), config.ActiveDeadlineSecondsLimit, "must be greater than 0"))
	}
	return allErrs
}
