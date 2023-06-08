package opts

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/types"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Initialization struct {
	//Html title
	PageTitle string `default:"Awesome go-echarts"`
	// Width of canvas
	Width string `default:"900px"`
	// Height of canvas
	Height string `default:"500px"`
	// BackgroundColor  of canvas
	BackgroundColor string
	//Chart unique  Id
	ChartID string
	// Assets  host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets"`
	// Theme of chart
	Theme string `default:"White"`
}

// Validate验证初始化配置。
// Validate validates the initialization configurations.
func (opt *Initialization) Validate() {
	SetDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = generateUniqueID()
	}
}

// SetDefaultValue set default values for the struct field.
// 为这个结构体字段设置默认值
// inspired from: https://github.com/mcuadros/go-defaults
func SetDefaultValue(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	walkField(elem)
}

func walkField(val reflect.Value) {
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		f := val.Field(i)
		if f.Kind() == reflect.Struct {
			walkField(f)
		}

		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(val.Field(i), defaultVal)
		}
	}
}

// setField handles String/Bool types only.

func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}

const (
	chartIDSize = 12
)

// generate the unique ID for each chart.
func generateUniqueID() string {
	var b [chartIDSize]byte
	for i := range b {
		b[i] = randByte()
	}
	return string(b[:])
}

func randByte() byte {
	c := 65 //A
	if rand.Intn(10) > 5 {
		c = 97
	}
	// The hyper link of main title text.
	//主标题文本的超链接。
	return byte(c + rand.Intn(26))
}

// Title is the option set for a title component.
///Title是为标题组件设置的选项。
// https://echarts.apache.org/en/option.html#title

type Title struct {
	// The main title text, supporting \n for newlines.
	////主标题文本，支持\n换行。
	Title string `json:"title,omitempty"`
	// TextStyle of the main title.
	TitleStyle *TextStyle `json:"titleStyle,omitempty""`
	// The hyper link of main title text.
	//主标题文本的超链接。
	Link string `json:"link,omitempty"`
	// Subtitle text, supporting \n for newlines.
	Subtitle string `json:"subtitle,omitempty"`
	// TextStyle of the sub title.
	SubLink string `json:"subLink,omitempty"`
	// Open the hyper link(超链接） of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target string `json:"target,omitempty"`

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	//标题组件与容器顶部之间的距离。
	//最高值可以是像20这样的即时像素值；也可以是百分比
	//相对于容器宽度的值，如“20%”；也可以是“顶部”、“中间”或“底部”。
	//如果左侧值设置为“top”、“middle”或“bottom”，
	//则部件将基于位置自动对准。
	Top string `json:"top,omitempty"`
	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`
	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`
	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`
}

/// Legend is the option set for a legend component.//Legend 是对于legend组成部分的一个选项
// Legend component shows symbol, color and name of different series. You can click legends to toggle displaying series in the chart.
// https://echarts.apache.org/en/option.html#legend

type Legend struct {

	// Whether to show the Legend, default true.
	// Once you set other options, need to manually set it to true // 手动设置为true

	Show bool `json:"show" default:"true"`
	// Type of legend. Optional values:
	// "plain": Simple legend. (default)
	// "scroll": Scrollable legend. It helps when too many legend items needed to be shown.
	Type string `json:"type"`

	// Distance between legend component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	//图例组件与容器左侧之间的距离。

	//左边的值可以是像20这样的即时像素值；也可以是百分比

	//相对于容器宽度的值，如“20%”；也可以是“左”、“中”或“右”。

	//如果左值设置为“左”、“中”或“右”，则组件

	//将根据位置自动对齐。
	Left string `json:"left,omitempty"`
	// Distance between legend component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`
	// Distance between legend component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`
	// Distance between legend component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Data array of legend. An array item is usually a name representing string.
	// set Data as []string{} if you wants to hide the legend.
	Data interface{} `json:"data,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`
	// Legend color when not selected.
	InactiveColor string `json:"inactiveColor,omitempty"`
	//所选图例的状态表。
	// State table of selected legend.
	// example:
	// var selected = map[string]bool{}
	// selected["series1"] = true
	// selected["series2"] = false
	Selected map[string]bool `json:"selected,omitempty"`

	// Selected mode of legend, which controls whether series can be toggled (切换）displaying by clicking legends.
	// It is enabled by default, and you may set it to be false to disabled it.
	// Besides, it can be set to 'single' or 'multiple', for single selection and multiple selection.
	//图例的选定模式，控制是否可以通过单击图例切换显示系列。
	//它在默认情况下是启用的，您可以将其设置为false以禁用它。

	//此外，可以设置为“单个”或“多个”，用于单选和多选。
	SelectedMode string `json:"selectedMode,omitempty"`
	// Legend space around content. The unit is px.
	// Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	// Examples:
	// 1. Set padding to be 5
	//    padding: 5
	// 2. Set the top and bottom paddings to be 5, and left and right paddings to be 10
	//    padding: [5, 10]
	// 3. Set each of the four paddings separately
	//    padding: [
	//      5,  // up
	//      10, // right
	//      5,  // down
	//      10, // left
	//    ]

	Padding interface{} `json:"padding,omitempty"`
	// Image width of legend symbol.
	ItemWidth int `json:"itemWidth,omitempty"`
	// Image height of legend symbol.
	ItemHeight int `json:"itemHeight,omitempty"`
	// Legend X position, right/left/center
	X string `json:"x,omitempty"`
	// Legend Y position, right/left/center
	Y string `json:"y,omitempty"`
	// Width of legend component. Adaptive by default.
	Width string `json:"width,omitempty"`
	// Height of legend component. Adaptive by default.
	Height string `json:"height,omitempty"`
	// Legend marker and text aligning.
	//图例标记和文本对齐。
	// By default, it automatically calculates from component location and orientation(定向）
	// When left value of this component is 'right' and orient is 'vertical', it would be aligned to 'right'.
	// Options: auto/left/right
	Align string `json:"align,omitempty"`
	// Legend text style.
	TextStyle *TextStyle `json:"textStyle,omitempty"`
}

