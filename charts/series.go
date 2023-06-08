package charts

import "github.com/go-echarts/go-echarts/v2/opts"

type SingleSeries struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	// Rectangular charts
	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	// Bar
	BarGap         string `json:"barGap,omitempty"`
	BarCategoryGap string `json:"barCategoryGap,omitempty"`
	ShowBackground bool   `json:"showBackground,omitempty"`
	RoundCap       bool   `json:"roundCap,omitempty"`

	// Bar3D
	Shading string `json:"shading,omitempty"`

	// Graph
	Links              interface{} `json:"links,omitempty"`
	Layout             string      `json:"layout,omitempty"`
	Force              interface{} `json:"force,omitempty"`
	Categories         interface{} `json:"categories,omitempty"`
	Roam               bool        `json:"roam,omitempty"`
	EdgeSymbol         interface{} `json:"edgeSymbol,omitempty"`
	EdgeSymbolSize     interface{} `json:"edgeSymbolSize,omitempty"`
	EdgeLabel          interface{} `json:"edgeLabel,omitempty"`
	Draggable          bool        `json:"draggable,omitempty"`
	FocusNodeAdjacency bool        `json:"focusNodeAdjacency,omitempty"`

	//KLine
	BarWidth    string `json:"barWidth,omitempty"`
	BarMinWidth string `json:"barMinWidth,omitempty"`
	BarMaxWidth string `json:"barMaxWidth,omitempty"`

	// Line
	Step         interface{} `json:"step,omitempty"`
	Smooth       bool        `json:"smooth"`
	ConnectNulls bool        `json:"connectNulls"`
	ShowSymbol   bool        `json:"showSymbol"`
	Symbol       string      `json:"symbol,omitempty"`
	Color        string      `json:"color,omitempty"`

	//Liquid

	IsLiquidOutline string `json:"isLiquidOutline,omitempty"`
	IsWaveAnimation bool   `json:"isWaveAnimation,omitempty"`

	//Map

	MapType     string `json:"mapType,omitempty"`
	CoordSystem string `json:"coordSystem,omitempty"`

	//Pie
	RoseType interface{} `json:"roseType,omitempty"`
	Center   interface{} `json:"center,omitempty"`
	Radius   interface{} `json:"radius,omitempty"`

	// Scatter
	SymbolSize interface{} `json:"symbolSize,omitempty"`

	// Tree
	Orient            string      `json:"orient,omitempty"`
	ExpandAndCollapse bool        `json:"expandAndCollapse,omitempty"`
	InitialTreeDepth  int         `json:"initialTreeDepth,omitempty"`
	Leaves            interface{} `json:"leaves,omitempty"`
	Left              string      `json:"left,omitempty"`
	Right             string      `json:"right,omitempty"`
	Top               string      `json:"top,omitempty"`
	Bottom            string      `json:"bottom,omitempty"`

	//TreeMap
	LeafDepth  int         `json:"leafDepth,omitempty"`
	Levels     interface{} `json:"levels,omitempty"`
	UpperLabel interface{} `json:"upperLabel,omitempty"`

	//WordClound

	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// Sunburst
	NodeClick               string `json:"nodeClick,omitempty"`
	Sort                    string `json:"sort,omitempty"`
	RenderLabelForZeroData  bool   `json:"renderLabelForZeroData"`
	SelectedMode            bool   `json:"selectedMode"`
	Animation               bool   `json:"animation" default:"true"`
	AnimationThreshold      int    `json:"animationThreshold,omitempty"`
	AnimationDuration       int    `json:"animationDuration,omitempty"`
	AnimationEasing         string `json:"animationEasing,omitempty"`
	AnimationDelay          int    `json:"animationDelay,omitempty"`
	AnimationDurationUpdate int    `json:"animationDurationUpdate,omitempty"`
	AnimationEasingUpdate   string `json:"animationEasingUpdate,omitempty"`
	AnimationDelayUpdate    int    `json:"animationDelayUpdate,omitempty"`

	// series  data
	Data interface{} `json:"data"`

	//series options

	*opts.Encode        `json:"encode,omitempty"`
	*opts.ItemStyle     `json:"itemStyle,omitempty"`
	*opts.Label         `json:"label,omitempty"`
	*opts.LabelLine     `json:"labelLine,omitempty"`
	*opts.Emphasis      `json:"emphasis,omitempty"`
	*opts.MarkLines     `json:"markLines,omitempty"`
	*opts.MarkAreas     `json:"markAreas,omitempty"`
	*opts.MarkPoints    `json:"markPoints,omitempty"`
	*opts.RippleEffect  `json:"rippleEffect,omitempty"`
	*opts.LineStyle     `json:"lineStyle,omitempty"`
	*opts.AreaStyle     `json:"areaStyle,omitempty"`
	*opts.TextStyle     `json:"textStyle,omitempty"`
	*opts.CircularStyle `json:"circularStyle,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

func WithSeriesAnimation(enable bool) SeriesOpts {
	return func(s *SingleSeries) {
		s.Animation = enable
	}
}

// 设置表

func WithLabelOpts(opt opts.Label) SeriesOpts {
	return func(s *SingleSeries) {
		s.Label = &opt
	}
}

// WithEmphasisOpts sets the emphasis.
// WithEmphasisOpts设置重点。
func WithEmphasisOpts(opt opts.Emphasis) SeriesOpts {
	return func(s *SingleSeries) {
		s.Emphasis = &opt
	}
}

// //WithAreaStyleOps设置区域样式。
func WithAreaStyleOpts(opt opts.AreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.AreaStyle = &opt
	}
}

// WithItemStyleOpts sets the item style.
// ///WithRippleEffectOpts设置项目格式
func WithItemStyleOpts(opt opts.ItemStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.ItemStyle = &opt
	}
}

// WithRippleEffectOpts //WithRippleEffectOpts设置波纹效果。
func WithRippleEffectOpts(opt opts.RippleEffect) SeriesOpts {
	return func(s *SingleSeries) {
		s.RippleEffect = &opt
	}
}

// WithLineStyleOps设置线条样式。
func WithLineStyleOpts(opt opts.LineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.LineStyle = &opt
	}
}

func WithCircularStyleOpts(opt opts.CircularStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.CircularStyle = &opt
	}
}

/* chart options */

// ///WithBarChartOpts设置条形图选项。
func WithBarCharOpts(opt opts.BarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Stack = opt.Stack
		s.BarGap = opt.BarGap
		s.BarCategoryGap = opt.BarCategoryGap
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ShowBackground = opt.ShowBackground
		s.RoundCap = opt.RoundCap
		s.CoordSystem = opt.CoordSystem

	}
}

// /WithSunburstOps设置SunburstChart选项。
func WithSunburstOpts(opt opts.SunburstChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.NodeClick = opt.NodeClick
		s.Sort = opt.Sort
		s.RenderLabelForZeroData = opt.RenderLabelForZeroData
		s.SelectedMode = opt.SelectedMode
		s.Animation = opt.Animation
		s.AnimationThreshold = opt.AnimationThreshold
		s.AnimationDuration = opt.AnimationDuration
		s.AnimationEasing = opt.AnimationEasing
		s.AnimationDelay = opt.AnimationDelay
		s.AnimationDurationUpdate = opt.AnimationDurationUpdate
		s.AnimationEasingUpdate = opt.AnimationEasingUpdate
		s.AnimationDelayUpdate = opt.AnimationDelayUpdate
	}
}

// //WithGraphChartOpts设置GraphChart选项。
func WithGraphChartOpts(opt opts.GraphChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Force = opt.Force
		s.Roam = opt.Roam
		s.EdgeSymbol = opt.EdgeSymbol
		s.EdgeSymbolSize = opt.EdgeSymbolSize
		s.Draggable = opt.Draggable
		s.FocusNodeAdjacency = opt.FocusNodeAdjacency
		s.Categories = opt.Categories
		s.EdgeLabel = opt.EdgeLabel
	}
}

// /WithHeatMapChartOpts设置HeatMapChart选项。
func WithHeatMapChartOpts(opt opts.HeatMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLineChartOpts设置“折线图”选项。
func WithLineChartOpts(opt opts.LineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.YAxisIndex = opt.YAxisIndex
		s.Stack = opt.Stack
		s.Smooth = opt.Smooth
		s.ShowSymbol = opt.ShowSymbol
		s.Symbol = opt.Symbol
		s.SymbolSize = opt.SymbolSize
		s.Step = opt.Step
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ConnectNulls = opt.ConnectNulls
		s.Color = opt.Color
	}
}

// WithKlineChartOpts设置“LineChart”选项。
func WithKlineChartOpts(opt opts.KlineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.BarWidth = opt.BarWidth
		s.BarMinWidth = opt.BarMinWidth
		s.BarMaxWidth = opt.BarMaxWidth
	}
}

// WithPieChartOpts 设置“PieChart”选项。
// WithPieChartOpts sets the PieChart option.
func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

// WithScatterChartOpts sets the ScatterChart option.
func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLiquidChartOpts sets the LiquidChart option.
func WithLiquidChartOpts(opt opts.LiquidChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.IsLiquidOutline = opt.IsShowOutline
		s.IsWaveAnimation = opt.IsWaveAnimation
	}
}

// WithBar3DChartOpts sets the Bar3DChart option.
func WithBar3DChartOpts(opt opts.Bar3DChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shading = opt.Shading
	}
}

// WithTreeOpts sets the TreeChart option.

func WithTreeOpts(opt opts.TreeChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Orient = opt.Orient
		s.ExpandAndCollapse = opt.ExpandAndCollapse
		s.InitialTreeDepth = opt.InitialTreeDepth
		s.Roam = opt.Roam
		s.Label = opt.Label
		s.Leaves = opt.Leaves
		s.Right = opt.Right
		s.Left = opt.Left
		s.Top = opt.Top
		s.Bottom = opt.Bottom
	}
}

// WithTreeMapOpts sets the TreeMapChart options.
func WithTreeMapOpts(opt opts.TreeMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Animation = opt.Animation
		s.LeafDepth = opt.LeafDepth
		s.Roam = opt.Roam
		s.Levels = opt.Levels
		s.UpperLabel = opt.UpperLabel
		s.Right = opt.Right
		s.Left = opt.Left
		s.Top = opt.Top
		s.Bottom = opt.Bottom
	}
}

// WithWorldCloudChartOpts sets the WorldCloudChart option.
func WithWorldCloudChartOpts(opt opts.WordCloudChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.SizeRange = opt.SizeRange
		s.RotationRange = opt.RotationRange
	}
}

// WithMarkLineNameTypeItemOpts sets the type of the MarkLine.

func WithMarkLineNameTypeItemOpts(opt ...opts.MarkLineNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameCoordItemOpts sets the coordinates of the MarkLine.

func WithMarkLineNameCoordItemOpts(opt ...opts.MarkLineNameCoordItem) SeriesOpts {
	type MLNameCoord struct {
		Name  string        `json:"name,omitempty"`
		Coord []interface{} `json:"coord"`
	}
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}

		for _, o := range opt {
			s.MarkLines.Data = append(
				s.MarkLines.Data,
				[]MLNameCoord{{Name: o.Name, Coord: o.Coordinate0}, {Coord: o.Coordinate1}},
			)
		}
	}
}

// WithMarkLineNameXAxisItemOpts sets the X axis of the MarkLine.

func WithMarkLineNameXAxisItemOpts(opt ...opts.MarkLineNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameYAxisItemOpts sets the Y axis of the MarkLine.

func WithMarkLineNameYAxisItemOpts(opt ...opts.MarkLineNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkAreaNameTypeItemOpts sets the type of the MarkArea.

func WithMarkAreaNameTypeOItemOpts(opt ...opts.MarkAreaNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkAreaStyleOpts sets the style of the MarkArea.
func WithMarkAreaStyleOpts(opt opts.MarkAreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}

		s.MarkAreas.MarkAreaStyle = opt
	}
}

// WithMarkAreaNameCoordItemOpts sets the coordinates of the MarkLine.

//WithMarkAreaNameCoordItemOpts设置标记线的坐标。

func WithMarkAreaNameCoordItemOpts(opt ...opts.MarkAreaNameCoordItem) SeriesOpts {
	type MANameCoord struct {
		Name      string          `json:"name,omitempty"`
		ItemStyle *opts.ItemStyle `json:"itemStyle"`
		Coord     []interface{}   `json:"coord"`
	}

	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}

		for _, o := range opt {
			s.MarkAreas.Data = append(
				s.MarkAreas.Data,
				[]MANameCoord{
					{Name: o.Name, ItemStyle: o.ItemStyle, Coord: o.Coordinate0},
					{Coord: o.Coordinate1},
				},
			)
		}
	}
}

// // WithMarkAreaNameXAxisItemOpts sets the X axis of the MarkLine.
// WithMarkAreaNameXAxisItemOpts设置标记线的X轴。
func WithMarkAreaNameXAxisItemOpts(opt ...opts.MarkAreaNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkAreaNameYAxisItemOpts sets the Y axis of the MarkLine.
func WithMarkAreaNameYAxisItemOpts(opt ...opts.MarkAreaNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkPointNameTypeItemOpts设置MarkPoint的类型。
func WithMarkPointNameTypeItemOpt(opt ...opts.MarkPointNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

// //WithMarkPointStyleOps设置MarkPoint的样式。
func WithMarkPointStyleOpts(opt opts.MarkPointStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}

		s.MarkPoints.MarkPointStyle = opt
	}
}

/////WithMarkPointNameCoordItemOpts设置MarkPoint的坐标。

func WithMarkPointNameCoordItemOpts(opt ...opts.MarkPointNameCoordItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

func (s *SingleSeries) InitSeriesDefaultOpts(c BaseConfiguration) {
	opts.SetDefaultValue(s)
	////从BaseConfiguration继承的一些特殊选项
	s.Animation = c.Animation
}

func (s *SingleSeries) ConfigureSeriesOpts(options ...SeriesOpts) {
	for _, opt := range options {
		opt(s)
	}
}

// ///MultiSeries表示多个系列。
type MultiSeries []SingleSeries

// 设置系列选项设置所有系列的选项。
// 以前的选项每次都会被覆盖，因此如果需要，请将它们设置在“AddSeries”上
// 单独自定义每个系列
// 此处->↓ <-
// func（c*Bar）AddSeries（名称字符串，data[]opts.BarData，options…SeriesOpts）
func (ms *MultiSeries) SetSeriesOptions(opts ...SeriesOpts) {
	s := *ms
	for i := 0; i < len(s); i++ {
		s[i].ConfigureSeriesOpts(opts...)
	}
}

// //WithEncodeOpts Set为数据集编码
func WithEncodeOpts(opt opts.Encode) SeriesOpts {
	return func(s *SingleSeries) {
		s.Encode = &opt
	}
}
