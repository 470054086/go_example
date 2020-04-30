package Subscribe

var d map[string][]int

type F func(i int)
type subscribe struct {
	topic     map[string]int
	queue     map[string]map[string][]int
	topicChan map[string]chan int
	queueChan chan map[string]map[string]int
}

var G_SUB *subscribe

func NewSubscribe() *subscribe {
	G_SUB = &subscribe{
		topic:     make(map[string]int),
		queue:     make(map[string]map[string][]int),
		topicChan: make(map[string]chan int),
		queueChan: make(chan map[string]map[string]int),
	}
	return G_SUB
}
func (s *subscribe) Create(topic string) {
	if _, ok := s.topic[topic]; ok {
		panic("当前top已经创建了")
	}
	s.topic[topic] = 0
	s.queue[topic] = make(map[string][]int)
}

func (s *subscribe) AddQueue(topic, queue string, data int) {
	s.queue[topic][queue] = append(s.queue[topic][queue], data)
}

func (s *subscribe) SendQueue(topic, queue string,f F) {
	for _, val := range s.queue[topic][queue] {
		go func(val int) {
			f(val)
		}(val)
	}
}
