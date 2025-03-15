
# DevChart - Terminal Charting Library in Go

**DevChart** is a simple Go library for generating tables and basic charts (bar graphs and line graphs) directly in your terminal. It's designed for command-line applications and tools where you need to display data in a visually informative way without relying on graphical interfaces.

## Features

*   **Table Generation:**
    *   Render data in formatted tables with customizable borders, text alignment, and padding.
    *   Supports single and double line borders, or no borders at all.
    *   Left, center, and right text alignment within table cells.
    *   Adjustable cell padding.
    *   Automatic column width calculation for clean formatting.
*   **Bar Graph Generation:**
    *   Create simple vertical bar graphs in the terminal.
    *   Customize graph width and the character used for bars.
    *   Optional labels for each bar on the X-axis.
*   **Line Graph Generation:**
    *   Generate basic text-based line graphs.
    *   Customize graph width, height, and the character used for data points.
    *   Optional labels for data points on the X-axis.

## Installation

To install DevChart, use `go get`:

```bash
go get github.com/Ziqian-Huang0607/devchart
```

Make sure you have Go installed and your `GOPATH` environment variable set up correctly.

## Usage

To use DevChart in your Go project, follow these steps:

**1. Import the necessary packages:**

   Depending on whether you want to use tables, bar graphs, or line graphs, import the relevant packages from DevChart in your Go file:

   ```go
   import (
       "fmt"
       "os"

       "github.com/Ziqian-Huang0607/devchart/table" // For tables
       "github.com/Ziqian-Huang0607/devchart/graph" // For graphs (bar and line)
   )
   ```

**2. Create a Chart Object:**

   *   **Tables:** Use `table.NewTable()` to create a new `Table` object.
       ```go
       myTable := table.NewTable()
       ```

   *   **Bar Graphs:** Use `graph.NewBarGraph()` to create a new `BarGraph` object.
       ```go
       barGraph := graph.NewBarGraph()
       ```

   *   **Line Graphs:** Use `graph.NewLineGraph()` to create a new `LineGraph` object.
       ```go
       lineGraph := graph.NewLineGraph()
       ```

**3. Set Data:**

   *   **Tables:** Use `SetHeaders([]string)` (optional) to set table headers and `SetData([][]string)` to provide the table data as a slice of string slices.
       ```go
       headers := []string{"Header 1", "Header 2", "Header 3"}
       data := [][]string{
           {"Row 1, Col 1", "Row 1, Col 2", "Row 1, Col 3"},
           {"Row 2, Col 1", "Row 2, Col 2", "Row 2, Col 3"},
           // ... more rows
       }
       myTable.SetHeaders(headers) // Optional: Set headers
       myTable.SetData(data)
       ```

   *   **Bar Graphs:** Use `SetData([]float64)` to provide the data values for the bars and `SetLabels([]string)` (optional) to set labels for each bar.
       ```go
       barGraphData := []float64{25, 40, 30, 55, 45}
       barGraphLabels := []string{"Category A", "Category B", "Category C", "Category D", "Category E"} // Optional labels
       barGraph.SetData(barGraphData)
       barGraph.SetLabels(barGraphLabels) // Optional: Set labels
       ```

   *   **Line Graphs:** Use `SetData([]float64)` to provide the data points for the line graph and `SetLabels([]string)` (optional) to set labels for each data point (typically X-axis labels).
       ```go
       lineGraphData := []float64{10, 15, 8, 12, 18, 14, 20}
       lineGraphLabels := []string{"Point 1", "Point 2", "Point 3", "Point 4", "Point 5", "Point 6", "Point 7"} // Optional labels
       lineGraph.SetData(lineGraphData)
       lineGraph.SetLabels(lineGraphLabels) // Optional: Set labels
       ```

