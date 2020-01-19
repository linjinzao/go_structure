package mysql

import (
	"strings"
	util "filesystem/util"
	meta "filesystem/meta"
)

/** insert拼接 **/
func GetInsertSql(fmeta meta.FileMeta, table string) string {
	data := util.StructToMap(fmeta)
	var field []string
	var values []string
	for key, val := range data{
		field = append(field, key)
		values = append(values, val)
	}
	return "INSERT INTO " + table + "(" + strings.Join(field, ",") + ") VALUES (" + strings.Join(values, ",") + ")" 
} 
