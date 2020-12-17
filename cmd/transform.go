package cmd

import (
	"fmt"
	"strconv"

	"github.com/Zioyi/tour/internal/csv_generator"
	"github.com/Zioyi/tour/internal/transformer"
	"github.com/spf13/cobra"
)

var srcFileName string
var srcColumnNo int
var dstFileName string
var alphabet string = ""

var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "处理user_id数字和user_id字符串之间的转换",
	Long:  "读取csv格式文件，将user_id字符串转换成user_id数字或将user_id数字转换成user_id字符串",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var batchEncodeUserIdCmd = &cobra.Command{
	Use:   "batch_encode",
	Short: "读取源csv文件，生成目标csv文件，增加一列是user_id字母格式",
	Long:  "指定源csv文件中user_id数字格式所在的列序号，根据此在行尾生成user_id字母所在格式",
	Run: func(cmd *cobra.Command, args []string) {
		sbayt := transformer.NewSbayIdTransformer(alphabet)
		csvg := csv_generator.CsvGenerator{}
		records := csvg.Open(srcFileName)
		titleLine := append(records[0], "user_id字母格式")
		newRecords := [][]string{titleLine}
		for _, line := range records[1:] {
			userID, err := strconv.Atoi(line[srcColumnNo-1])
			userIDStr := "user_id错误"
			if err == nil {
				userIDStr = sbayt.ID2String(uint64(userID))
			}
			newLine := append(line, userIDStr)
			newRecords = append(newRecords, newLine)
		}
		csvg.Generate(newRecords, dstFileName, true)
	},
}

var batchDecodeUserIdCmd = &cobra.Command{
	Use:   "batch_decode",
	Short: "读取源csv文件，生成目标csv文件，增加一列是user_id数据格式",
	Long:  "指定源csv文件中user_id数字格式所在的列序号，根据此在行尾生成user_id字母所在格式",
	Run: func(cmd *cobra.Command, args []string) {
		sbayt := transformer.NewSbayIdTransformer(alphabet)
		csvg := csv_generator.CsvGenerator{}
		records := csvg.Open(srcFileName)
		titleLine := append(records[0], "user_ids数字格式")
		newRecords := [][]string{titleLine}
		for _, line := range records[1:] {
			userIDStr := line[srcColumnNo-1]
			userID := sbayt.String2ID(userIDStr)
			newLine := append(line, fmt.Sprint(userID))
			newRecords = append(newRecords, newLine)
		}
		csvg.Generate(newRecords, dstFileName, true)
	},
}

func init() {
	transformCmd.AddCommand(batchEncodeUserIdCmd)
	transformCmd.AddCommand(batchDecodeUserIdCmd)

	batchEncodeUserIdCmd.Flags().StringVarP(&srcFileName, "src_filename", "s", "", `指定源csv文件名，带表头`)
	batchEncodeUserIdCmd.Flags().IntVarP(&srcColumnNo, "src_column_no", "n", 0, `指定源csv文件中user_id数字类型所在的列`)
	batchEncodeUserIdCmd.Flags().StringVarP(&dstFileName, "dst_filename", "d", "", `指定生成csv文件名`)

	batchDecodeUserIdCmd.Flags().StringVarP(&srcFileName, "src_filename", "s", "", `指定源csv文件名，带表头`)
	batchDecodeUserIdCmd.Flags().IntVarP(&srcColumnNo, "src_column_no", "n", 0, `指定源csv文件中user_id数字类型所在的列`)
	batchDecodeUserIdCmd.Flags().StringVarP(&dstFileName, "dst_filename", "d", "", `指定生成csv文件名`)
}
