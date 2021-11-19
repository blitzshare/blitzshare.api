package registry_test

import (
	"blitzshare.api/app/config"
	"blitzshare.api/app/dependencies"
	"blitzshare.api/app/model"
	"blitzshare.api/app/server/services/registry"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPeerRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registry test")
}

var _ = Describe("Test Registry", func() {
	Context("Given a registry", func() {
		var deps *dependencies.Dependencies
		BeforeSuite(func() {
			server := config.Server{Port: 323}
			settings := config.Settings{RedisUrl: "redis.svc.cluster.local"}
			config := config.Config{
				Server:   server,
				Settings: settings,
			}
			deps, _ = dependencies.NewDependencies(config)
		})
		It("should fail to return multiaddress for unregistrated peer", func() {
			fetchedMultiAddr, err := registry.GetPeerMultiAddr(deps, "glycogen-descanting-booing-crassness")
			Expect(fetchedMultiAddr).To(Equal(""))
			Expect(err.Error()).To(Equal("PeerNotFoundError"))
		})
		It("should return error for not found peers", func() {
			r := model.PeerRegistry{MultiAddr: "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR", OneTimePass: "clown-descanting-booing-crassness"}
			multiAddr, err := registry.RegisterPeer(deps, &r)
			Expect(multiAddr).To(Equal(r.MultiAddr))
			Expect(err).To(BeNil())
			fetchedMultiAddr, err := registry.GetPeerMultiAddr(deps, "neverseen-pass")
			Expect(fetchedMultiAddr).To(Equal(""))
			Expect(err.Error()).To(Equal("PeerNotFoundError"))
		})
		It("should register peer and return multiaddress", func() {
			r := model.PeerRegistry{MultiAddr: "/ip4/10.101.18.26/tcp/63785/p2p/12D3KooWPGR", OneTimePass: "boom-descanting-booing-crassness"}
			multiAddr, err := registry.RegisterPeer(deps, &r)
			Expect(err).To(BeNil())
			Expect(multiAddr).To(Equal(r.MultiAddr))
			fetchedMultiAddr, err := registry.GetPeerMultiAddr(deps, r.OneTimePass)
			Expect(fetchedMultiAddr).To(Equal(r.MultiAddr))
			Expect(err).To(BeNil())
		})
	})
})
