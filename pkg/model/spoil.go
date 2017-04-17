package model

import "runtime"

// Spoiler consumes CPU
type Spoiler struct {
	stop chan bool
}

var spoiler *Spoiler

// GetSpoiler returns application info struct
func GetSpoiler() *Spoiler {
	if spoiler == nil {
		spoiler = &Spoiler{}
	}
	return spoiler
}

// Run starts consuming CPU
func (s *Spoiler) Run() {
	if s.stop == nil {
		s.spoil()
	}
}

// Stop consuming CPU
func (s *Spoiler) Stop() {
	if s.stop != nil {
		s.stop <- true
	}
}

func (s *Spoiler) spoil() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)
	s.stop = make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case <-s.stop:
					s.stop = nil
					return
				default:
				}
			}
		}()
	}

}
