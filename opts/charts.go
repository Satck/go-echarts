package opts

// BarChart
// https://echarts.apache.org/en/option.html#series-bar
type BarChart struct {
	Type string
	//堆栈的名称。在同一类别轴上，具有
	//相同的堆栈名称将被放在彼此的顶部。
	Stack string
	//不同系列之间的条之间的间隙是一个百分比值，如“30%”，
	//这意味着棒宽度的30%。
	//将barGap设置为“-100%”可以重叠属于不同系列的条，
	//这在放置一系列条形图作为背景时是有用的。
	//在单个坐标系中，该属性由多个“条形”系列共享。
	//该属性应设置在坐标系中的最后一个“条形”系列上，
	//那么它将被坐标系中的所有“bar”系列所采用。
	BarGap string
	//单个系列的条间距默认为类别间距的20%，
	//可以设置为固定值。
	//在单个坐标系中，该属性由多个“条形”系列共享。
	//该属性应设置在坐标系中的最后一个“条形”系列上，
	//那么它将被坐标系中的所有“bar”系列所采用。
	BarCategoryGap string
	//要与之组合的x轴索引，这对于一个图表中的多个x轴很有用。
	XAxisIndex int
	//要组合的y轴索引，这对于一个图表中的多个y轴很有用。
	YAxisIndex int

	ShowBackground bool
	RoundCap       bool
	CoordSystem    string
}

// SunburstChart
// https://echarts.apache.org/en/option.html#series-sunburst
type SunburstChart struct {
	////单击扇区的操作
	NodeClick string `json:"nodeClick,omitempty"`
	//扇区基于值使用的排序方法
	Sort string `json:"sort,omitempty"`
	//如果没有名称，是否需要渲染。
	RenderLabelForZeroData bool `json:"renderLabelForZeroData"`
	//所选模式
	SelectedMode bool `json:"selectedMode"`
	//是否启用动画。
	Animation bool `json:"animation"`
	//是否将图形数量阈值设置为动画
	AnimationThreshold int `json:"animationThreshold,omitempty"`
	//用于第一个动画的缓和方法
	AnimationDuration int `json:"animationDuration,omitempty"`
	//第一个动画的持续时间
	AnimationEasing string `json:"animationEasing,omitempty"`
	//更新第一个动画之前的延迟
	AnimationDelay int `json:"animationDelay,omitempty"`
	//动画完成时间
	AnimationDurationUpdate int `json:"animationDurationUpdate,omitempty"`
	//用于动画的缓和方法。
	AnimationEasingUpdate string `json:"animationEasingUpdate,omitempty"`
	//更新动画之前的延迟
	AnimationDelayUpdate int `json:"animationDelayUpdate,omitempty"`
}

// BarData
// https://echarts.apache.org/en/option.html#series-bar.data

