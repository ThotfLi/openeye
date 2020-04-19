package taskrunner

import "time"

type ScheduledTask struct {
	t *time.Ticker
	r *Runner
}

func NewScheduledTask (t time.Duration,l bool,disp fn,exec fn,datasize int) *ScheduledTask{
	return &ScheduledTask{
		t: time.NewTicker(t*time.Second),
		r: NewRunner(l,disp,exec,datasize),
	}
}

func (st *ScheduledTask)StartWork(){
	for {
		select {
		case <- st.t.C:
			go st.r.StartAll()
		}
	}
}

func Start(){
	st := NewScheduledTask(3,true,VideoDeleteProducer,VideoDeleteConsumers,3)
	st.StartWork()
}