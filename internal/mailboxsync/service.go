package mailboxsync

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
)

type Service struct {
	Syncs []Sync
}

func NewService() Service {
	return Service{}
}

func (s Service) SyncMailboxes(syncs []Sync) {
	concurrentSyncs := 3
	for _, sync := range syncs {
		eg, ctx := errgroup.WithContext(context.Background())
		for i, mailbox := range sync.Mailboxes {
			var args = []string{
				"--host1", sync.Src.getHost(),
				"--port1", sync.Src.getPort(),
				"--user1", mailbox.GetSrcUser(),
				"--password1", mailbox.GetSrcPassword(),
				"--host2", sync.Dst.getHost(),
				"--port2", sync.Dst.getPort(),
				"--user2", mailbox.GetDstUser(),
				"--password2", mailbox.GetDstPassword(),
				"--logdir", "var/log",
			}

			cmd := exec.CommandContext(ctx, "imapsync", args...)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			eg.Go(func() error {
				return cmd.Run()
			})

			if i+1%concurrentSyncs == 0 {
				eg.Wait()
			}
		}

		eg.Wait()
	}
}
