package mysql

import(
	"strings"
	"strconv"
)
type Search struct{
	sel []string
	from string
	where []string
	order []string
	group []string
	limit int
	offset int
}

/** 查询表名 **/
func (s *Search) From(table string) *Search {
	s.from = table
	return s
}

/** 查询字段  **/
func (s *Search) Field(field string) *Search {
	s.sel = append(s.sel, field)
	return s
}

/** 查询条件 **/
func (s *Search) Where(condition string) *Search {
	s.where = append(s.where, condition)
	return s
}

/** 排序 **/
func (s *Search) Order(orderBy string) *Search {
	s.order = append(s.order, orderBy)
	return s
}

/** 分组 **/
func (s *Search) Group(groupBy string) *Search {
	s.group = append(s.group, groupBy)
	return s
}

/** limit **/
func (s *Search) Limit(limit int) *Search {
	s.limit = limit
	return s
}

/** offset **/
func (s *Search) Offset(offset int) *Search {
	s.offset = offset
	return s
}

/** sql拼接 **/
func (s *Search) ToString() string {
	var sel, from, where, order, group, limit, offset string

	sel = "SELECT "
	if len(s.sel) == 0 {
		sel = sel + "*"
	}else{
		sel = sel + strings.Join(s.sel, ",")
	}
	
	from = " FROM "+ s.from
	
	if len(s.where) != 0 {
		where = " WHERE "
		where = where + strings.Join(s.where, " AND ")
	}

	if len(s.order) != 0 {
		order = " ORDER BY "
		order = order + strings.Join(s.order, ",")
	}
	
	if len(s.group) != 0 {
		group = " GROUP BY "
		group = group + strings.Join(s.group, ",")
	}

	if s.limit != 0 {
		limit = " LIMIT " + strconv.Itoa(s.limit)
	}

	if s.offset != 0 {
		offset = " OFFSET " + strconv.Itoa(s.offset)
	}

	return sel + from + where + order + group + limit + offset
}