**4. Customize Options (Optional):**

   You can customize the appearance of your tables and graphs using option structs and their methods.

   *   **Tables:** Use `table.DefaultOptions()` to get a default `Options` struct, then use methods like `SetBorderStyle()`, `SetAlignment()`, and `SetPadding()` to modify it. Finally, apply the options using `myTable.SetOptions(options)`.
       ```go
       tableOptions := table.DefaultOptions()
       tableOptions.SetBorderStyle(table.BorderDouble).SetAlignment(table.AlignCenter).SetPadding(2)
       myTable.SetOptions(tableOptions)
       ```
       See the [Table Options](#table-options) section below for available options.

   *   **Bar Graphs:** Use `graph.DefaultBarGraphOptions()` to get a default `BarGraphOptions` struct and methods like `SetWidth()` and `SetBarCharacter()` to customize. Apply using `barGraph.SetOptions(barGraphOptions)`.
       ```go
       barGraphOptions := graph.DefaultBarGraphOptions()
       barGraphOptions.SetWidth(60).SetBarCharacter("=")
       barGraph.SetOptions(barGraphOptions)
       ```
       See the [Bar Graph Options](#bar-graph-options) section below for available options.

   *   **Line Graphs:** Use `graph.DefaultLineGraphOptions()` to get a default `LineGraphOptions` struct and methods like `SetWidth()`, `SetHeight()`, and `SetPointCharacter()` to customize. Apply using `lineGraph.SetOptions(lineGraphOptions)`.
       ```go
       lineGraphOptions := graph.DefaultLineGraphOptions()
       lineGraphOptions.SetWidth(70).SetHeight(12).SetPointCharacter("*")
       lineGraph.SetOptions(lineGraphOptions)
       ```
       See the [Line Graph Options](#line-graph-options) section below for available options.

**5. Render and Print:**

   Use the `Print()` method (or `Render(io.Writer)` for more control) to render the table or graph to the terminal (standard output).

   ```go
   fmt.Println("--- My Chart/Table ---")
   err := myTable.Print() // Or barGraph.Print() or lineGraph.Print()
   if err != nil {
       fmt.Fprintf(os.Stderr, "Error rendering: %v\n", err)
   }
   ```

### <a name="table-options"></a>Table Options

The `table.Options` struct allows you to customize tables with the following methods:

*   `SetBorderStyle(style table.BorderStyle)`: Sets the border style. Available `table.BorderStyle` constants are:
    *   `table.BorderSingle`: Single line borders.
    *   `table.BorderDouble`: Double line borders.
    *   `table.BorderNone`: No borders.
*   `SetAlignment(alignment table.Alignment)`: Sets the text alignment within cells. Available `table.Alignment` constants are:
    *   `table.AlignLeft`: Left alignment.
    *   `table.AlignCenter`: Center alignment.
    *   `table.AlignRight`: Right alignment.
*   `SetPadding(padding int)`: Sets the number of spaces for padding around cell content.

### <a name="bar-graph-options"></a>Bar Graph Options

The `graph.BarGraphOptions` struct allows you to customize bar graphs with the following methods:

*   `SetWidth(width int)`: Sets the maximum width of the graph in characters.
*   `SetBarCharacter(char string)`: Sets the character used to draw the bars.

### <a name="line-graph-options"></a>Line Graph Options

The `graph.LineGraphOptions` struct allows you to customize line graphs with the following methods:

*   `SetWidth(width int)`: Sets the maximum width of the graph in characters.
*   `SetHeight(height int)`: Sets the maximum height of the graph in lines.
*   `SetPointCharacter(char string)`: Sets the character used to represent data points on the line graph.

**Example Program:**

For a complete example demonstrating tables, bar graphs, and line graphs, please refer to the `main.go` file in the repository.

```go
package main

import (
	"fmt"
	"os"

	"github.com/Ziqian-Huang0607/devchart/graph"
	"github.com/Ziqian-Huang0607/devchart/table"
)

func main() {
	// -------------------- Table Examples --------------------
	fmt.Println("-------------------- Table Examples --------------------")

	tableData := [][]string{
		{"Product", "Price", "Quantity", "Description"},
		{"Laptop", "$1200", "50", "High-performance laptop with SSD"},
		{"Mouse", "$25", "200", "Wireless ergonomic mouse"},
		{"Keyboard", "$75", "100", "Mechanical keyboard with RGB"},
		{"Monitor", "$300", "75", "27-inch 4K monitor"},
	}
	headers := []string{"Item", "Cost", "Stock", "Details"} // Different headers for example

	// Example 1: Default Table
	defaultTable := table.NewTable()
	defaultTable.SetHeaders(headers)
	defaultTable.SetData(tableData)
	fmt.Println("\n--- Default Table (Single Border, Left Align, Padding 1) ---")
	defaultTable.Print()

	// Example 2: Double Border, Center Align, Padding 2
	centerTable := table.NewTable()
	centerTable.SetHeaders(headers)
	centerTable.SetData(tableData)
	centerTableOptions := table.DefaultOptions()
	centerTableOptions.SetBorderStyle(table.BorderDouble).SetAlignment(table.AlignCenter).SetPadding(2)
	centerTable.SetOptions(centerTableOptions) // Corrected: Pass value directly
	fmt.Println("\n--- Double Border, Center Align, Padding 2 Table ---")
	centerTable.Print()

	// Example 3: No Borders, Right Align, Padding 0
	noBorderTable := table.NewTable()
	noBorderTable.SetHeaders(headers)
	noBorderTable.SetData(tableData)
	noBorderOptions := table.DefaultOptions()
	noBorderOptions.SetBorderStyle(table.BorderNone).SetAlignment(table.AlignRight).SetPadding(0)
	noBorderTable.SetOptions(noBorderOptions) // Corrected: Pass value directly
	fmt.Println("\n--- No Borders, Right Align, Padding 0 Table ---")
	noBorderTable.Print()

	// -------------------- Bar Graph Example --------------------
	fmt.Println("\n-------------------- Bar Graph Example --------------------")

	barGraphData := []float64{35, 52, 40, 65, 48, 25, 30}
	barGraphLabels := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	barGraph := graph.NewBarGraph()
	barGraph.SetData(barGraphData)
	barGraph.SetLabels(barGraphLabels)

	barGraphOptions := graph.DefaultBarGraphOptions()
	barGraphOptions.SetWidth(60).SetBarCharacter("=") // Customize bar character
	barGraph.SetOptions(barGraphOptions) // Corrected: Pass value directly
	fmt.Println("\n--- Bar Graph (Width 60, Bar Char '=') ---")
	barGraph.Print()

	// Example with different data and labels length
	barGraphData2 := []float64{15, 28, 22, 35}
	barGraphLabels2 := []string{"Q1", "Q2", "Q3", "Q4"}

	barGraph2 := graph.NewBarGraph()
	barGraph2.SetData(barGraphData2)
	barGraph2.SetLabels(barGraphLabels2)
	fmt.Println("\n--- Bar Graph (Shorter Data and Labels) ---")
	barGraph2.Print()

	// -------------------- Line Graph Example --------------------
	fmt.Println("\n-------------------- Line Graph Example --------------------")

	lineGraphData := []float64{10, 15, 8, 12, 18, 14, 20, 16, 22, 19}
	lineGraphLabels := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct"}

	lineGraph := graph.NewLineGraph()
	lineGraph.SetData(lineGraphData)
	lineGraph.SetLabels(lineGraphLabels)

	lineGraphOptions := graph.DefaultLineGraphOptions()
	lineGraphOptions.SetWidth(70).SetHeight(12).SetPointCharacter("*") // Customize width, height, point char
	lineGraph.SetOptions(lineGraphOptions) // Corrected: Pass value directly
	fmt.Println("\n--- Line Graph (Width 70, Height 12, Point Char '*') ---")
	lineGraph.Print()

	// Example with different data and labels length
	lineGraphData2 := []float64{7, 9, 5, 11}
	lineGraphLabels2 := []string{"Week 1", "Week 2", "Week 3", "Week 4"}

	lineGraph2 := graph.NewLineGraph()
	lineGraph2.SetData(lineGraphData2)
	lineGraph2.SetLabels(lineGraphLabels2)
	fmt.Println("\n--- Line Graph (Shorter Data and Labels) ---")
	lineGraph2.Print()
}
```

## Contributing

Contributions are welcome! If you have ideas for improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request on GitHub.

## License

[MIT License](LICENSE)

## Author

[Ziqian-Huang0607](https://github.com/Ziqian-Huang0607)
```
