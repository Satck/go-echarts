package charts

import (
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

type TreeMap struct {
	BaseActions
	BaseConfiguration
}

func (TreeMap) Type() string {
	return types.ChartTreeMap
}
func NewTreeMap() *TreeMap {
	c := &TreeMap{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *TreeMap) SetGlobalOptions(options ...GlobalOpts) *TreeMap {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

func (c *TreeMap) SetDispatchActions(actions ...GlobalActions) *TreeMap {
	c.BaseActions.setBaseGloActions(actions...)
	return c
}

func (c *TreeMap) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
