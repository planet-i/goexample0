`
type Time struct {
    sec int64
// offset within the second named by Seconds；range [0, 999999999].
    nsec int32
//该时间点是距离 UTC 时间0年1月1日0时 sec 秒、nsec 纳秒的时间点
    loc *Location
//默认程序所在的时区，从操作系统中获得；UTC时间用loc == nil表示
}`


#### Format
func (t Time) Format(layout string) string
#### Parse
func Parse（layout，value string）（Time，error）
- format的逆向操作
- layout(2006-01-02 15:04:05)要和 string 相同的格式
- time.Parse()并不会默认使用本地时区，会尝试从 value 里读出时区信息，当且仅当：有时区信息、时区信息以 zone offset 形式（如+0800）表示、表示结果与本地时区等价时，才会使用本地时区，否则使用读出的时区。若 value 里没有时区信息，则使用 UTC 时间
#### Equal
func (t Time) Equal(u Time) bool
sec和nsec 这两个字段相等，两个time实例指的就是同一个时间点