// Tooltip is the option set for a tooltip component.
// https://echarts.apache.org/en/option.html#tooltip

type Tooltip struct {
	// Whether to show the tooltip component, including tooltip floating layer and axisPointer.
	//是否显示工具提示组件，包括工具提示浮动层和轴指针。
	Show bool `json:"show"`
	// Type of triggering.
	// Options:
	// * 'item': Triggered by data item, which is mainly used for charts that
	//    don't have a category axis like scatter charts or pie charts.
	// * 'axis': Triggered by axes, which is mainly used for charts that have category axes,
	//    like bar charts or line charts.
	// * 'none': Trigger nothing.

	Trigger string `json:"trigger,omitempty"`
	// Conditions to trigger tooltip. Options:
	// * 'mousemove': Trigger when mouse moves.
	// * 'click': Trigger when mouse clicks.
	// * 'mousemove|click': Trigger when mouse clicks and moves.
	// * 'none': Do not triggered by 'mousemove' and 'click'. Tooltip can be triggered and hidden
	//    manually by calling action.tooltip.showTip and action.tooltip.hideTip.
	//    It can also be triggered by axisPointer.handle in this case.

	TriggerOh string `json:"triggerOh,omitempty"`
	// Whether mouse is allowed to enter the floating layer of tooltip, whose default value is false.
	// If you need to interact in the tooltip like with links or buttons, it can be set as true.
	//是否允许鼠标进入工具提示的浮动层，默认值为false。
	//如果需要在工具提示中进行交互（如与链接或按钮交互），可以将其设置为true。
	Enterable bool `json:"enterable,omitempty"`

	// The content formatter of tooltip's floating layer which supports string template and callback function.
	//
	// 1. String template
	// The template variables are {a}, {b}, {c}, {d} and {e}, which stands for series name,
	// data name and data value and ect. When trigger is set to be 'axis', there may be data from multiple series.
	// In this time, series index can be refereed as {a0}, {a1}, or {a2}.
	// {a}, {b}, {c}, {d} have different meanings for different series types:
	//
	// * Line (area) charts, bar (column) charts, K charts: {a} for series name,
	//   {b} for category name, {c} for data value, {d} for none;
	// * Scatter (bubble) charts: {a} for series name, {b} for data name, {c} for data value, {d} for none;
	// * Map: {a} for series name, {b} for area name, {c} for merging data, {d} for none;
	// * Pie charts, gauge charts, funnel charts: {a} for series name, {b} for data item name,
	//   {c} for data value, {d} for percentage.
	//
	// 2. Callback function
	// The format of callback function:
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// The first parameter params is the data that the formatter needs. Its format is shown as follows:
	// {
	//    componentType: 'series',
	//    // Series type
	//    seriesType: string,
	//    // Series index in option.series
	//    seriesIndex: number,
	//    // Series name
	//    seriesName: string,
	//    // Data name, or category name
	//    name: string,
	//    // Data index in input data array
	//    dataIndex: number,
	//    // Original data as input
	//    data: Object,
	//    // Value of data. In most series it is the same as data.
	//    // But in some series it is some part of the data (e.g., in map, radar)
	//    value: number|Array|Object,
	//    // encoding info of coordinate system
	//    // Key: coord, like ('x' 'y' 'radius' 'angle')
	//    // value: Must be an array, not null/undefined. Contain dimension indices, like:
	//    // {
	//    //     x: [2] // values on dimension index 2 are mapped to x axis.
	//    //     y: [0] // values on dimension index 0 are mapped to y axis.
	//    // }
	//    encode: Object,
	//    // dimension names list
	//    dimensionNames: Array<String>,
	//    // data dimension index, for example 0 or 1 or 2 ...
	//    // Only work in `radar` series.
	//    dimensionIndex: number,
	//    // Color of data
	//    color: string,
	//
	//    // the percentage of pie chart
	//    percent: number,
	// }

	Formatter string `json:"formatter,omitempty"`

	// Configuration item for axisPointer
	//轴指针的配置项
	AxisPointer *AxisPointer `json:"axisPointer,omitempty"`
}

