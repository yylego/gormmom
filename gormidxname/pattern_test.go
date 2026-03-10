package gormidxname

import (
	"testing"

	"github.com/yylego/gormmom/internal/simpleindexname"
	"github.com/yylego/must"
)

func TestPattern(t *testing.T) {
	must.Same(simpleindexname.IdxPatternTagName, IdxPatternTagName)
	must.Same(simpleindexname.UdxPatternTagName, UdxPatternTagName)
}
