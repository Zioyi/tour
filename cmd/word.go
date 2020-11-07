package cmd

import (
	"log"
	"strings"

	"github.com/Zioyi/tour/internal/word"
	"github.com/spf13/cobra"
)

// const
const (
	ModeUpper                       = iota + 1 // 单词转大写
	ModeLower                                  // 单词转小写
	ModeUnderscoreToUpperCamelCase             // 下划线单词转大写驼峰单词
	ModeUnderscoreToCLowerCamelCase            // 下划线单词转小写驼峰单词
	ModeCamelCaseToUnderscore                  // 驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToCLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
