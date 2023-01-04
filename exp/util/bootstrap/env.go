package bootstrap

type EnvConfig struct {
	Set   map[string]string
	Unset []string

	Files []string
}

func EnvBootstrap(cfg EnvConfig) func() error {
	return nil
}
