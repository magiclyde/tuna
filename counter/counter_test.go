/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/13 上午10:20
 * @note:
 */

package counter

import "testing"

func TestNewCounter(t *testing.T) {
	c := NewCounter()
	c.Add(1)
	if c.Read() != 1 {
		t.Error("sth went wrong")
	}
}