type AxisPointer struct {
	// Indicator type.
	// Options:
	//   - 'line' line indicator.
	//   - 'shadow' shadow crosshair indicator.
	//   - 'none' no indicator displayed.
	//   - 'cross' crosshair indicator, which is actually the shortcut of enable two axisPointers of two orthometric axes.
	Type string `json:"type,omitempty"`

	// 	Whether snap to point automatically. The default value is auto determined.
	// This feature usually makes sense in value axis and time axis, where tiny points can be seeked automatically.
	//是否自动捕捉到点。默认值是自动确定的。
	//这个功能通常在值轴和时间轴上是有意义的，在这些轴上可以自动查找微小的点。
	Snap bool `json:"snap,omitempty"`

	Link []AxisPointer `json:"link,omitempty"`

	Axis string `json:"axis,omitempty"`
}

type AxisPointerLink struct {
	XAxisIndex []int  `json:"xAxisIndex,omitempty"`
	YAxisIndex []int  `json:"yAxisIndex,omitempty"`
	XAxisName  string `json:"xAxisName,omitempty"`
	YAxisName  string `json:"yAxisName,omitempty"`
}

//Brush是一个区域选择组件，用户可以使用它从图表中选择部分数据以详细显示，或者使用它们进行计算。
// Brush is an area-selecting component, with which user can select part of data from a chart to display in detail, or do calculations with them.
// https://echarts.apache.org/en/option.html#brush

type Brush struct {
	//XAxisIndex Assigns which of the xAxisIndex can use brush selecting.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
	//Brushlink is a mapping of dataIndex. So data of every series with brushLink should be guaranteed to correspond to the other.

	Brushlink interface{} `json:"brushlink,omitempty"`
	//OutOfBrush Defines visual effects of items out of selection
	OutOfBrush *BrushOutOfBrush `json:"outOfBrush,omitempty"`
}

// BrushOutOfBrush
// https://echarts.apache.org/en/option.html#brush.outOfBrush

type BrushOutOfBrush struct {
	ColorAlpha float32 `json:"ColorAlpha,omitempty"`
}

// Toolbox is the option set for a toolbox component.
// https://echarts.apache.org/en/option.html#toolbox

type Toolbox struct {
	// Whether to show toolbox component.
	Show bool `json:"show"`
	// The layout orientation of toolbox's icon.
	/////工具箱图标的布局方向。
	// Options: 'horizontal','vertical'
	Orient string `json:"orient,omitempty"`

	// Distance between toolbox component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	//工具箱组件与容器左侧之间的距离。
	//左边的值可以是像20这样的即时像素值；也可以是百分比
	//相对于容器宽度的值，如“20%”；也可以是“左”、“中”或“右”。
	//如果左值设置为“左”、“中”或“右”，则组件
	//将根据位置自动对齐。
	Left string `json:"left,omitempty"`
	// Distance between toolbox component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`
	// Distance between toolbox component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`
	// Distance between toolbox component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`
	// The configuration item for each tool.
	// Besides the tools we provide, user-defined toolbox is also supported.
	Feature *ToolBoxFeature `json:"feature,omitempty"`
}

// ToolBoxFeature is a feature component under toolbox.
// https://echarts.apache.org/en/option.html#toolbox
type ToolBoxFeature struct {
	// Save as image tool
	SaveAsImage *ToolBoxFeatureSaveAsImage `json:"saveAsImage,omitempty"`
	// Save as image tool
	Brush *ToolBoxFeatureBrush `json:"brush,omitempty"`
	// Data area zooming, which only supports rectangular coordinate by now.
	//数据区域缩放，目前只支持直角坐标。
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`
	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`
	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
// https://echarts.apache.org/en/option.html#toolbox.feature.saveAsImage

type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show bool `json:"Show"`
	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default),
	// then 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when  chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).

	Type string `json:"Type,omitempty"`
	// Name to save the image, whose default value is title.text.
	Name string `json:"Name,omitempty"`
	// Title for the tool.
	Title string `json:"Title,omitempty"`
}

// ToolBoxFeatureBrush  brush-selecting icon.
// https://echarts.apache.org/en/option.html#toolbox.feature.brush

type ToolBoxFeatureBrush struct {
	//Icons used, whose values are:
	// 'rect': Enabling selecting with rectangle area. // “矩形”  能够选择矩形区域
	// 'polygon': Enabling selecting with any shape.
	// 'lineX': Enabling horizontal selecting.  // horizontal  水平的
	// 'lineY': Enabling vertical selecting.   vertical  竖的
	// 'keep': Switching between single selecting and multiple selecting. The latter one can select multiple areas, while the former one cancels previous selection.
	// 'clear': Clearing all selection.
	Type []string `json:"Type,omitempty"`
}

// ToolBoxFeatureDataZoom
// https://echarts.apache.org/en/option.html#toolbox.feature.dataZoom

type ToolBoxFeatureDataZoom struct {

	// Whether to show the tool.
	Show bool `json:"Show"`
	//Defines which yAxis should be controlled. By default, it controls all y axes.
	//If it is set to be false, then no y axis is controlled.
	//If it is set to be then it controls axis with axisIndex of 3.
	//If it is set to be [0, 3], it controls the x-axes with axisIndex of 0 and 3.
	YAxisIndex interface{} `json:"YAxisIndex,omitempty"`
	// Restored and zoomed title text.
	// m["zoom"] = "area zooming"
	// m["back"] = "restore area zooming"
	//恢复并缩放标题文本。
	//m[“zoom”]=“区域缩放”
	//m[“back”]=“恢复区域缩放”

	Title map[string]string `json:"Title"`
}

