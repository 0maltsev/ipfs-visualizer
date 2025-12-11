package config

type Config struct {
    HTTP HTTPConfig `envPrefix:"HTTP_"`
    Kube KubeConfig `envPrefix:"KUBE_"`
}

type HTTPConfig struct {
    Host string `env:"HOST" envDefault:"0.0.0.0"`
    Port int    `env:"PORT" envDefault:"8080"`
}

type KubeConfig struct {
    InCluster   bool   `env:"IN_CLUSTER" envDefault:"false"`
    Kubeconfig string `env:"KUBECONFIG" envDefault:"~/.kube/config"`
}
