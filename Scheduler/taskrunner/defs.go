package taskrunner
//三种信号
const (
	READ_TO_DISPATCH = "a"   //派遣任务
	READ_TO_EXECUTE  = "e"   //处理任务
	CLOSE            = "c"   //退出
)

type SignChan chan string

type DataChan chan interface{}

type  fn func(da DataChan) error

func NewRunner (longLive bool,disPathch fn,execute fn,datasize int) *Runner{
	return &Runner{
		signChan:  make(chan string,1),
		dataChan:  make(chan interface{},1),
		dataSize:  datasize,
		longLive:  longLive,
		Dispathch: disPathch,
		execute:   execute,
		err:       make(chan string,1),
	}
}

type Runner struct {
	signChan SignChan
	dataChan DataChan
	dataSize int
	longLive bool
	Dispathch fn
	execute   fn
	err      SignChan
}

func (r *Runner)StartDispatch(){
	if !r.longLive {
		close(r.err)
		close(r.dataChan)
		close(r.signChan)
	}
	Loop:
		//生产者 消费者模型
	for {
		select {
		case sign := <-r.signChan:
			//第一次执行需要预支一个通知生产者的任务
			//往后是在消费者执行完任务以后再给生产者发送消息
			if sign == READ_TO_DISPATCH{
				err := r.Dispathch(r.dataChan)
				if err != nil {
					r.err <- CLOSE
				}else {
					r.signChan <- READ_TO_EXECUTE
				}
			}
			//生产者通知消费者开始执行任务
			if sign == READ_TO_EXECUTE{
				err := r.execute(r.dataChan)
				if err != nil {
					r.err <- CLOSE
				}else {
					r.signChan <- READ_TO_DISPATCH
				}

			}

		case err := <-r.err:
			if err == CLOSE{
				return
			}
		default:
			break Loop
		}
	}
}

func (r *Runner)StartAll(){
	r.signChan <- READ_TO_DISPATCH
	go r.StartDispatch()
}