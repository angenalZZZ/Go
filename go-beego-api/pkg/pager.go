package pkg

import (
	"github.com/go-xorm/builder"
	"github.com/xormplus/xorm"
)

type Pager struct {
	*builder.Builder
}

func NewPagingBuilder(build func(builder *Pager)) *Pager {
	b := &Pager{builder.Dialect(builder.MSSQL)}
	if build != nil {
		build(b)
	}
	return b
}

func (b *Pager) Paging(session *xorm.Session, limitN int, offset int, rowsSlicePtr interface{}, condiBean ...interface{}) error {
	if sql, args, err := b.ToSQL(); err != nil {
		return err
	} else {
		return session.SQL(sql+" OFFSET ? ROW FETCH NEXT ? ROW ONLY", append(args, offset, limitN)...).Find(rowsSlicePtr, condiBean...)
	}
}
