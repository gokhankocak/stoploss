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
    Implements (MFI) Money Flow Index
*/

package stoploss

import(
  "math"
)

type MoneyFlow struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewMoneyFlow() *MoneyFlow {
  mf := new(MoneyFlow)
  mf.DataSet = DataSet
  return mf
}

// Calculate Money Flow Index
func (mf *MoneyFlow) MFI( Period int ) float64 {

	var TypicalPrice []float64
	var RawMoneyFlow []float64
	var PositiveFlow float64
	var NegativeFlow float64
	var Index int
	var Result float64

	LastIndex := len(mf.DataSet) - 1

	if LastIndex - Period - 1 < 0 { return math.NaN() }

	TypicalPrice = make([]float64, Period)
	RawMoneyFlow = make([]float64, Period)

	Index = Period - 1
	for k := LastIndex; k > LastIndex - Period; k-- {

		TypicalPrice[Index] = (mf.DataSet[k].High + mf.DataSet[k].Low + mf.DataSet[k].Close) / 3.0
		RawMoneyFlow[Index] = TypicalPrice[Index] * mf.DataSet[k].Volume
		Index--
	}

	PositiveFlow = 0.0
	NegativeFlow = 0.0
	for k := 1; k < Period - 1; k++ {

		if TypicalPrice[k] > TypicalPrice[k - 1] {

			PositiveFlow += RawMoneyFlow[k]
		} else if TypicalPrice[k] < TypicalPrice[k - 1] {

			NegativeFlow += RawMoneyFlow[k]
		}
	}

	Result = (100.0 - 100.0 / (1.0 + PositiveFlow / NegativeFlow))

	return Result
}
