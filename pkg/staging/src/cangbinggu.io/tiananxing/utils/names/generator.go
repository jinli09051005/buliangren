package names

import (
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	k8snames "k8s.io/apiserver/pkg/storage/names"
	"strings"
)

type generator struct {
}

var Generator k8snames.NameGenerator = generator{}

func (generator) GenerateName(base string) string {
	if strings.HasSuffix(base, "-") {
		return base + utilrand.String(8)
	}
	return base + "-" + utilrand.String(8)
}
