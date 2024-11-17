package constant

const (
	SimpleStrings    = "+%s\r\n"
	BulkStrings      = "$%d\r\n%s\r\n"
	EmptyBulkStrings = "$%d\r\n\r\n"
	NullBulkStrings = "$-1\r\n"
	Arrays           = "*%d\r\n"
)