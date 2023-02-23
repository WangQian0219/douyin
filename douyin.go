type Video struct {
    ID         int
    Title      string
    Author     string
    Description string
    CoverURL   string
    VideoURL   string
    PublishTime time.Time
}
var videos []Video
func handleVideos(w http.ResponseWriter, r *http.Request) {
    // 显示所有的视频
    for _, video := range videos {
        fmt.Fprintf(w, "Title: %s\nAuthor: %s\nDescription: %s\nCover URL: %s\nVideo URL: %s\nPublish Time: %s\n\n",
            video.Title, video.Author, video.Description, video.CoverURL, video.VideoURL, video.PublishTime.String())
    }
}

func main() {
    // 注册路由
    http.HandleFunc("/videos", handleVideos)

    // 启动服务器
    http.ListenAndServe(":8080", nil)
}
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Upload Video</title>
</head>
<body>
    <h1>Upload Video</h1>
    <form method="post" action="/upload" enctype="multipart/form-data">
        <label>Title: <input type="text" name="title"></label><br>
        <label>Description: <textarea name="description"></textarea></label><br>
        <label>Cover: <input type="file" name="cover"></label><br>
        <label>Video: <input type="file" name="video"></label><br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
func handleUpload(w http.ResponseWriter, r *http.Request) {
    // 获取上传的视频文件
    file, header, err := r.FormFile("video")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // 保存视频文件到磁盘
    videoPath := filepath.Join("videos", header.Filename)
    f, err := os.Create(videoPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer f.Close()
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 获取视频信息
    title := r.FormValue("title")
    author := "Anonymous" // 默认匿名上传
    description := r.FormValue("description")
    cover, _, err := r.FormFile("cover")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer cover.Close()
    coverURL := "" // TODO: 保存封面图片并获取其URL
    publishTime := time.Now()

    // 保存视频信息到数组
    video := Video{
        ID:          len(videos) + 1,
        Title:       title,
        Author:      author,
        Description: description,
        CoverURL:    coverURL,
        VideoURL:    videoPath,
        PublishTime: publishTime,
    }
    videos = append(videos, video)

    // 返回上传成功的消息
    fmt.Fprintf(w, "Upload successful!")
}
func handleUserVideos(w http.ResponseWriter, r *http.Request) {
    // 获取用户ID
    userID, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 显示用户上传的视频
    for _, video := range videos {
        if video.Author == "Anonymous" && userID == 0 {
            // 如果视频是匿名上传的，并且用户没有登录，则显示该视频
            fmt.Fprintf(w, "Title: %s\nDescription: %s\nCover URL: %s\nVideo URL: %s\nPublish Time: %s\n\n",
                video.Title, video.Description, video.CoverURL, video.VideoURL, video.PublishTime.String())
        } else if video.Author != "Anonymous" && video.ID == userID {
            // 如果视频是某个用户上传的，并且用户已经登录，则显示该视频
            fmt.Fprintf(w, "Title: %s\nDescription: %%s\nCover URL: %s\nVideo URL: %s\nPublish Time: %s\n\n",
              video.Title, video.Description, video.CoverURL, video.VideoURL, video.PublishTime.String())
        }
        }
        }
                   
