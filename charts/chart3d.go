package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Chart3D struct {
	BaseConfiguration
}

// WithXAxis 3DOpts设置Chart3D实例的X轴。
func WithXAxis3D0pts(opt opts.XAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.XAxis3D = opt
	}
}

// WithXAxis3Opts设置Chart3D实例的Y轴。
func WithYAxis3DOpts(opt opts.YAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.YAxis3D = opt
	}
}

// WithZAxis3Opts设置Chart3D实例的Z轴。
func WithZAxis3D0pts(opt opts.ZAxis3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.ZAxis3D = opt
	}
}

// WithGrid3DOpts设置Chart3D实例的网格。
func WithGrid3D0pts(opt opts.Grid3D) GlobalOpts {
	return func(bc *BaseConfiguration) {
		bc.Grid3D = opt
	}
}

func (c *Chart3D) initChart3D() {
	c.JSAssets.Add("echarts-gl.min.js")
	c.has3DAxis = true
}

func (c *Chart3D) addSeries(chartType, name string, data []opts.Chart3DData, options ...SeriesOpts) {
	series := SingleSeries{
		Name:        name,
		Type:        chartType,
		Data:        data,
		CoordSystem: types.ChartCartesian3D,
	}
	series.ConfigureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
}

// SetGlobalOptions设置Chart3D实例的选项。
func (c *Chart3D) SetGlobalOptions(options ...GlobalOpts) *Chart3D {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate验证给定的配置。
func (c *Chart3D) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