type BarData struct {
	////数据项的名称。
	Name string `json:"name,omitempty"`
	//单个数据项的值。
	Value interface{} `json:"value,omitempty"`
	////单个条形图中文本标签的样式设置。
	Label *Label `json:"label,omitempty"`
	//此系列数据中的ItemStyle设置。
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
	////此系列数据中的工具提示设置。
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

// Bar3DChart is the option set for a 3D bar chart.
type Bar3DChart struct {

	// Shading is the coloring effect of 3D graphics in 3D Bar.
	// The following three coloring methods are supported in echarts-gl:
	// Options:
	//
	// * "color": Only display colors, not affected by other factors such as lighting.
	// * "lambert": Through the classic [lambert] coloring, can express the light and dark that the light shows.
	// * "realistic": Realistic rendering, combined with light.ambientCubemap and postEffect,
	//   can improve the quality and texture of the display.
	//   [Physical Based Rendering (PBR)] (https://www.marmoset.co/posts/physically-based-rendering-and-you-can-too/)
	//   is used in ECharts GL to represent realistic materials.
	//着色是三维条形图中三维图形的着色效果。
	//echarts gl支持以下三种着色方法：
	//选项：
	//
	//*“颜色”：仅显示颜色，不受照明等其他因素影响。
	//*“蓝伯特”：通过经典的[蓝伯特]着色，可以表达光线所表现出的明暗。
	//*“逼真”：逼真渲染，结合light.ambientCubemap和postEffect，
	//可以提高显示器的质量和纹理。
	//[基于物理的绘制（PBR）](https://www.marmoset.co/posts/physically-based-rendering-and-you-can-too/)
	//在ECharts GL中用于表示逼真的材质。
	Shading string
}

// BoxPlotData
// https://echarts.apache.org/en/option.html#series-boxplot.data

type BoxPlotData struct {
	////数据项的名称
	Name string `json:"name,omitempty"`
	//单个数据项的值。
	Value interface{} `json:"value,omitempty"`
	//单个条形图中文本标签的样式设置。
	Label *Label `json:"label,omitempty"`
	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"ItemStyle,omitempty"`
	// Emphasis settings in this series data.
	Emphasis *Emphasis `json:"emphasis,omitempty"`
	// Tooltip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

// EffectScatterData
// https://echarts.apache.org/en/option.html#series-effectScatter.data

type EffectScatterData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// FunnelData
// https://echarts.apache.org/en/option.html#series-funnel.data

type FunnelData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GeoData
type GeoData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GaugeData
// https://echarts.apache.org/en/option.html#series-gauge.data
type GaugeData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GraphChart is the option set for graph chart.
// https://echarts.apache.org/en/option.html#series-graph
type GraphChart struct {
	// Graph layout.
	// * 'none' No layout, use x, y provided in node as the position of node.
	// * 'circular' Adopt circular layout, see the example Les Miserables.
	// * 'force' Adopt force-directed layout, see the example Force, the
	// detail about layout configurations are in graph.force1
	Layout string
	// Force is the option set for graph force layout.
	////Force 是为图形Force布局设置的选项。
	Force *GraphForce

	//是否启用鼠标缩放和平移。默认情况下为false。

	//如果需要缩放或平移，可以将其设置为“缩放”或“移动”。

	//否则，将其设置为true以同时启用这两个选项。

	Roam bool
	// EdgeSymbol is the symbols of two ends of edge line.
	// * 'circle'
	// * 'arrow'
	// * 'none'
	// example: ["circle", "arrow"] or "circle"
	EdgeSymbol interface{}
	// EdgeSymbolSize is size of symbol of two ends of edge line. Can be an array or a single number
	//边缘线两端的符号大小。可以是数组或单个数字
	// example: [5,10] or 5
	EdgeSymbolSize interface{}
	// Draggable allows you to move the nodes with the mouse if they are not fixed.
	Draggable bool
	// Whether to focus/highlight the hover node and it's adjacencies.
	////是否聚焦/高亮显示悬停节点及其相邻区域。
	FocusNodeAdjacency bool
	//节点的类别，这是可选的。如果存在节点的分类，
	//每个节点的类别可以通过data[i].category来分配。
	//类别的样式也将应用于节点的样式。类别也可以在图例中使用。
	Categories []*GraphCategory
	////EdgeLabel是边的标签的属性。
	EdgeLabel *EdgeLabel `json:"edgeLabel"`
}

// GraphNode represents a data node in graph chart.
// https://echarts.apache.org/en/option.html#series-graph.data
type GraphNode struct {
	Name string `json:"name,omitempty"`
	// x value of node position.
	//节点位置的x值。
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
	//Value of data item.
	Value float32 `json:"value,omitempty"`
	//If node are fixed when doing force directed layout.
	Fixed bool `json:"fixed,omitempty"`
	// Index of category which the data item belongs to.
	////数据项所属类别的索引。
	Category interface{} `json:"category,omitempty"`
	//此类别节点的符号。
	//ECharts提供的图标类型包括
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	//它可以设置为带有'的图像image://url'，其中URL是指向图像的链接，或图像的dataURI。
	Symbol string `json:"symbol,omitempty"`
	// node of this category symbol size. It can be set to single numbers like 10,
	// or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	SymbolSize interface{} `json:"symbolSize,omitempty"`

	// The style of this node.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// GraphLink represents relationship between two data nodes.
// https://echarts.apache.org/en/option.html#series-graph.links

type GraphLink struct {
	// A string representing the name of source node on edge. Can also be a number representing the node index.
	//表示边缘上源节点名称的字符串。也可以是表示节点索引的数字。
	Source interface{} `json:"source,omitempty"`
	// A string representing the name of target node on edge. Can also be a number representing node index.

	Target interface{} `json:"target,omitempty"`
	// value of edge, can be mapped to edge length in force graph.

	Value float32 `json:"value,omitempty"`
	// Label for this link.

	//此链接的标签。
	Label *EdgeLabel `json:"label,omitempty"`
}

// GraphCategory represents a category for data nodes.
// The categories of node, which is optional. If there is a classification of nodes,
// the category of each node can be assigned through data[i].category.
// And the style of category will also be applied to the style of nodes. categories can also be used in legend.
// GraphCategory表示数据节点的类别。
// 节点的类别，这是可选的。如果存在节点的分类，
// 每个节点的类别可以通过data[i].category来分配。
// 类别的样式也将应用于节点的样式。类别也可以在图例中使用。
type GraphCategory struct {
	//类别的名称，用于与图例和工具提示的内容相对应。
	Name string `json:"name"`

	// The label style of node in this category.
	///此类别中节点的标签样式。
	Label *Label `json:"label,omitempty"`
}

// HeatMapChart is the option set for a heatmap chart.
// https://echarts.apache.org/en/option.html#series-heatmap
type HeatMapChart struct {

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	//要与之组合的x轴索引，这对于一个图表中的多个x轴很有用。
	XAxisIndex int
	//y轴
	YAxisIndex int
}

// HeatMapData
// https://echarts.apache.org/en/option.html#series-heatmap.data

type HeatMapData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// KlineData
// https://echarts.apache.org/en/option.html#series-candlestick.data

type KlineData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// LineChart is the options set for a line chart.
// https://echarts.apache.org/en/option.html#series-line
type LineChart struct {
	// If stack the value. On the same category axis, the series with the same stack name would be put on top of each other.
	// The effect of the below example could be seen through stack switching of toolbox on the top right corner:
	//如果堆栈值。在相同的类别轴上，具有相同堆栈名称的系列将被放在彼此的顶部。
	//通过右上角工具箱的堆栈切换可以看出以下示例的效果：
	Stack string
	// Whether to show as smooth curve.
	// If is typed in boolean, then it means whether to enable smoothing. If is
	// typed in number, valued from 0 to 1, then it means smoothness. A smaller value makes it less smooth.
	//是否显示为平滑曲线。
	//如果以布尔值键入，则表示是否启用平滑。如果是
	//键入数字，值从0到1，则表示平滑。值越小，其平滑度越低。
	Smooth bool

	// Whether to show as a step line. It can be true, false. Or 'start', 'middle', 'end'.
	// Which will configure the turn point of step line.
	//是否显示为阶梯线。它可以是真的，也可以是假的。或者“开始”、“中间”、“结束”。
	//这将配置步进线的转折点。

	Step interface{}

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	//要与之组合的x轴索引，这对于一个图表中的多个x轴很有用。
	XAxisIndex int
	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	//要组合的y轴索引，这对于一个图表中的多个y轴很有用。
	YAxisIndex int
	//Whether to connect the line across null points.
	//是否连接跨空点的线。
	ConnectNulls bool
	// Whether to show symbol. It would be shown during tooltip hover.
	//是否显示符号。它将在工具提示悬停期间显示。
	ShowSymbol bool
	// Icon types provided by ECharts includes
	//  'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbol

	Symbol string
	// symbol size. It can be set to single numbers like 10, or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbolSize
	SymbolSize interface{}
	// color for Line series. it affects Line series including symbols, unlike LineStyle.Color
	Color string
}

// LineChart is the options set for a chandlestick chart.
// https://echarts.apache.org/en/option.html#series-candlestick

type KlineChart struct {
	// Specify bar width. Absolute value (like 10) or percentage (like '20%', according to band width) can be used. Auto adapt by default.
	//指定条形图宽度。可以使用绝对值（如10）或百分比（如“20%”，根据带宽）。默认情况下自动调整。
	BarWidth string
	// min  width
	BarMinWidth string
	// max width
	BarMaxWidth string
}

// LineData
// https://echarts.apache.org/en/option.html#series-line.data

type LineData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
	// Symbol of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`
	// single data symbol size. It can be set to single numbers like 10, or
	// use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10
	SymbolSize int `json:"symbolSize,omitempty"`
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int
	//y
	YAxisIndex int
}

// LiquidChart
// reference https://github.com/ecomfe/echarts-liquidfill

type LiquidChart struct {
	// Shape of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Shape string

	// Whether to show outline
	IsShowOutline string

	// Whether to stop animation
	IsWaveAnimation bool
}

// LiquidData
// reference https://github.com/ecomfe/echarts-liquidfill
type LiquidData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// MapData
// https://echarts.apache.org/en/option.html#series-map.data

type MapData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// ParallelData
// https://echarts.apache.org/en/option.html#series-parallel.data
type ParallelData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// PieChart is the option set for a pie chart.
// https://echarts.apache.org/en/option.html#series-pie

type PieChart struct {
	// Whether to show as Nightingale chart, which distinguishes data through radius. There are 2 optional modes:
	// * 'radius' Use central angle to show the percentage of data, radius to show data size.
	// * 'area' All the sectors will share the same central angle, the data size is shown only through radiuses.

	////是否显示为Nightingale chart，该图通过半径区分数据。有两种可选模式：
	RoseType string
	// Center position of Pie chart, the first of which is the horizontal position, and the second is the vertical position.
	// 第一个是水平位置，第二个是垂直位置。
	// Percentage is supported. When set in percentage, the item is relative to the container width,
	// and the second item to the height.
	// Example:
	// Set to absolute pixel values ->> center: [400, 300]
	// Set to relative percent ->> center: ['50%', '50%']
	Center interface{}

	// Radius of Pie chart. Value can be:
	// * number: Specify outside radius directly.
	// * string: For example, '20%', means that the outside radius is 20% of the viewport
	// size (the little one between width and height of the chart container).
	//
	// Array.<number|string>: The first item specifies the inside radius, and the
	// second item specifies the outside radius. Each item follows the definitions above.
	Radius interface{}
}

// PieData
// https://echarts.apache.org/en/option.html#series-pie.data
type PieData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
	// Whether the data item is selected.
	Selected bool `json:"selected,omitempty"`
	// The label configuration of a single sector.
	////单个扇区的标签配置。
	Label *Label `json:"label,omitempty"`

	// Graphic style of , emphasis is the style when it is highlighted, like being hovered by mouse, or highlighted via legend connect.
	//Graphic 的图形样式，强调是高亮显示时的样式，如鼠标悬停或通过图例连接高亮显示。

	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
	//tooltip settings in this series data.
	//该系列数据中的工具提示设置。
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

// RadarData
// https://echarts.apache.org/en/option.html#series-radar

type RadarData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// SankeyLink represents relationship between two data nodes.
// https://echarts.apache.org/en/option.html#series-sankey.links

type SankeyLink struct {
	///边的源节点名称
	Source interface{} `json:"source,omitempty"`
	////边缘的目标节点的名称
	Target interface{} `json:"target,omitempty"`
	//边的值，它决定边的宽度。
	Value float32 `json:"value,omitempty"`
}

// SankeyNode represents a data node.
// https://echarts.apache.org/en/option.html#series-sankey.nodes

type SankeyNode struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value string `json:"value,omitempty"`
	// Depth of the node within the chart
	//图表中节点的深度
	Depth *int `json:"depth,omitempty"`
	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// ScatterChart is the option set for a scatter chart.
// https://echarts.apache.org/en/option.html#series-scatter

type ScatterChart struct {
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	//要与之组合的x轴索引，这对于一个图表中的多个x轴很有用。
	XAxisIndex int
	// y轴

	YAxisIndex int
}

// ScatterData
// https://echarts.apache.org/en/option.html#series-scatter.data

type ScatterData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	SymbolSize int `json:"symbolSize,omitempty"`

	SymbolRotate int `json:"symbolRotate,omitempty"`
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	//要与之组合的x轴索引，这对于一个图表中的多个x轴很有用。
	XAxisIndex int `json:"XAxisIndex,omitempty"`

	YAxisIndex int `json:"YAxisIndex,omitempty"`
}