// ToolBoxFeatureDataView
// https://echarts.apache.org/en/option.html#toolbox.feature.dataView

type ToolBoxFeatureDataView struct {
	// Whether to show the tool.
	Show bool `json:"Show"`
	// title for the tool.
	Title string `json:"Title,omitempty"`
	// There are 3 names in data view
	// you could set them like this: []string["data view", "turn off", "refresh"]
	Lang []string `json:"Lang"`
	// Background color of the floating layer in data view.
	BackgroundColor string `json:"BackgroundColor"`
}

// ToolBoxFeatureRestore
// https://echarts.apache.org/en/option.html#toolbox.feature.restore
type ToolBoxFeatureRestore struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}

// AxisLabel settings related to axis label.
// https://echarts.apache.org/en/option.html#xAxis.axisLabel

type AxisLabel struct {
	// Set this to false to prevent the axis label from appearing.
	// 将此设置为false可防止出现轴标签。
	Show bool `json:"Show"`
	// Set this to true so the axis labels face the inside direction.
	//将此设置为true，以便轴标签面向内部方向。
	Interval string `json:"interval,omitempty"`
	//轴标签的旋转度，当类别轴没有足够的空间时，这一点特别有用。
	//旋转角度为-90到90。
	Rotate float64 `json:"rotate,omitempty"`
	// The margin(边缘） between the axis label and the axis line.
	Margin float64 `json:"margin,omitempty"`

	//Formatter of axis label, which supports string template and callback function.
	////轴标签的格式化程序，支持字符串模板和回调函数。
	// Example:
	//
	// Use string template; template variable is the default label of axis {value}
	//使用字符串模板；模板变量是轴｛value｝的默认标签
	// formatter: '{value} kg'
	//
	// Use callback function; function parameters are axis index
	//
	//
	//  formatter: function (value, index) {
	//    // Formatted to be month/day; display year only in the first label
	//    var date = new Date(value);
	//    var texts = [(date.getMonth() + 1), date.getDate()];
	//    if (idx === 0) {
	//        texts.unshift(date.getYear());
	//    }
	//    return texts.join('/');
	// }

	Color string `json:"color,omitempty"`
	// axis label font style   //轴标签字体样式
	FontStyle string `json:"fontStyle,omitempty"`
	// axis label font weight
	FontWeight string `json:"fontWeight,omitempty"`
	// axis label font family
	FontFamily string `json:"fontFamily,omitempty"`
	// axis label font size
	FontSize string `json:"fontSize,omitempty"`
	// Horizontal alignment of axis label  //轴标签的水平对齐
	Align string `json:"align,omitempty"`
	// Vertical alignment of axis label
	VerticalAlign string `json:"verticalAlign,omitempty"`
	// Line height of the axis label
	//轴标签的线高度
	LineHeight string `json:"lineHeight,omitempty"`
}
type AxisTick struct {
	// Set this to false to prevent the axis tick from showing.
	//将此设置为false可防止显示轴刻度。
	Show bool `json:"show"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	//指定X轴还是Y轴位于另一个的原点位置，其中轴上的值为0。
	//仅当另一个轴为值类型并且包含0值时有效。
	OnZero bool `json:"onZero,omitempty"`

	// When multiple axes exists, this option can be used to specify which axis can be "onZero" to.
	//当存在多个轴时，此选项可用于指定哪个轴可以“onZero”到。
	OnZeroAxisIndex int `json:"onZeroAxisIndex,omitempty"`
	// Symbol of the two ends of the axis. It could be a string, representing the same symbol for two ends; or an array
	// with two string elements, representing the two ends separately. It's set to be 'none' by default, meaning no
	//arrow for either end. If it is set to be 'arrow', there shall be two arrows. If there should only one arrow
	//at the end, it should set to be ['none', 'arrow'].
	Symbol string `json:"symbol,omitempty"`

	// Size of the arrows at two ends. The first is the width perpendicular to the axis, the next is the width parallel to the axis.
	//两端箭头的大小。第一个是垂直于轴的宽度，第二个是平行于轴的长度。
	SymbolSize []float64 `json:"symbolSize,omitempty"`

	// Arrow offset of axis. If is array, the first number is the offset of the arrow at the beginning, and the second
	// number is the offset of the arrow at the end. If is number, it means the arrows have the same offset.
	//轴的箭头偏移。如果是数组，则第一个数字是箭头开头的偏移量，第二个数字是
	//数字是箭头末端的偏移量。如果是数字，则表示箭头具有相同的偏移量。
	SymbolOffset []float64 `json:"symbolOffset,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// AxisLine controls settings related to axis line.
// https://echarts.apache.org/en/option.html#yAxis.axisLine
type AxisLine struct {
	// Set this to false to prevent the axis line from showing.
	Show bool `json:"show"`

	// Specifies whether X or Y axis lies on the other's origin position, where value is 0 on axis.
	// Valid only if the other axis is of value type, and contains 0 value.
	OnZero bool `json:"onZero,omitempty"`

	// When multiple axes exists, this option can be used to specify which axis can be "onZero" to.
	OnZeroAxisIndex int `json:"onZeroAxisIndex,omitempty"`

	// Symbol of the two ends of the axis. It could be a string, representing the same symbol for two ends; or an array
	// with two string elements, representing the two ends separately. It's set to be 'none' by default, meaning no
	//arrow for either end. If it is set to be 'arrow', there shall be two arrows. If there should only one arrow
	//at the end, it should set to be ['none', 'arrow'].
	Symbol string `json:"symbol,omitempty"`

	// Size of the arrows at two ends. The first is the width perpendicular to the axis, the next is the width parallel to the axis.
	SymbolSize []float64 `json:"symbolSize,omitempty"`

	// Arrow offset of axis. If is array, the first number is the offset of the arrow at the beginning, and the second
	// number is the offset of the arrow at the end. If is number, it means the arrows have the same offset.
	SymbolOffset []float64 `json:"symbolOffset,omitempty"`

	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// XAxis is the option set for X axis.
// https://echarts.apache.org/en/option.html#xAxis
type XAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`
	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.  // 适合离散类别数据
	//   Category data can be auto retrieved(检索） from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.

	Type string `json:"type,omitempty"`
	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`
	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`
	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	//仅在数值轴中可用，即类型：“value”。
	//指定是否强制包含轴的零位。
	//当它被设置为真时，轴可以不包含零位置，
	//这在两个值轴的散点图中都是有用的。
	//当设置了最小值和最大值时，此配置项不可用。
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	//轴的最小值。
	//可以将其设为特殊值“dataMin”，以便将该轴上的最小值设置为最小标签。
	//它将自动计算，以确保未设置时轴刻度均匀分布。
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`
	// Minimum gap between split lines. For 'time' axis, MinInterval is in unit of milliseconds.
	//分割线之间的最小间隙。对于“时间”轴，MinInterval以毫秒为单位。
	MinInterval float64 `json:"minInterval,omitempty"`
	// Maximum gap between split lines. For 'time' axis, MaxInterval is in unit of milliseconds.
	MaxInterval float64 `json:"maxInterval,omitempty"`

	// The index of grid which the x axis belongs to. Defaults to be in the first grid.
	// default 0
	//x轴所属网格的索引。默认位于第一个网格中。

	//默认值0
	GridIndex int `json:"gridIndex,omitempty"`
	// Split area of X axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`
	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`
	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`
	// Settings related to axis tick.
	AxisTick *AxisTick `json:"axisTick,omitempty"`
}

// YAxis is the option set for Y axis.
// https://echarts.apache.org/en/option.html#yAxis
type YAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// The index of grid which the Y axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of Y axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of Y axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Settings related to axis label.
	AxisLabel *AxisLabel `json:"axisLabel,omitempty"`

	// Settings related to axis line.
	AxisLine *AxisLine `json:"axisLine,omitempty"`
}

