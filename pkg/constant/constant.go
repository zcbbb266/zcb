// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constant

// 定义了一天的秒数 SecondsOneDay。一天有 24 小时，每小时有 3600 秒，因此一天总共有 24 * 3600 秒。
const (
	SecondsOneDay = 24 * 3600
)

// 定义了一个字符串常量 StringElementPlaceHolder，其值为 __null__。通常用于在字符串中作为占位符来表示空值或缺失的元素。
const (
	StringElementPlaceHolder = `__null__`
)

// 定义了默认的分组阈值 DefaultGroupByThreshold，其值为 4。该常量通常用于配置默认情况下在分组操作中可能使用的最小阈值。
const (
	DefaultGroupByThreshold uint64 = 4
)

// 定义了一组字符串常量，这些常量表示各种原因（Reason*）和动作名称（ActionName*）。这些常量通常用于在日志记录、错误处理、状态报告等场景中，标识特定的原因或动作。
const (
	ReasonCallbackFrontendFail = "CallbackFrontendFail"
	ReasonSessionNotFound      = "SessionNotFound"
	ReasonSessionAbnormalQuit  = "SessionAbnormalQuit"
	ReasonSessionNormalQuit    = "SessionNormalQuit"
	ReasonInvalidRequest       = "InvalidRequest"
	ReasonInvalidRequestFormat = "InvalidRequestFormat"
	ReasonInvalidResponse      = "InvalidResponse"
	ReasonCallEngineFail       = "CallEngineFail"

	ActionNameSCDBQueryJobDone     = "SCDBQueryJobDone"
	ActionNameSCDBCallbackFrontend = "SCDBCallbackFrontend"
)

// 这些映射在实际编程中可以用于检查和识别代码中使用的类型别名，从而实现更灵活和语义化的类型检查和处理。
// NOTE: the type values are defined by DataType Enum in "api/v1/column.proto"
var StringTypeAlias = map[string]bool{"string": true, "str": true}

var IntegerTypeAlias = map[string]bool{"int32": true, "int64": true, "integer": true, "int": true}
var FloatTypeAlias = map[string]bool{"float": true, "float32": true}
var DoubleTypeAlias = map[string]bool{"double": true, "float64": true}
var DateTimeTypeAlias = map[string]bool{"datetime": true}
var TimeStampTypeAlias = map[string]bool{"timestamp": true}

// merge 函数接受多个 map[string]bool 类型的参数，并将它们合并成一个 map。这个函数遍历所有传入的 map，将它们的键值对合并到一个新的 map 中并返回。
func merge(ms ...map[string]bool) map[string]bool {
	res := map[string]bool{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// 定义了 SupportTypes 变量，它是通过调用 merge 函数将所有类型别名 map 合并后得到的一个 map。这个 map 包含了所有支持的类型名作为键，对应的布尔值均为 true。
var SupportTypes = merge(StringTypeAlias, IntegerTypeAlias, FloatTypeAlias, DoubleTypeAlias, DateTimeTypeAlias, TimeStampTypeAlias)

// 定义了两个哈希算法常量，GMsm3Hash 表示国密 SM3 哈希算法，Sha256Hash 表示 SHA-256 哈希算法。可以在代码中使用这些常量来指定或标识哈希算法。
const (
	GMsm3Hash  = "GMSM3"
	Sha256Hash = "SHA256"
)
