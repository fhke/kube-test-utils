package kind

import "sigs.k8s.io/kind/pkg/cluster"

func (n *NewKindClusterOpts) newKindCreateOpts() []cluster.CreateOption {
	ret := []cluster.CreateOption{}

	if n.KubeconfigPath != "" {
		ret = append(ret, cluster.CreateWithKubeconfigPath(n.KubeconfigPath))
	}

	if n.NodeImage != "" {
		ret = append(ret, cluster.CreateWithNodeImage(n.NodeImage))
	}

	return ret
}
