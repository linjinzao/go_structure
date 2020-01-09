package handle
import(
	"net/http"
	"io/ioutil"
	"io"
	"log"
	"os"
	"time"
	"encoding/json"

	meta "filesystem/meta"
	util "filesystem/util"
)
func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		/** 输出html页面**/
		html , err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w,"inter server error")
			return
		}
		io.WriteString(w,string(html))
	}else if r.Method == "POST"{
		/** 接受文件流存储到本地**/
		file,head,err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		//创建文件元信息
		fmeta := meta.FileMeta {
			FileName : head.Filename,
			Location : "D:/GO/work/src/filesystem/file/"+ head.Filename,
			UploadAt : time.Now().Format("2006-01-02 15:04:05"),
		}

		localFile,err := os.Create(fmeta.Location)
		if err != nil {
			log.Println(err)
			return
		}
		defer localFile.Close()

		fmeta.FileSize,err = io.Copy(localFile,file)
		if err != nil {
			log.Println(err)
			return
		}
		localFile.Seek(0,0)
		fmeta.FileSha1 = util.FileSha1(localFile)

		meta.UpdateFileMeta(fmeta)
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"upload fishned")
}

func GetFileMeta(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	fileSha1 := r.Form["fileHash"][0]
	fmeta := meta.GetFileMeta(fileSha1)

	date,err := json.Marshal(fmeta)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(date)
	
}

func DownloadHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")

	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type","application/octet-stream")
	w.Header().Set("Content-Disposition","attachment;filename=\""+fm.FileName+"\"")
	log.Println(w.Header());
	w.Write(data)
}

func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")
	opType := r.Form.Get("op")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	currentMeta := meta.GetFileMeta(fsha1)

	currentMeta.FileName = newFileName

	meta.UpdateFileMeta(currentMeta)

	data, err := json.Marshal(currentMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")

	meta.DeleteFileMeta(fsha1)


	fileMeta := meta.GetFileMeta(fsha1)

	err := os.Remove(fileMeta.Location)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}