package day19

import "slices"

type PartDimension struct {
	minValue       int
	maxValue       int
	excludedValues []int
}

func constrainSlice(values []int, minValue, maxValue int) {
	slices.DeleteFunc(values, func(value int) bool {
		return value < minValue || value > maxValue
	})
}

func (pd PartDimension) Cardinality() int {
	return pd.maxValue - pd.minValue + 1 - len(pd.excludedValues)
}

func (pd PartDimension) MustBeMoreThan(floor int) PartDimension {
	pd.minValue = max(floor+1, pd.minValue)
	pd.excludedValues = slices.Clone(pd.excludedValues)
	constrainSlice(pd.excludedValues, pd.minValue, pd.maxValue)
	return pd
}

func (pd PartDimension) MustBeLessThan(ceil int) PartDimension {
	pd.maxValue = min(ceil-1, pd.maxValue)
	pd.excludedValues = slices.Clone(pd.excludedValues)
	constrainSlice(pd.excludedValues, pd.minValue, pd.maxValue)
	return pd
}

func (pd PartDimension) MustBe(value int) PartDimension {
	pd.minValue = value
	pd.maxValue = value
	pd.excludedValues = slices.Clone(pd.excludedValues)
	constrainSlice(pd.excludedValues, pd.minValue, pd.maxValue)
	return pd
}

func (pd PartDimension) MustNotBe(value int) PartDimension {
	pd.excludedValues = slices.Clone(pd.excludedValues)
	pd.excludedValues = append(pd.excludedValues, value)
	constrainSlice(pd.excludedValues, pd.minValue, pd.maxValue)
	return pd
}
