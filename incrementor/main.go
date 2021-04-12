// Package incrementor provides type Incrementor
package incrementor

import (
	"fmt"
	"math"
	"sync"
)

// Incrementor provides simple incrementer
type Incrementor struct {
	sync.RWMutex
	maxValue     int64
	currentValue int64
}

// NewIncrementor возвращает указатель на структуру Incrementor
func NewIncrementor() *Incrementor {
	return &Incrementor{
		maxValue:     math.MaxInt64,
		currentValue: 0,
	}
}

// GetNumber возвращает текущее число. В самом начале это ноль.
func (inc *Incrementor) GetNumber() int64 {
	inc.RLock()
	defer inc.RUnlock()
	return inc.currentValue
}

// IncrementNumber увеличивает текущее число на один. После каждого вызова этого
// метода getNumber() будет возвращать число на один больше.
func (inc *Incrementor) IncrementNumber() {
	inc.Lock()
	defer inc.Unlock()
	newValue := inc.currentValue + 1
	if newValue > inc.maxValue {
		inc.currentValue = 0
	} else {
		inc.currentValue = newValue
	}
}

// SetMaximumValue Устанавливает максимальное значение текущего числа.
// Хранимое число не может превышать установленное максимальное
// значение.
// Когда при вызове incrementNumber() текущее число достигает
// этого значения, оно обнуляется, т.е. getNumber() начинает
// снова возвращать ноль, и снова один после следующего
// вызова incrementNumber() и так далее.
// По умолчанию максимум -- максимальное значение int.
// Нельзя позволять установить тут число меньше нуля.
func (inc *Incrementor) SetMaximumValue(maxValue int64) error {
	if maxValue < 0 {
		return fmt.Errorf("Нельзя установить число меньше нуля.")
	}
	inc.Lock()
	inc.maxValue = maxValue
	inc.Unlock()
	return nil
}
