package validate

import (
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service"
)

func RecordListParams(resource *model.Resource, propertyPathMap map[string]bool, params service.RecordListParams) error {
	// check filters
	if err := Filters(propertyPathMap, params.Filters); err != nil {
		return err
	}

	// check query
	if params.Query != nil {
		if err := BooleanExpression(resource, propertyPathMap, params.Query); err != nil {
			return err
		}
	}

	// check sorting
	if params.Sorting != nil {
		for _, sort := range params.Sorting.Items {
			if !propertyPathMap["$."+sort.Property] {
				return errors.RecordValidationError.WithDetails(fmt.Sprintf("Sorting property %s is not a valid property path", sort.Property))
			}
		}
	}

	if params.Aggregation != nil {
		for _, agg := range params.Aggregation.Items {
			if !propertyPathMap["$."+agg.Property] {
				return errors.RecordValidationError.WithDetails(fmt.Sprintf("Aggregation property %s is not a valid property path", agg.Property))
			}

			if !NamePattern.MatchString(agg.Name) {
				return errors.RecordValidationError.WithDetails(fmt.Sprintf("Aggregation name %s should match pattern %s", agg.Name, NamePattern.String()))
			}
		}
	}

	return nil
}

func Filters(propertyPathMap map[string]bool, filters map[string]interface{}) error {
	for k := range filters {
		if !propertyPathMap["$."+k] {
			return errors.RecordValidationError.WithDetails(fmt.Sprintf("Filter %s is not a valid property path", k))
		}
	}

	return nil
}

func BooleanExpression(resource *model.Resource, propertyPathMap map[string]bool, exp *model.BooleanExpression) error {
	if exp.Expression != nil {
		if exp.GetAnd() != nil {
			for _, e := range exp.GetAnd().Expressions {
				if err := BooleanExpression(resource, propertyPathMap, e); err != nil {
					return err
				}
			}
		}

		if exp.GetOr() != nil {
			for _, e := range exp.GetOr().Expressions {
				if err := BooleanExpression(resource, propertyPathMap, e); err != nil {
					return err
				}
			}
		}

		if exp.GetNot() != nil {
			if err := BooleanExpression(resource, propertyPathMap, exp.GetNot()); err != nil {
				return err
			}
		}

		if exp.GetEqual() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetEqual()); err != nil {
				return err
			}
		}

		if exp.GetGreaterThan() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetGreaterThan()); err != nil {
				return err
			}
		}

		if exp.GetLessThan() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetLessThan()); err != nil {
				return err
			}
		}

		if exp.GetGreaterThanOrEqual() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetGreaterThanOrEqual()); err != nil {
				return err
			}
		}

		if exp.GetLessThanOrEqual() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetLessThanOrEqual()); err != nil {
				return err
			}
		}

		if exp.GetIn() != nil {
			if err := PairExpression(resource, propertyPathMap, exp.GetIn()); err != nil {
				return err
			}
		}

		if exp.GetFilters() != nil {
			for k := range exp.GetFilters() {
				if !propertyPathMap["$."+k] {
					return errors.RecordValidationError.WithDetails(fmt.Sprintf("Filter %s is not a valid property path", k))
				}
			}
		}

		if exp.GetRegexMatch() != nil {
			if err := RegexMatchExpression(resource, propertyPathMap, exp.GetRegexMatch()); err != nil {
				return err
			}
		}
	}

	if exp.Filters != nil {
		for k := range exp.Filters {
			if !propertyPathMap["$."+k] {
				return errors.RecordValidationError.WithDetails(fmt.Sprintf("Filter %s is not a valid property path", k))
			}
		}
	}

	return nil
}

func PairExpression(resource *model.Resource, pathMap map[string]bool, equal *model.PairExpression) error {
	if equal.Left != nil {
		if err := Expression(resource, pathMap, equal.Left); err != nil {
			return err
		}
	} else {
		return errors.RecordValidationError.WithDetails("PairExpression left is nil")
	}

	if equal.Right != nil {
		if err := Expression(resource, pathMap, equal.Right); err != nil {
			return err
		}
	} else {
		return errors.RecordValidationError.WithDetails("PairExpression right is nil")
	}

	return nil
}

func Expression(resource *model.Resource, pathMap map[string]bool, exp *model.Expression) error {
	if exp.GetProperty() != "" {
		if !pathMap["$."+exp.GetProperty()] {
			return errors.RecordValidationError.WithDetails(fmt.Sprintf("Expression property %s is not a valid property path", exp.GetProperty()))
		}
	} else if exp.GetValue() != nil {
		return nil
	} else {
		return errors.RecordValidationError.WithDetails("either property or value must be set in expression")
	}

	return nil
}

func RegexMatchExpression(resource *model.Resource, pathMap map[string]bool, regex *model.RegexMatchExpression) error {
	return errors.RecordValidationError.WithDetails("RegexMatchExpression is not implemented")
}
