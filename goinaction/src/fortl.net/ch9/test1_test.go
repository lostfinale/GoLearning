package ch9

import (
	"testing"
	"net/http"
)

const checkMark = "\u2713"
const ballotX = "\u2717"


//Go 语言的测试工具只会认为以_test.go 结尾的文件是测试文件。如果没有遵从这个约定，在包里
//运行 go test 的时候就可能会报告没有测试文件。一旦测试工具找到了测试文件，就会查找里
//面的测试函数并执行。



//一个测试函数必须是公开的函数，并且以 Test 单词开头
//而且函数的签名必须接收一个指向 testing.T 类型的指针

//基础测试(单个参数)
//表组测试（就是给一个列表的参数）
//模仿调用
//测试服务断点
func TestDownload(t *testing.T) {

	url := "http://www.baidu.com"

	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}




