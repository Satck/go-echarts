package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type WordCloud struct {
	BaseActions
	BaseConfiguration
}

func (WordCloud) Type() string {
	return types.ChartWordCloud
}

var wcTextColor = `function () {
	return 'rgb(' + [
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160)].join(',') + ')';
}`

func NewWordCloud() *WordCloud {
	c := &WordCloud{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.JSAssets.Add("echarts-wordcloud.min.js")
	return c
}

func (c *WordCloud) AddSeries(name string, data []opts.WordCloudData, options ...SeriesOpts) *WordCloud {
	series := SingleSeries{Name: name, Type: types.ChartWordCloud, Data: data}
	series.ConfigureSeriesOpts(options...)

	////设置WordCloud图表的默认随机颜色
	if series.TextStyle == nil {
		series.TextStyle = &opts.TextStyle{Normal: &opts.TextStyle{}}
	}
	if series.TextStyle.Normal.Color == "" {
		series.TextStyle.Normal.Color = opts.FuncOpts(wcTextColor)
	}

	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *WordCloud) SetGlobalOptions(options ...GlobalOpts) *WordCloud {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *WordCloud) SetDispatchActions(actions ...GlobalActions) *WordCloud {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c

}

func (c *WordCloud) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
