package function

import (
	"context"
	"log/slog"

	"github.com/aceld/kis-flow/kis"
)

type KisFunctionV struct {
	BaseFunction
}

func NewKisFunctionV() kis.Function {
	f := new(KisFunctionV)

	// 初始化metaData
	f.metaData = make(map[string]interface{})

	return f
}

func (f *KisFunctionV) Call(ctx context.Context, flow kis.Flow) error {
	slog.Debug("KisFunctionV", "flow", flow)

	// 通过KisPool 路由到具体的执行计算Function中
	if err := kis.Pool().CallFunction(ctx, f.Config.FName, flow); err != nil {
		slog.ErrorContext(ctx, "Function Called Error", "err", err)
		return err
	}

	return nil
}
