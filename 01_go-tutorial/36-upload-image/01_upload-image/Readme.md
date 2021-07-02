## Logic for uploading images

**读取配置**

上传图片之前服务器会对图片做一个基本的限制，比如图片大小限制等等

**获取图片的相关信息**

- GetImageFullUrl：获取图片完整访问 URL
- GetImageName：获取图片名称
- GetImagePath：获取图片路径
- GetImageFullPath：获取图片完整路径
- CheckImageExt：检查图片后缀
- CheckImageSize：检查图片大小
- CheckImage：检查图片

**实际处理上传图片的逻辑**

- Request.FormFile：获取上传的图片（返回提供的表单键的第一个文件）
- CheckImageExt、CheckImageSize 检查图片大小，检查图片后缀
- CheckImage：检查上传图片所需（权限、文件夹）
- SaveUploadedFile：保存图片