/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/13 上午10:04
 * @note: channel counter
 */

package counter

type Counter struct {
	queue  chan func()
	number uint64
}

func NewCounter() *Counter {
	counter := &Counter{
		queue:  make(chan func(), 64),
		number: 0,
	}
	go func(counter *Counter) {
		for f := range counter.queue {
			f()
		}
	}(counter)
	return counter
}

func (c *Counter) Add(num uint64) {
	c.queue <- func() {
		c.number += num
	}
}

func (c *Counter) Read() uint64 {
	ret := make(chan uint64)
	c.queue <- func() {
		ret <- c.number
		close(ret)
	}
	return <-ret
}
