package shared_testutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"

	datatransfer "github.com/filecoin-project/go-data-transfer/v2"

	"github.com/brossetti1/go-fil-markets/shared"
)

// StartAndWaitable is any interface that can be started up and will be asynchronously ready later
type StartAndWaitable interface {
	Start(ctx context.Context) error
	OnReady(shared.ReadyFunc)
}

// StartAndWaitForReady is a utility function to start a module and verify it reaches the ready state
func StartAndWaitForReady(ctx context.Context, t *testing.T, startAndWaitable StartAndWaitable) {
	ready := make(chan error, 1)
	startAndWaitable.OnReady(func(err error) {
		ready <- err
	})
	require.NoError(t, startAndWaitable.Start(ctx))
	select {
	case <-ctx.Done():
		t.Fatal("did not finish starting up module")
	case err := <-ready:
		require.NoError(t, err)
	}
}

// StartAndWaitForReadyDT is a utility function to start a go-data-transfer and verify it reaches the ready state
func StartAndWaitForReadyDT(ctx context.Context, t *testing.T, startAndWaitable datatransfer.Manager) {
	ready := make(chan error, 1)
	startAndWaitable.OnReady(func(err error) {
		ready <- err
	})
	require.NoError(t, startAndWaitable.Start(ctx))
	select {
	case <-ctx.Done():
		t.Fatal("did not finish starting up module")
	case err := <-ready:
		require.NoError(t, err)
	}
}