// TextStyle is the option set for a text style component.
////TextStyle是为文本样式组件设置的选项。

type TextStyle struct {
	// Font color
	Color string `json:"color,omitempty"`
	// Font style
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`
	// Font size
	FontSize int `json:"fontSize,omitempty"`
	// Font family the main title font family.
	// Options: "sans-serif", 'serif' , 'monospace', 'Arial', 'Courier New', 'Microsoft YaHei', ...
	FontFamily string `json:"fontFamily,omitempty"`
	// Padding title space around content.
	// The unit is px. Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	Padding interface{} `json:"padding,omitempty"`

	// compatible for WordCloud
	//兼容WordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}

// SplitArea is the option set for a split area.

type SplitArea struct {
	// Set this to true to show the splitArea.
	//将其设置为true以显示splitArea。
	Show bool `json:"show"`
	// Split area style.
	AreaStyle *AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show bool `json:"show"`
	// Split line style.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
	// Align split line with label, which is available only when boundaryGap is set to be true in category axis.
	//将分割线与标签对齐，只有在类别轴中将boundaryGap设置为true时，标签才可用。
	AlignWithLabel bool `json:"alignWithLabel,omitempty"`
}

// Used to customize how to slice continuous data, and some specific view style for some pieces.
//用于自定义如何对连续数据进行切片，以及某些片段的某些特定视图样式。

type Piece struct {
	Min float32 `json:"min,omitempty"`

	Max float32 `json:"max,omitempty"`

	Lt float32 `json:"lt,omitempty"`

	Lte float32 `json:"lte,omitempty"`

	Gt float32 `json:"gt,omitempty"`

	Gte float32 `json:"gte,omitempty"`
	// Symbol color
	Color string `json:"color,omitempty"`
}

// VisualMap is a type of component for visual encoding, which maps the data to visual channels.
//VisualMap是一种用于视觉编码的组件，它将数据映射到视觉通道。
// https://echarts.apache.org/en/option.html#visualMap

type VisualMap struct {
	// Mapping type.
	// Options: "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`
	// Whether show handles, which can be dragged to adjust "selected range".
	Calculable bool `json:"calculable"`
	// Specify the min dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	//指定visualMap组件的最小数据值。
	//[visualMap.min，visualMax.max]构成了视觉映射的领域。
	Min float32 `json:"min,omitempty"`
	// Specify the max dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Max float32 `json:"max,omitempty"`
	// Specify selected range, that is, the dataValue corresponding to the two handles.
	Range []float32 `json:"range,omitempty"`
	// The label text on both ends, such as ['High', 'Low'].
	Text []string `json:"text,omitempty"`
	// Specify which dimension should be used to fetch dataValue from series.data, and then map them to visual channel.
	//指定应使用哪个维度从series.data中获取dataValue，然后将它们映射到视觉通道。
	Dimension string `json:"dimension,omitempty"`
	// Define visual channels that will mapped from dataValues that are in selected range.
	InRange string `json:"inRange,omitempty"`

	// Used to customize how to slice continuous data, and some specific view style for some pieces.
	//用于自定义如何对连续数据进行切片，以及某些片段的某些特定视图样式。
	Pieces []Piece `json:"pieces,omitempty"`

	// Whether to show visualMap-piecewise component. If set as false,
	// visualMap-piecewise component will not show,
	// but it can still perform visual mapping from dataValue to visual channel in chart.
	//是否显示visualMap分段组件。如果设置为假，
	//visualMap分段组件将不会显示，
	//但它仍然可以执行从dataValue到图表中可视通道的可视映射。
	Show bool `json:"show"`
	// Distance between visualMap component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	//visualMap组件与容器左侧之间的距离。
	//左边的值可以是像20这样的即时像素值；也可以是百分比
	//相对于容器宽度的值，如“20%”；也可以是“左”、“中”或“右”。
	//如果左值设置为“左”、“中”或“右”，
	//则部件将基于位置自动对准。
	Left string `json:"left.,omitempty"`
	// Distance between visualMap component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between visualMap component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between visualMap component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

type VisualMapInRange struct {
	//Color
	Color []string `json:"color,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	//标记线两端的符号类型。它可以是两端的数组，也可以单独指定。
	//选项：“圆形”、“矩形”、“圆形”，“三角形”、“菱形”、“销”、“箭头”、“无”
	Symbol string `json:"symbol,omitempty"`
	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// DataZoom is the option set for a zoom component.
// dataZoom component is used for zooming a specific area, which enables user to
// investigate data in detail, or get an overview of the data, or get rid of outlier points.
// DataZoom是为缩放组件设置的选项。
// dataZoom组件用于缩放特定区域，使用户可以 详细调查数据，或者对数据进行概述，或者消除异常点。
// https://echarts.apache.org/en/option.html#dataZoom
type DataZoom struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type string `json:"type"default:"inside"`
	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	//窗口超出数据范围的起始百分比，范围为0~100。
	//默认值0
	Start float32 `json:"start,omitempty"`
	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`
	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	//指定视图刷新的帧速率，单位为毫秒（ms）。
	//如果animation设置为true并且animationDurationUpdate设置为大于0，
	//您可以将throttle保持为默认值100（或将其设置为大于0的值），
	//否则，拖动时可能会不平滑。
	//如果动画设置为false或animationDurationUpdate设置为0并且数据大小不是很大，
	//并且拖动时似乎不平滑，可以将throttle设置为0来改善这一点。

	Throttle float32 `json:"throttle,omitempty"`
	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	//指定使用笛卡尔坐标系时，哪些xAxis由内部的dataZoom控制。

	//默认情况下，当dataZoom在内部时，将控制与dataZoom平行的第一个xAxis。

	//方向设置为“水平”。但建议明确指定它，但不要使用默认值。

	//如果设置为单个数字，则控制一个轴，而如果设置为数组，

	//控制多个轴。
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.

	// y 轴

	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}

// SingleAxis is the option set for single axis.
// https://echarts.apache.org/en/option.html#singleAxis
type SingleAxis struct {

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	//轴的最小值。
	//可以将其设置为特殊值“dataMin”，以便将该轴上的最小值设置为最小标签。
	//它将自动计算，以确保未设置时轴刻度均匀分布。
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.

	Type string `json:"type,omitempty"`
	// Distance between grid component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	//栅格组件与容器左侧之间的距离。

	//左边的值可以是像20这样的即时像素值；也可以是百分比

	//相对于容器宽度的值，如“20%”；也可以是“左”、“中”或“右”。

	//如果左值设置为“左”、“中”或“右”，

	//则部件将基于位置自动对准。
	Left string `json:"left,omitempty"`
	// Distance between grid component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`
	// Distance between grid component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// Indicator is the option set for a radar chart.

type Indicator struct {
	// Indicator name
	Name string `json:"name,omitempty"`

	// The maximum value of indicator. It is an optional configuration, but we recommend to set it manually.
	//指示器的最大值。这是一个可选配置，但我们建议手动设置。

	Max float32 `json:"max,omitempty"`
	// The minimum value of indicator. It is an optional configuration, with default value of 0.
	Min float32 `json:"min,omitempty"`

	// Specify a color the indicator.

	Color string `json:"color,omitempty"`
}

// RadarComponent is the option set for a radar component.
// https://echarts.apache.org/en/option.html#radar
type RadarComponent struct {

	// Indicator of radar chart, which is used to assign multiple variables(dimensions) in radar chart.
	//radar的指示符，用于分配radar中的多个变量（维度）。
	Indicator []*Indicator `json:"indicator,omitempty"`
	// Radar render type, in which 'polygon' and 'circle' are supported.
	//radar 渲染类型，其中支持“多边形”和“圆形”。
	Shape string `json:"shape,omitempty"`
	// Segments of indicator axis.// 指使轴的分段
	// default 5

	SplitNumber int `json:"splitNumber,omitempty"`

	// Center position of , the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width and height.

	Center interface{} `json:"center,omitempty"`
	// Split area of axis in grid area.
	//轴在网格区域中的拆分区域。
	SplitArea *SplitArea `json:"splitArea,omitempty"`
	// Split line of axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`
}

// GeoComponent is the option set for geo component.
// https://echarts.apache.org/en/option.html#geo
type GeoComponent struct {
	// Map charts.
	Map string `json:"map,omitempty"`
	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	//地图区域边界的图形样式，强调是突出显示时的样式，
	//比如鼠标悬停，或者通过图例连接突出显示。
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
	// Set this to true, to prevent interaction with the axis.
	Silent bool `json:"silent,omitempty"`
}

type ParallelComponent struct {

	// Distance between parallel component and the left side of the container.
	// Left value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%';
	// and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	//平行构件与容器左侧之间的距离。
	//左边的值可以是像20这样的即时像素值。
	//它也可以是相对于容器宽度的百分比值，如“20%”；
	//也可以是“左”、“中”或“右”。
	//如果左值设置为“左”、“中”或“右”，
	//则部件将基于位置自动对准。
	Left string `json:"left,omitempty"`
	// Distance between parallel component and the top side of the container.
	// Top value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	// and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`
	// Distance between parallel component and the right side of the container.
	// Right value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between parallel component and the bottom side of the container.
	// Bottom value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

//ParallelAxis is the  option  set for a parallel axis

type ParallelAxis struct {
	// Dimension index of coordinate axis.
	Dim int `json:"dim,omitempty"`
	// Name of axis.
	Name string `json:"name,omitempty"`
	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	// In category axis, it can also be set as the ordinal number.
	//轴的最大值。
	//可以将其设置为特殊值“dataMax”，以便将此轴上的最小值设置为最大标签。
	//它将自动计算，以确保未设置时轴刻度均匀分布。
	//在类别轴中，它也可以设置为序号。
	Max interface{} `json:"max,omitempty"`
	// Compulsively set segmentation interval for axis.
	//强制设置轴的分割间隔。
	Inverse bool `json:"inverse,omitempty"`
	// Location of axis name. Options: "start", "middle", "center", "end"
	// default "end"
	NameLocation string `json:"nameLocation,omitempty"`
	// Type of axis.
	// Options：
	// "value"：Numerical axis, suitable for continuous data.
	// "category" Category axis, suitable for discrete category data. Category data can be auto retrieved from series.
	// "log"  Log axis, suitable for log data.
	Type string `json:"type,omitempty"`
	// Category data，works on (type: "category").
	Data interface{} `json:"data"`
}

// Polar Bar options
type Polar struct {
	ID      string    `json:"id,omitempty"`
	Zlevel  int       `json:"zlevel,omitempty"`
	Z       int       `json:"z,omitempty"`
	Center  [2]string `json:"center,omitempty"`
	Radius  [2]string `json:"radius,omitempty"`
	Tooltip Tooltip   `json:"tooltip,omitempty"`
}

type PolarAxisBase struct {
	ID           string  `json:"id,omitempty"`
	PolarIndex   int     `json:"polarIndex,omitempty"`
	StartAngle   float64 `json:"startAngle,omitempty"`
	Type         string  `json:"type,omitempty"`
	BoundaryGap  bool    `json:"boundaryGap,omitempty"`
	Min          float64 `json:"min,omitempty"`
	Max          float64 `json:"max,omitempty"`
	Scale        bool    `json:"scale,omitempty"`
	SplitNumber  int     `json:"splitNumber,omitempty"`
	MinInterVal  float64 `json:"minInterVal,omitempty"`
	MaxInterval  float64 `json:"maxInterval,omitempty"`
	InterVal     float64 `json:"interVal,omitempty"`
	LogBase      float64 `json:"logBase,omitempty"`
	Silent       float64 `json:"silent,omitempty"`
	TriggerEvent bool    `json:"triggerEvent,omitempty"`
}

type AngleAxis struct {
	PolarAxisBase
	Clockwise bool `json:"clockwise,omitempty"`
}

type RadiusAxis struct {
	PolarAxisBase
	Name          string    `json:"name,omitempty"`
	NameLocation  string    `json:"nameLocation,omitempty"`
	NameTestStyle TextStyle `json:"nameTextStyle,omitempty"`
	NameGap       float64   `json:"nameGap,omitempty"`
	NameRadius    float64   `json:"nameRadius,omitempty"`
	Inverse       bool      `json:"inverse,omitempty"`
}

var funcPat = regexp.MustCompile(`\n|\t`)

const funcMarker = "__f__"

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.

func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, funcPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts is the option set for handling function type.

func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

func replaceJsFuncs(fn string) string {
	return fmt.Sprintf("%s%s%s", funcMarker, funcPat.ReplaceAllString(fn, ""), funcMarker)
}

type Colors []string

type Assets struct {
	JSAssets  types.OrderedSet
	CSSAssets types.OrderedSet

	CustomizedJSAssets  types.OrderedSet
	CustomizedCSSAssets types.OrderedSet
}

// InitAssets inits the static assets storage.

func (opt *Assets) InitAssets() {

	opt.JSAssets.Init("echars.min.js")
	opt.CSSAssets.Init()
	opt.CustomizedJSAssets.Init()
	opt.CustomizedCSSAssets.Init()
}

// AddCustomizedJSAssets adds the customized javascript assets which will not be added the `host` prefix.
// AddCustomizedJSAssets添加自定义的javascript资产，这些资产不会添加“host”前缀。
func (opt *Assets) AddCustomizedJSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedJSAssets.Add(assets[i])
	}
}

// AddCustomizedCSSAssets adds the customized css assets which will not be added the `host` prefix.

func (opt *Assets) AddCustomizedCSSAssets(assets ...string) {
	for i := 0; i < len(assets); i++ {
		opt.CustomizedCSSAssets.Add(assets[i])
	}
}

// Validate validates the static assets configurations

func (opt *Assets) Validate(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		if !strings.HasPrefix(opt.JSAssets.Values[i], host) {
			opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
		}
	}

	for i := 0; i < len(opt.CSSAssets.Values); i++ {
		if !strings.HasPrefix(opt.CSSAssets.Values[i], host) {
			opt.CSSAssets.Values[i] = host + opt.CSSAssets.Values[i]
		}
	}
}

