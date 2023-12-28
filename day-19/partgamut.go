package day19

type PartGamut struct {
	x PartDimension
	m PartDimension
	a PartDimension
	s PartDimension
}

func NewPartGamut() PartGamut {
	return PartGamut{
		x: PartDimension{minValue: 1, maxValue: 4000},
		m: PartDimension{minValue: 1, maxValue: 4000},
		a: PartDimension{minValue: 1, maxValue: 4000},
		s: PartDimension{minValue: 1, maxValue: 4000},
	}
}

func (pg PartGamut) Cardinality() int {
	return pg.x.Cardinality() *
		pg.m.Cardinality() *
		pg.a.Cardinality() *
		pg.s.Cardinality()
}

func (pg *PartGamut) GetDimension(name rune) PartDimension {
	switch name {
	case 'x':
		return pg.x
	case 'm':
		return pg.m
	case 'a':
		return pg.a
	case 's':
		return pg.s
	default:
		panic("invalid dimension name")
	}
}

func (pg *PartGamut) SetDimension(name rune, dimension PartDimension) {
	switch name {
	case 'x':
		pg.x = dimension
	case 'm':
		pg.m = dimension
	case 'a':
		pg.a = dimension
	case 's':
		pg.s = dimension
	default:
		panic("invalid dimension name")
	}
}

func (pg PartGamut) ThatSatisfiesCondition(condition WorkflowCondition) PartGamut {
	dimension := pg.GetDimension(condition.property)
	switch condition.operator {
	case '>':
		dimension = dimension.MustBeMoreThan(condition.value)
	case '<':
		dimension = dimension.MustBeLessThan(condition.value)
	case '=':
		dimension = dimension.MustBe(condition.value)
	default:
		panic("invalid workflow operator")
	}
	pg.SetDimension(condition.property, dimension)
	return pg
}

func (pg PartGamut) ThatDoesNotSatisfyCondition(condition WorkflowCondition) PartGamut {
	dimension := pg.GetDimension(condition.property)
	switch condition.operator {
	case '>':
		dimension = dimension.MustBeLessThan(condition.value + 1)
	case '<':
		dimension = dimension.MustBeMoreThan(condition.value - 1)
	case '=':
		dimension = dimension.MustNotBe(condition.value)
	default:
		panic("invalid workflow operator")
	}
	pg.SetDimension(condition.property, dimension)
	return pg
}
