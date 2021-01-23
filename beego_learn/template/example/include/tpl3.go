package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

func handleMapAction(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, _ := template.ParseFiles("html/mainlayout.html","html/quote.html")

	// 设置模板数据
	data := map[string]interface{}{
		"User": "说课",
		"List": []string{"C","Go", "Python", "PHP", "JavaScript"},
	}

	fmt.Println(t.Name())

	// 渲染模板，发送响应
	t.ExecuteTemplate(w, "mainlayout", data)
}


func main() {
	// 设置 处理函数
	http.HandleFunc("/", handleMapAction)
	//// 开启监听（监听浏览器请求）
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*
*https://www.cnblogs.com/duhuo/p/5695256.html

[1]http://localhost:8081

curl是一个利用URL规则在命令行下工作的文件传输工具;
支持的通信协议有FTP、FTPS、HTTP、HTTPS、TFTP、SFTP、Gopher、SCP、Telnet、DICT、FILE、LDAP、LDAPS、IMAP、POP3、SMTP和RTSP等

执行后http://localhost:8081的html就会显示在屏幕上
[2]curl http://localhost:8081

可以使用curl的内置option:-o(小写)保存网页
[3]curl -o linux.html   http://localhost:8081
*/