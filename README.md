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
