package synckit

import "github.com/soverenio/vanilla/throw"

type SignalChannel = <-chan struct{}

type ClosableSignalChannel = chan struct{}

func ClosedChannel() SignalChannel {
	return closedChan
}

var closedChan = func() SignalChannel {
	c := make(ClosableSignalChannel)
	close(c)
	return c
}()

func SafeClose(c ClosableSignalChannel) (err error) {
	defer func() {
		err = throw.R(recover(), err)
	}()
	close(c)
	return nil
}
