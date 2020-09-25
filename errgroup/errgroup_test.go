package errgroup

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

func TestErrgroup(t *testing.T) {
	var count int64 = 1
	countBackup := count
	eg := WithContext(context.Background())

	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return errors.New("error ,reset count")
	})
	// wait 协程 Done
	if err := eg.Wait(); err != nil {
		// do some thing
		count = countBackup
		xlog.Error(err)
		//return
	}
	xlog.Debug(count)
}
