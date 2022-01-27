module fybrik.io/fybrik

go 1.16

require (
	emperror.dev/errors v0.7.0
	github.com/IBM/satcon-client-go v0.2.1-0.20211027144622-4f54f37377a3
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/apache/arrow/go/arrow v0.0.0-20210907151234-f40856a768f2
	github.com/aws/aws-sdk-go v1.40.37
	github.com/distribution/distribution v2.7.1+incompatible
	github.com/gdexlab/go-render v1.0.1
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/gin-gonic/gin v1.7.7
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-test/deep v1.0.7 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/hashicorp/vault/api v1.3.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/open-policy-agent/opa v0.34.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.0
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/vdemeester/k8s-pkg-credentialprovider v1.22.4
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
	helm.sh/helm/v3 v3.8.0
	k8s.io/api v0.23.1
	k8s.io/apiextensions-apiserver v0.23.1
	k8s.io/apimachinery v0.23.1
	k8s.io/cli-runtime v0.23.1
	k8s.io/client-go v0.23.1
	sigs.k8s.io/cli-utils v0.19.2
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/yaml v1.3.0
)

// This replace is for https://github.com/advisories/GHSA-w73w-5m7g-f7qc
replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible

replace helm.sh/helm/v3 v3.8.0 => github.com/fybrik/helm/v3 v3.8.0-fybrik-update
