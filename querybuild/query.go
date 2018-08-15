package querybuild

import (
	"fmt"
	"strings"
)

func Build(fs ...Fragment) string {
	var parts []string
	for _, f := range fs {
		parts = append(parts, f())
	}
	return strings.Join(parts, " ")
}

type Fragment func() string

func Org(name string) Fragment {
	return func() string {
		return fmt.Sprintf("org:%s", name)
	}
}

func Filename(name string) Fragment {
	return func() string {
		return fmt.Sprintf("filename:%s", name)
	}
}

func Query(q string) Fragment {
	return func() string {
		return q
	}
}