// ThemeRiverData
// https://echarts.apache.org/en/option.html#series-themeRiver

type ThemeRiverData struct {
	// the time attribute of time and theme.
	//时间和主题的时间属性。
	Date string `json:"date,omitempty"`
	// the value of an event or theme at a time point.
	//事件或主题在某个时间点的价值。
	Value float64 `json:"value,omitempty"`
	// the name of an event or theme.
	Name string `json:"name,omitempty"`
}

// ToList  converts the themeriver data to a list
// ToList将中间数据转换为列表
func (trd ThemeRiverData) ToList() [3]interface{} {
	return [3]interface{}{trd.Date, trd.Value, trd.Name}
}

//WordCloudChart is the option set for a word cloud  chart
//WordCloudChart是为单词云图表设置的选项。

type WordCloudChart struct {
	// Shape of WordCloud
	// Optional: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	Shape string
	// range of font size
	SizeRange []float32
	// range of font rotation angle
	//字体旋转角度范围
	RotationRange []float32
}

// WordCloudData
type WordCloudData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

type Chart3DData struct {
	// Name of the data item.
	Name string `json:"name,omitempty"`
	// Value of the data item.
	// []interface{}{1, 2, 3}
	Value []interface{} `json:"value,omitempty"`
	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
	// The style setting of the text label in a single bar.
	//单个条形图中文本标签的样式设置。
	Label *Label `json:"label,omitempty"`
}