// XAxis3D contains options for X axis in the 3D coordinate.

type XAxis3D struct {
	// Whether to display the x-axis.
	Show bool `json:"show,omitempty"`
	// The name of the axis.
	Name string `json:"name,omitempty"`
	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	//轴使用的栅格3D组件的索引。默认情况是使用第一个grid3D构件。
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	/// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	///轴的类型。
	//可选：
	//*'value'值轴。适用于连续数据。
	//*'category'类别轴。适用于离散类别数据。
	//对于这种类型，必须通过数据设置类别数据。
	//*“时间”时间线。适用于连续计时数据。时间轴有一个
	//时间格式与数值轴相比，尺度计算也有所不同。
	//例如，月、周、日和小时范围的比例可以根据跨度的范围来确定。
	//*“log”对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.

	//轴的最小值
	//可以将其设置为特殊值“dataMin”，以便将该轴上的最小值设置最小标签。
	//它将自动计算，以确保未设置时轴刻度均匀分布。
	//在类别轴中，它也可以设置为序号。例如
	//如果类别轴具有数据：[类别a，类别B，类别C]，
	//序号2表示“类别C”。此外，它可以设置为负数，如-3。
	Min interface{} `json:"min,omitempty"`
	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`
	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// YAxis3D contains options for Y axis in the 3D coordinate.
type YAxis3D struct {
	// Whether to display the y-axis.
	Show bool `json:"show,omitempty"`

	// The name of the axis.
	Name string `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// ZAxis3D contains options for Z axis in the 3D coordinate.
type ZAxis3D struct {
	// Whether to display the z-axis.
	Show bool `json:"show,omitempty"`

	// The name of the axis.
	Name string `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// Grid3D contains options for the 3D coordinate.
//Grid3D包含三维坐标的选项。

type Grid3D struct {
	// Whether to show the coordinate.
	Show bool `json:"show,omitempty"`
	// 3D Cartesian coordinates width
	// default 100
	BoxWidth float32 `json:"boxWidth,omitempty"`
	// 3D Cartesian coordinates height
	// default 100
	BoxHeight float32 `json:"boxHeight,omitempty"`
	// 3D Cartesian coordinates depth
	// default 100
	BoxDepth float32 `json:"boxDepth,omitempty"`
	// Rotate or scale fellows the mouse  //旋转或缩放鼠标
	ViewControl *ViewControl `json:"viewControl,omitempty"`
}

// ViewControl contains options for view controlling.
type ViewControl struct {
	// Whether to rotate automatically.
	AutoRotate bool `json:"autoRotate,omitempty"`
	// Rotate Speed, (angle/s).
	// default 10
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}

// Grid Drawing grid in rectangular coordinate.
// https://echarts.apache.org/en/option.html#grid

type Grid struct {
	// Distance between grid component and the left side of the container.
	Left string `json:"left,omitempty"`
	// Distance between grid component and the right side of the container
	Right string `json:"right,omitempty"`
	// Distance between grid component and the top side of the container.
	Top string `json:"top,omitempty"`
	// Distance between grid component and the bottom side of the container.
	Bottom string `json:"bottom,omitempty"`
	// Width of grid component.
	Width string `json:"width,omitempty"`

	ContainLabel bool `json:"containLabel,omitempty"`
	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`
}

// Dataset brings convenience in data management separated with styles and enables data reuse by different series.
// More importantly, it enables data encoding from data to visual, which brings convenience in some scenarios.
// https://echarts.apache.org/en/option.html#dataset.id
//数据集为以样式分隔的数据管理带来了便利，并允许不同系列的数据重用。
//更重要的是，它实现了从数据到可视化的数据编码，这在某些场景中带来了便利。

type Dataset struct {
	//source
	Source interface{} `json:"source"`
}
