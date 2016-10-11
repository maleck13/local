package councillors

import (
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
)

//Service handles actions around Councillors
type Service struct {
	Config         *config.Config
	CouncillorRepo domain.CouncillorRepo
}
