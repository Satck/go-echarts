package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Overlaper interface {
	overlap() MultiSeries
}

type XYAxis struct {
	XAxisList []opts.XAxis `json:"xaxis"`
	YAxisList []opts.YAxis `json:"yaxis"`
}

func (xy *XYAxis) initXYAxis() {
	xy.XAxisList = append(xy.XAxisList, opts.XAxis{})
	xy.YAxisList = append(xy.YAxisList, opts.YAxis{})
}

func (xy *XYAxis) ExtendXAxis(xAxis ...opts.XAxis) {
	xy.XAxisList = append(xy.XAxisList, xAxis...)
}

func (xy *XYAxis) ExtendYAXis(yAxis ...opts.YAxis) {
	xy.YAxisList = append(xy.YAxisList, yAxis...)
}

//设置x轴

func WithXAxisOpts(opt opts.XAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.XAxisList[index[i]] = opt
		}
	}
}

// 设置y轴
func WithYAxisOpts(opt opts.YAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.YAxisList[index[i]] = opt
		}
	}
}

type RectConfiguration struct {
	BaseConfiguration
	BaseActions
}

func (rect *RectConfiguration) setRectGlobalOptions(options ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(options...)
}
func (rect *RectConfiguration) setRectGlobalActions(options ...GlobalActions) {
	rect.BaseActions.setBaseGlobalActions(options...)
}

type RectChart struct {
	RectConfiguration

	xAxisData interface{}
}

func (rc *RectChart) overlap() MultiSeries {
	return rc.MultiSeries
}

func (rc *RectChart) SetGlobalOptions(options ...GlobalOpts) *RectChart {
	rc.RectConfiguration.setRectGlobalOptions(options...)
	return rc
}
func (rc *RectChart) SetDispatchActions(options ...GlobalActions) *RectChart {
	rc.RectConfiguration.setRectGlobalActions(options...)
	return rc

}

//Overlap重叠将多个图表组成一个画布。

//它只适用于某些直角坐标系的图表。

//支持的图表：条形图/BoxPlot/Line/Stratter/EffetStratter/Kline/HeatMap

func (rc *RectChart) Overlap(a ...Overlaper) {
	for i := 0; i < len(a); i++ {
		rc.MultiSeries = append(rc.MultiSeries, a[i].overlap()...)
	}
}

func (rc *RectChart) Validate() {
	//确保XAxisOpts的X轴数据不会被清除
	rc.XAxisList[0].Data = rc.xAxisData
	//确保Y轴的标签正确显示
	for i := 0; i < len(rc.YAxisList); i++ {
		rc.YAxisList[i].AxisLabel.Show = true
	}
	rc.Assets.Validate(rc.AssetsHost)
}
