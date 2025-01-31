package mock

import (
	"context"
	"errors"
	"fmt"

	"github.com/celestiaorg/celestia-core/light/provider"
	"github.com/celestiaorg/celestia-core/types"
)

var errNoResp = errors.New("no response from provider")

type deadMock struct {
	id string
}

// NewDeadMock creates a mock provider that always errors. id is used in case of multiple providers.
func NewDeadMock(id string) provider.Provider {
	return &deadMock{id: id}
}

func (p *deadMock) String() string {
	return fmt.Sprintf("DeadMock-%s", p.id)
}

func (p *deadMock) LightBlock(_ context.Context, height int64) (*types.LightBlock, error) {
	return nil, errNoResp
}

func (p *deadMock) DASLightBlock(_ctx context.Context, height int64) (*types.LightBlock, error) {
	return nil, errNoResp
}

func (p *deadMock) ReportEvidence(_ context.Context, ev types.Evidence) error {
	return errNoResp
}
