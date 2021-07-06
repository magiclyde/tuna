/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午11:29
 * @note:
 */

package logger

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	DefaultLogger.Info("DefaultLogger test")

	l := NewLogger(zap.InfoLevel, "./output.log")

	l.Info("info test")
	l.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second))

	l.Sugar().Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	l.Sugar().Infof("sugar.Infof: %s", "xx")
}
