# Go Validator 简称 GV验证器

GV验证器主要用于添加数据或更新数据时对数据格式进行验证，以下列出GV的使用方法，非常容易理解和上手：


##### 用法

>在数据模型验证数据,如

```go

// 系统配置模型
type SysConfig struct {
	Id        int64
	Title     string    `xorm:"varchar(255)"`
	Name      string    `xorm:"unique"`
	Value     string    `xorm:"varchar(255)"`
	Tag       string    `xorm:"varchar(255)"`
	Sort      int64     `xorm:"tinyint(5)"`
	Status    int64     `xorm:"tinyint(1)"`
	UpdatedAt time.Time `xorm:"updated"`
}

//验证逻辑
func (m SysConfig) Validate() (bool, string) {
	vd := gv.NewValidator()

	vd.Size(m.Title, 1, 20).Message("配置名称1至20个字符!")

	vd.MaxSize(m.Name, 20).Message("配置标识不超过20字符!")
	vd.AlphaDash(m.Name).Message("配置标识由数字、字母和下划线组成!")
	vd.MatchBool(false == m.ExistByName()).Message("配置标识已存在!")

	vd.MaxSize(m.Value, 65535).Message("配置值不超过65535个字符!")

	vd.Require(m.Tag).Message("配置标签不能为空!")
	vd.AlphaDash(m.Tag).Message("配置标签由数字、字母和下划线组成!")
	vd.MaxSize(m.Tag, 20).Message("配置标签不超过20个字符!")

	vd.MinInt64(m.Sort, 0).Message("排序值不能小于0!")
	vd.MaxInt64(m.Sort, 65535).Message("排序值不大于65535!")

	vd.RangeInt(int(m.Status), []int{0, 1}).Message("状态值不正确!")

	return vd.Validate()
}
```

##### 实例化验证器

```go
vd := gv.NewValidator()
```

##### 要求字段值不能为空

```go
vd.Require(value).Message("验证不通过提示信息")
```

##### 验证字符串长度范围

```go
vd.Size(str, minSize, maxSize).Message("验证不通过提示信息")
```

##### 验证字符串最小长度

```go
vd.MinSize(str, minSize).Message("验证不通过提示信息")
```

##### 验证字符串最大长度

```go
vd.MaxSize(str, maxSize).Message("验证不通过提示信息")
```

##### 验证数字不小于某个值

> 同时还提供int8、int16、int32、int64、float32、float64的验证

```go
vd.MinInt(num, min).Message("验证不通过提示信息")
vd.MinInt8(num, min).Message("验证不通过提示信息")
vd.MinInt16(num, min).Message("验证不通过提示信息")
vd.MinInt32(num, min).Message("验证不通过提示信息")
vd.MinInt64(num, min).Message("验证不通过提示信息")
vd.MinFloat32(num, min).Message("验证不通过提示信息")
vd.MinFloat64(num, min).Message("验证不通过提示信息")
```
##### 验证数字不大于某个值

> 同时还提供int8、int16、int32、int64、float32、float64的验证

```go
vd.MaxInt(num, max).Message("验证不通过提示信息")
vd.MaxInt(num, min).Message("验证不通过提示信息")
vd.MaxInt8(num, min).Message("验证不通过提示信息")
vd.MaxInt16(num, min).Message("验证不通过提示信息")
vd.MaxInt32(num, min).Message("验证不通过提示信息")
vd.MaxInt64(num, min).Message("验证不通过提示信息")
vd.MaxFloat32(num, min).Message("验证不通过提示信息")
vd.MaxFloat64(num, min).Message("验证不通过提示信息")
```

##### 电子邮箱格式验证

```go
vd.Email(email).Message("验证不通过提示信息")
```

##### 手机号码格式验证

```go
vd.Mobile(mobile).Message("验证不通过提示信息")
```

##### 邮政编码验证

```go
vd.ZipCode(code).Message("验证不通过提示信息")
```

##### IPV4格式验证

```go
vd.IPv4(ip).Message("验证不通过提示信息")
```


##### Domain 格式验证

```go
vd.Domain(domain).Message("验证不通过提示信息")
```


##### Date 日期格式验证

```go
vd.Date(date).Message("验证不通过提示信息")
```


##### DateTime 日期时间格式验证

```go
vd.DateTime(datetime).Message("验证不通过提示信息")
```

##### 纯字母验证

```go
vd.Letter(str).Message("验证不通过提示信息")
```

##### 数字、字母组合验证

```go
vd.Alpha(str).Message("验证不通过提示信息")
```

##### 数字、字母、下划线组合验证

```go
vd.AlphaDash(str).Message("验证不通过提示信息")
```


##### 验证是否字符串数字

```go
vd.Number(str).Message("验证不通过提示信息")
```

##### 验证是否字符串大于0正整数

```go
vd.Positive(str).Message("验证不通过提示信息")
```

##### 验证是否字符串小于0正整数

```go
vd.Negative(str).Message("验证不通过提示信息")
```

##### 验证是否字符串浮点数

```go
vd.Float(str).Message("验证不通过提示信息")
```

##### 验证是否中文字符串

```go
vd.Chinese(str).Message("验证不通过提示信息")
```

##### 验证是否中文字符串

```go
vd.Chinese(str).Message("验证不通过提示信息")
```

##### 验证字符串存在集合中

```go
vd.RangeString(str, arrString).Message("验证不通过提示信息")
```

##### 验证int型数字存在集合中

```go
vd.RangeInt(int, arrInt).Message("验证不通过提示信息")
```

##### 验证float64型数字存在集合中

```go
vd.RangeFloat(float64, arrFloat64).Message("验证不通过提示信息")
```

##### 正则表达式验证

```go
vd.Match(pattern, str).Message("验证不通过提示信息")
```

##### 布尔值验证

```go
vd.MatchBool(bool).Message("验证不通过提示信息")
```


##### 验证结果

```go
vd.Validate() // 返回 (bool, string)
```
