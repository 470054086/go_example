package registry
// 定义服务发现接口
type Registry interface {
	Register(option RegisterOption, provider ...Provider)
	Unregister(option RegisterOption, provider ...Provider)
	GetServiceList() []Provider
	Watch() Watcher
	Unwatch(watcher Watcher)
}

// 服务的唯一key 类型 com.douban.rpc.server
type RegisterOption struct {
	AppKey string
}

// 监听的对象
type Watcher interface {
	// 获取下一个地址
	Next() (*Event, error)
	Close()
}

// 定义操作的方法
type EventAction byte
const (
	Create EventAction = iota
	Update
	Delete
)
// 定义事件
type Event struct {
	Action    EventAction
	AppKey    string
	Providers []Provider
}
// 一个服务
type Provider struct {
	ProviderKey string // Network+"@"+Addr
	Network     string
	Addr        string
	Meta        map[string]string
}

type Peer2PeerDiscovery struct {
	providers []Provider
}

func (p *Peer2PeerDiscovery) Register(option RegisterOption, providers ...Provider) {
	p.providers = providers
}

func (p *Peer2PeerDiscovery) Unregister(option RegisterOption, provider ...Provider) {
	p.providers = []Provider{}
}

func (p *Peer2PeerDiscovery) GetServiceList() []Provider {
	return p.providers
}

func (p *Peer2PeerDiscovery) Watch() Watcher {
	return nil
}

func (p *Peer2PeerDiscovery) Unwatch(watcher Watcher) {
	return
}

func (p *Peer2PeerDiscovery) WithProvider(provider Provider) *Peer2PeerDiscovery {
	p.providers = append(p.providers, provider)
	return p
}

func (p *Peer2PeerDiscovery) WithProviders(providers []Provider) *Peer2PeerDiscovery {
	for _, provider := range providers {
		p.providers = append(p.providers, provider)
	}
	return p
}

func NewPeer2PeerRegistry() *Peer2PeerDiscovery {
	r := &Peer2PeerDiscovery{}
	return r
}