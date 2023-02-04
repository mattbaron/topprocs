package line

import (
	"fmt"
	"strings"
)

type Line struct {
	measurement string
	tags        map[string]string
	fields      map[string]string
}

func NewLine(measurement string) *Line {
	line := Line{
		measurement: measurement,
		tags:        make(map[string]string),
		fields:      make(map[string]string),
	}

	return &line
}

func mapToString(hash map[string]string) string {
	items := make([]string, 0)
	for key, value := range hash {
		items = append(items, key+"="+strings.ReplaceAll(value, " ", "\\ "))
	}
	return strings.Join(items, ",")
}

func (line *Line) AddTag(key string, value any) {
	line.tags[key] = fmt.Sprint(value)
}

func (line *Line) AddField(key string, value any) {
	line.fields[key] = fmt.Sprint(value)
}

// TODO: This assumes there is at least 1 tag and 1 field
func (line Line) ToString() string {
	return fmt.Sprintf("%s,%s %s", line.measurement, mapToString(line.tags), mapToString(line.fields))
}

func (line Line) Dump() {
	fmt.Printf("Tags: %s\n", mapToString(line.tags))
	fmt.Printf("Fields: %s\n", mapToString(line.fields))
}
