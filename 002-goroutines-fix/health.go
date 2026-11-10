package health

import "context"

// CheckHealth runs probe and returns its error, or ctx.Err() if the
// context is done before the probe finishes.
func CheckHealth(ctx context.Context, probe func() error) error {
	result := make(chan error)

	go func() {
		result <- probe()
	}()

	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