type TreeChart struct {
	// The layout of the tree, which can be orthogonal and radial.
	// * 'orthogonal' refer to the horizontal and vertical direction.
	// * 'radial' refers to the view that the root node as the center and each layer of nodes as the ring.
	//树的布局，可以是正交的，也可以是放射状的。
	//*“正交”是指水平方向和垂直方向。
	//*“径向”是指以根节点为中心，
	Layout string
	// The direction of the orthogonal layout in the tree diagram.
	// * 'from left to right' or 'LR'
	// * 'from right to left' or 'RL'
	// * 'from top to bottom' or 'TB'
	// * 'from bottom to top' or 'BT'
	//树状图中正交布局的方向。
	//*“从左到右”或“LR”
	//*“从右向左”或“RL”
	//*“从上到下”或“TB”
	//*“自下而上”或“BT”
	Orient string `json:"orient,omitempty"`

	// Whether to enable mouse zooming and translating. false by default.
	// If either zooming or translating is wanted, it can be set to 'scale' or 'move'.
	// Otherwise, set it to be true to enable both.
	//是否启用鼠标缩放和平移。默认情况下为false。

	//如果需要缩放或平移，可以将其设置为“缩放”或“移动”。

	//否则，将其设置为true以同时启用这两个选项。
	Roam bool `json:"roam"`

	// Subtree collapses and expands interaction, default true.
	//子树折叠并展开交互，默认为true。
	ExpandAndCollapse bool `json:"expandAndCollapse,omitempty"`

	// The initial level (depth) of the tree. The root node is the 0th layer, then the first layer, the second layer, ... , until the leaf node.
	// This configuration item is primarily used in conjunction with collapsing and expansion interactions.
	// The purpose is to prevent the nodes from obscuring each other. If set as -1 or null or undefined, all nodes are expanded.
	//树的初始级别（深度）。根节点是第0层，然后是第一层，第二层，直到叶节点。

	//此配置项主要用于折叠和展开交互。

	//其目的是防止节点相互遮挡。如果设置为-1、null或未定义，则所有节点都将展开。

	InitialTreeDepth int `json:"initialTreeDepth,omitempty"`
	// The style setting of the text label in a single bar.
	//单个条形图中文本标签的样式设置。
	Label *Label `json:"label,omitempty"`
	// Leaf node special configuration, the leaf node and non-leaf node label location is different.
	//叶节点特殊配置，叶节点与非叶节点的标签位置不同。
	Leaves *TreeLeaves `json:"leaves,omitempty"`

	// Distance between tree component and the sides of the container.
	// value can be instant pixel value like 20;
	// It can also be a percentage value relative to container width like '20%';
	//树组件与容器侧面之间的距离。
	//值可以是像20这样的即时像素值；
	//它也可以是相对于容器宽度的百分比值，如“20%”；

	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
	Top    string `json:"top,omitempty"`
	Bottom string `json:"bottom,omitempty"`
}

