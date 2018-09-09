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
    Implements Fibonacci NUmbers and Analysis
*/

package stoploss

import(
)

type Fibonacci struct {
  DataSet Series
  DebugMode bool
}

func (DataSet Series) NewFibonacci() *Fibonacci {

  fib := new(Fibonacci)
  fib.DataSet = DataSet
  return fib
}

func FibonacciNumbers( Count int ) []int {

  Numbers := make([]int, Count)
  Numbers[0] = 1
  Numbers[1] = 2

  First := 1
  Second := 2

  for k := 2; k < Count; k++ {
    Numbers[k] = First + Second
    First = Second
    Second = Numbers[k]
  }

  return Numbers
}

func (f *Fibonacci) IsFibonacciTimeSinceLowestClose() (RetVal bool, FibNumber int) {

  LastIndex := len(f.DataSet) - 1
  Index, _ := f.DataSet.LowestClose()

  First := 1
  Second := 2
  for {

    FibNumber = First + Second
    if Index + FibNumber == LastIndex { return true, FibNumber }
    if Index + FibNumber > LastIndex { break }
    First = Second
    Second = FibNumber
  }

  return false, -1
}

func (f *Fibonacci) IsFibonacciTimeSinceHighestClose() (RetVal bool, FibNumber int) {

  LastIndex := len(f.DataSet) - 1
  Index, _ := f.DataSet.HighestClose()

  First := 1
  Second := 2
  for {

    FibNumber = First + Second
    if Index + FibNumber == LastIndex { return true, FibNumber }
    if Index + FibNumber > LastIndex { break }
    First = Second
    Second = FibNumber
  }

  return false, -1
}

//
// TODO: Functional test, is the algorithm correct?
//

func (f *Fibonacci) FibonacciRetracement() (RetLevels [6]float64) {

  var Lowest, Highest, Delta float64
  var Levels = [4]float64{ 23.6, 38.2, 50.0, 61.8 }

  _, Lowest = f.DataSet.LowestClose()
  _, Highest = f.DataSet.HighestClose()
  Delta = Highest - Lowest

  RetLevels[0] = Highest
  for k := 1; k < len(RetLevels) - 1; k++ {
    RetLevels[k] = Delta * Levels[k - 1]
  }
  RetLevels[len(RetLevels) - 1] = Lowest

  return RetLevels
}
