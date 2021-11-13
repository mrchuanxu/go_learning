package go_design_pattern_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

// observer mode
// func params is 统一的参数
type observer func(ctx context.Context, params []byte) error

type BeObserver struct {
	Obs    []observer
	params paramsStruct
}

// 建造者 register
func (b *BeObserver) SetObserver(ctx context.Context, ob observer) *BeObserver {
	b.Obs = append(b.Obs, ob)
	return b
}

type paramsStruct struct {
	A string
	B string
}

func (b *BeObserver) RunObs(ctx context.Context) {
	params, _ := json.Marshal(b.params)
	// 协程可以管理起来
	for i := range b.Obs {
		go b.Obs[i](ctx, params)
	}
}

func Test_observer(t *testing.T) {
	beObs := &BeObserver{}
	beObs.Obs = make([]observer, 0)
	beObs.params.A = "A"
	beObs.params.B = "B"
	// 注册观察者
	beObs.SetObserver(context.Background(), func(ctx context.Context, params []byte) error {
		par := &paramsStruct{}
		_ = json.Unmarshal(params, par)
		fmt.Println(par.A)
		return nil
	})

	beObs.SetObserver(context.Background(), func(ctx context.Context, params []byte) error {
		par := &paramsStruct{}
		_ = json.Unmarshal(params, par)
		fmt.Println(par.B)
		return nil
	})
	beObs.RunObs(context.Background())

}
