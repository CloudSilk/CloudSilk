using System;
using System.IO;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;
using System.Web;
using System.Web.Http;
using Newtonsoft.Json;
using PrintService.Models;
using PrintService.Utils;

namespace PrintService.Controllers
{
    [RoutePrefix("api/mom/service/print")]
    public class PrintController : ApiController
    {
        static string tmpDir = Path.Combine(HttpRuntime.AppDomainAppPath, "tmp");
        static string outDir = Path.Combine(HttpRuntime.AppDomainAppPath, "result");
        [HttpPost]
        [Route("preview")]
        public async Task<HttpResponseMessage> Preview()
        {
            // 尝试从请求中获取文件和数组
            if (!Request.Content.IsMimeMultipartContent())
            {
                return Request.CreateResponse(HttpStatusCode.UnsupportedMediaType);
            }
            var provider = new MultipartMemoryStreamProvider();
            await Request.Content.ReadAsMultipartAsync(provider);
            Stream fileStream = null;
            string fileName = "";
            foreach (var content in provider.Contents)
            {
                var name = content.Headers.ContentDisposition.Name.Replace("\"", "");
                if (name == "btw")
                {
                    fileName = content.Headers.ContentDisposition.FileName;
                    fileName=fileName.Replace("\"", "");
                    fileStream = await content.ReadAsStreamAsync();
                }
            }

            if (fileStream == null)
            {
                return Request.CreateResponse(HttpStatusCode.BadRequest, "请上传模版");
            }
  
            // 这里处理文件，比如保存到服务器
            var filePath = Path.Combine(tmpDir, $"{Guid.NewGuid().ToString()}.btw");
            if (!Directory.Exists(tmpDir))
            {
                Directory.CreateDirectory(tmpDir);
            }

            using (var stream = new FileStream(filePath, FileMode.Create))
            {
                await fileStream.CopyToAsync(stream);
            }
            if (!Directory.Exists(outDir))
            {
                Directory.CreateDirectory(outDir);
            }

            var outFilePath = Path.Combine(outDir, $"{Guid.NewGuid().ToString()}.jpeg");
            BartenderUtil.Preview(filePath, outFilePath);
            // 读取图片文件到内存流中
            MemoryStream memoryStream = new MemoryStream();
            using (FileStream fileStream2 = new FileStream(outFilePath, FileMode.Open))
            {
                fileStream2.CopyTo(memoryStream);
            }

            // 准备HTTP响应，设置内容类型为JPEG
            HttpResponseMessage result = new HttpResponseMessage(HttpStatusCode.OK);
            result.Content = new ByteArrayContent(memoryStream.ToArray());
            result.Content.Headers.ContentType = new System.Net.Http.Headers.MediaTypeHeaderValue("image/jpeg");
            File.Delete(outFilePath);
            File.Delete(filePath);
            return result;
        }
        [HttpPost]
        [Route("print")]
        public async Task<HttpResponseMessage> Print()
        {
            // 尝试从请求中获取文件和数组
            if (!Request.Content.IsMimeMultipartContent())
            {
                return Request.CreateResponse(HttpStatusCode.UnsupportedMediaType);
            }
            var provider = new MultipartMemoryStreamProvider();
            await Request.Content.ReadAsMultipartAsync(provider);
            Stream fileStream = null;
            string fileName = "";
            PrintRequest parameters=null;
            foreach (var content in provider.Contents)
            {
                var name = content.Headers.ContentDisposition.Name.Replace("\"", "");
                if (name == "btw")
                {
                    fileName = content.Headers.ContentDisposition.FileName;
                    fileName = fileName.Replace("\"", "");
                    fileStream = await content.ReadAsStreamAsync();
                }else if(name == "parameters")
                {
                    var jsonArrayString = await content.ReadAsStringAsync();
                    parameters = JsonConvert.DeserializeObject<PrintRequest>(jsonArrayString);
                }
            }

            if (fileStream == null)
            {
                return Request.CreateResponse(HttpStatusCode.BadRequest, "请上传模版");
            }
            if(parameters == null)
            {
                return Request.CreateResponse(HttpStatusCode.BadRequest, "无效参数");
            }
            var filePath = Path.Combine(tmpDir, $"{Guid.NewGuid().ToString()}.btw");
            if (!Directory.Exists(tmpDir))
            {
                Directory.CreateDirectory(tmpDir);
            }

            using (var stream = new FileStream(filePath, FileMode.Create))
            {
                await fileStream.CopyToAsync(stream);
            }
            string message="";
            var result=BartenderUtil.Print(filePath, parameters,ref message);
            return Request.CreateResponse(HttpStatusCode.OK, message);
        }
    }
}
