package concurrence_mode

func orChannelMode(){
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0: //递归终止条件一
			return nil
		case 1: //递归终止条件二
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
				case 2:
					select {
					case <-channels[0]:
					case <-channels[1]:
					}
				default:
					select {
					case <-channels[0]:
					case <-channels[1]:
					case <-channels[2]:
					case <-or(append(channels[3:], orDone)...): //递归
					}
				}
		}()
		return orDone
	}
}