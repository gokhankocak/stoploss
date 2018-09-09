package stoploss

import (
)

type PivotPoints struct {
	DataSet Series
	DebugMode bool
}

func (DataSet Series) NewPivotPoints() *PivotPoints {
	pp := new(PivotPoints)
	pp.DataSet = DataSet
	return pp
}

type PivotPointsResult struct {
	Pivot float64
	SupLevels [3]float64
	ResLevels [3]float64
}

// https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/pivot-points-resistance-support

func (pp *PivotPoints) Standard() (Result *PivotPointsResult) {

	LastIndex := len(pp.DataSet) - 1
	Result.Pivot = (pp.DataSet[LastIndex].High + pp.DataSet[LastIndex].Low + pp.DataSet[LastIndex].Close) / 3.0

	Result.SupLevels[0] = 2.0 * Result.Pivot - pp.DataSet[LastIndex].High
	Result.SupLevels[1] = Result.Pivot - (pp.DataSet[LastIndex].High - pp.DataSet[LastIndex].Low)

	Result.ResLevels[0] = 2.0 * Result.Pivot - pp.DataSet[LastIndex].Low
	Result.ResLevels[1] = Result.Pivot + (pp.DataSet[LastIndex].High - pp.DataSet[LastIndex].Low)

	return Result
}

func (pp *PivotPoints) Fibonacci() (Result *PivotPointsResult) {

	LastIndex := len(pp.DataSet) - 1
	Result.Pivot = (pp.DataSet[LastIndex].High + pp.DataSet[LastIndex].Low + pp.DataSet[LastIndex].Close) / 3.0
	Delta := pp.DataSet[LastIndex].High - pp.DataSet[LastIndex].Low

	Result.SupLevels[0] = Result.Pivot - 0.382 * Delta
	Result.SupLevels[1] = Result.Pivot - 0.618 * Delta
	Result.SupLevels[2] = Result.Pivot - Delta

	Result.ResLevels[0] = Result.Pivot + 0.382 * Delta
	Result.ResLevels[1] = Result.Pivot + 0.618 * Delta
	Result.ResLevels[2] = Result.Pivot + Delta

	return Result
}
