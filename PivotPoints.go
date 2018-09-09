/*
Copyright 2018 www.gokhankocak.com gokhan.kocak@mail.ru

The MIT License
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

/*
		Implements Pivot Points
		Standard Pivot Points
		Fibonacci Pivot Points
*/

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
