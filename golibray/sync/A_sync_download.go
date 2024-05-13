//通过go使用的多线程包sync,中的waitgroup 和 锁 读写锁，实现一个本地的文件下载和上传服务

//todo:
//1. 文件下载，支持图片，视频,压缩包，
//2. 文件在上传到指定的文件目录下面的时候使用，上传到指定目录，可以选择将文件进行压缩上传

//3.分配用户权限，首先用户连接的时候需要输入系统分发的某个key，通过key来区分用户的全线，
//普通用户是只能读取文件和下载文件，管理员用户则可以进行文件的所有权限（上传，下载，修改，读取）

package syncTest

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	uploadDir = "./uploads"
	adminKey  = "admin123"
	wg        sync.WaitGroup
	lock      sync.RWMutex
)

func downw() {
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 校验权限
	if !checkPermission(r) {
		http.Error(w, "Permission denied", http.StatusForbidden)
		return
	}

	filename := r.FormValue("filename")
	filepath := filepath.Join(uploadDir, filename)

	lock.RLock()
	defer lock.RUnlock()

	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", getFileSize(file)))

	// 读取文件并写入响应
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 校验权限
	if !checkPermission(r) {
		http.Error(w, "Permission denied", http.StatusForbidden)
		return
	}

	// 解析上传的文件
	err := r.ParseMultipartForm(10 << 20) // 最大10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 遍历上传的文件
	for _, headers := range r.MultipartForm.File {
		for _, header := range headers {
			wg.Add(1)
			go func(header *multipart.FileHeader) {
				defer wg.Done()

				// 打开上传的文件
				file, err := header.Open()
				if err != nil {
					fmt.Println("Error opening uploaded file:", err)
					return
				}
				defer file.Close()

				// 创建目标文件
				dstFile, err := os.Create(filepath.Join(uploadDir, header.Filename))
				if err != nil {
					fmt.Println("Error creating destination file:", err)
					return
				}
				defer dstFile.Close()

				// 将上传的文件复制到目标文件
				_, err = io.Copy(dstFile, file)
				if err != nil {
					fmt.Println("Error copying file:", err)
					return
				}
				fmt.Println("File uploaded:", header.Filename)
			}(header)
		}
	}
	wg.Wait()
	fmt.Fprintln(w, "Upload completed")
}

func checkPermission(r *http.Request) bool {
	key := r.FormValue("key")
	return key == adminKey
}

func getFileSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
