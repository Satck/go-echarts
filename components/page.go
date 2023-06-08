package components

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
)

// Page represents a page chart.
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	FillDefaultValues()
	Validate()
}

type Page struct {
	render.Renderer
	opts.Initialization
	opts.Assets

	Charts []interface{}
	Layout Layout
}

func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate)
	page.Layout = PageCenterLayout
	return page
}

func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()

		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}

		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}
