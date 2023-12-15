package sync

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
)

type Service struct {
	ConcurrentSyncs int
	Syncs           []Sync
}

func NewService(concurrentSyncs int) Service {
	return Service{ConcurrentSyncs: concurrentSyncs}
}

func (s Service) SyncMailboxes(syncs []Sync) {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, sync := range syncs {
		for i, mailbox := range sync.GetActiveMailboxes() {
			var args = []string{
				"--host1", sync.Src.GetHost(),
				"--port1", sync.Src.GetPort(),
				"--user1", mailbox.GetSrcUser(),
				"--host2", sync.Dst.GetHost(),
				"--port2", sync.Dst.GetPort(),
				"--user2", mailbox.GetDstUser(),
				"--logdir", "var/log",
				"--tmpdir", "var/tmp",
			}

			cmd := exec.CommandContext(ctx, "imapsync", args...)
			cmd.Env = append(cmd.Env, fmt.Sprintf("IMAPSYNC_PASSWORD1=%s", mailbox.GetSrcPassword()), fmt.Sprintf("IMAPSYNC_PASSWORD2=%s", mailbox.GetDstPassword()))

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			eg.Go(func() error {
				return cmd.Run()
			})

			if i+1%s.ConcurrentSyncs == 0 {
				eg.Wait()
			}
		}

		eg.Wait()
	}
}
