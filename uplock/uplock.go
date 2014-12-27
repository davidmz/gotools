package uplock

import "sync"

type Mutex struct{ sync.RWMutex }

type Locker interface {
	sync.Locker
	RLocker() sync.Locker
}

func Do(lk Locker, rAction func() bool, wAction func()) {

	ok := func() bool {
		rl := lk.RLocker()
		rl.Lock()
		defer rl.Unlock()
		return rAction()
	}()

	if !ok {
		return
	}

	func() {
		lk.Lock()
		defer lk.Unlock()
		if rAction() {
			wAction()
		}
	}()
}

func (m *Mutex) Do(rAction func() bool, wAction func()) { Do(m, rAction, wAction) }