type TreeData struct {
	// Name of the data item.
	Name string `json:"name,omitempty"`
	// Value of the data item.
	Value int `json:"value,omitempty"`

	Children []*TreeData `json:"children,omitempty"`
	// Symbol of node of this category.
	// Icon types provided by ECharts includes
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`
	//此类别符号大小的节点。
	// node of this category symbol size. It can be set to single numbers like 10,
	// or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	SymbolSize interface{} `json:"symbolSize,omitempty"`
	// If set as `true`, the node is collapsed in the initialization.
	//节点在初始化过程中被折叠。
	Collapsed bool `json:"collapsed,omitempty"`
	// LineStyle settings in this series data.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}
type TreeMapChart struct {
	Animation bool `json:"animation"`

	LeafDepth int `json:"leafDepth,omitempty"`
	//Roam描述了是否启用鼠标缩放和平移。默认情况下为false。
	Roam bool `json:"roam"`
	//Label描述每个节点中标签的样式。
	Label *Label `json:"label,omitempty"`
	//UpperLabel用于指定当树映射节点具有子节点时是否显示标签。
	UpperLabel *UpperLabel `json:"upperLabel,omitempty"`

	// ColorMappingBy specifies the rule according to which each node obtain color from color list.
	//ColorMappingBy指定每个节点从颜色列表中获取颜色的规则。
	ColorMappingBy string `json:"colorMappingBy,omitempty"`
	//Levels为每个节点级别提供配置
	Levels *[]TreeMapLevel `json:"levels,omitempty"`

	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
	Top    string `json:"top,omitempty"`
	Bottom string `json:"bottom,omitempty"`
}
type TreeMapNode struct {
	// Name of the tree node item.
	Name string `json:"name"`
	// Value of the tree node item.
	Value int `json:"value,omitempty"`

	Children []TreeMapNode `json:"children,omitempty"`
}

type SunBurstData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`
	// Value of data item.
	Value float64 `json:"value,omitempty"`
	// sub item of data item
	Children []*SunBurstData `json:"children,omitempty"`
}
