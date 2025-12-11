package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}

func part1(s []string) int {
	var x1, y1, x2, y2 int
	maxArea := 0
	for i := 0; i < len(s)-1; i++ {
		_, _ = fmt.Sscanf(s[i], "%d,%d", &x1, &y1)
		for j := i + 1; j < len(s); j++ {
			_, _ = fmt.Sscanf(s[j], "%d,%d", &x2, &y2)
			area := int((math.Abs(float64(x2-x1)) + 1) * (math.Abs(float64(y2-y1)) + 1))
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

type Point struct {
	x, y int
}

func getPoints(s []string) []Point {
	var (
		x, y   int
		points []Point
	)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
		points = append(points, Point{x, y})
	}
	return points
}

type Segment struct {
	a, b Point
}

func getSegments(points []Point) []Segment {
	var segments []Segment
	n := len(points)
	for i := range n {
		segments = append(segments, Segment{points[i], points[(i+1)%n]})
	}
	return segments
}

func intersect(segment Segment, p1, p2 Point) bool {
	minX := min(p1.x, p2.x) + 1
	maxX := max(p1.x, p2.x) - 1
	minY := min(p1.y, p2.y) + 1
	maxY := max(p1.y, p2.y) - 1

	segMinX := min(segment.a.x, segment.b.x)
	segMaxX := max(segment.a.x, segment.b.x)
	segMinY := min(segment.a.y, segment.b.y)
	segMaxY := max(segment.a.y, segment.b.y)

	if (segMaxX < minX || segMinX > maxX) || (segMaxY < minY || segMinY > maxY) {
		return false
	}
	return true
}

func part2(s []string) int {
	points := getPoints(s)
	segments := getSegments(points)
	maxArea := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			area := int((math.Abs(float64(points[j].x-points[i].x)) + 1) * (math.Abs(float64(points[j].y-points[i].y)) + 1))
			if area > maxArea {
				valid := true
				for _, segment := range segments {
					if intersect(segment, points[i], points[j]) {
						valid = false
						break
					}
				}
				if valid {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func main() {
	filename := "tests.txt"
	output, err := ReadFile(filename)
	if err != nil {
		println(err)
	}

	fmt.Println("Part1:", part1(output))
	fmt.Println("Part2:", part2(output))
}
