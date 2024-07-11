package casbin

import (
	"archv1/internal/pkg/config"
	"github.com/casbin/casbin/v2"
	"log"
)

func NewEnforcer(cfg *config.Config) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatalf("failed to create new casbin enforcer: %v", err)
	}

	return enforcer
}
