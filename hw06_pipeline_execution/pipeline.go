package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		ch := make(chan interface{})
		close(ch)
		return ch
	}

	for _, stage := range stages {
		stageChan := make(Bi)

		go func(stageChan Bi, read Out) {
			defer close(stageChan)

			for {
				select {
				case <-done:
					return
				case v, ok := <-read:
					if !ok {
						return
					}
					stageChan <- v
				}
			}
		}(stageChan, in)

		in = stage(stageChan)
	}

	return in
}
