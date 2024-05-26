using MQTTnet.Client;
using PrintService.Utils;
using System;
using System.Web;
using System.Web.Http;
using System.Web.Mvc;
using System.Web.Routing;
using System.Threading.Tasks;
using Newtonsoft.Json;
using PrintService.Models;
using System.Text;
using System.IO;

namespace PrintService
{
    public class WebApiApplication : System.Web.HttpApplication
    {
        static MQTTClient client;
        static string tmpDir = Path.Combine(HttpRuntime.AppDomainAppPath, "tmp");
        protected void Application_Start()
        {
            Start();
            AreaRegistration.RegisterAllAreas();
            GlobalConfiguration.Configure(WebApiConfig.Register);
            FilterConfig.RegisterGlobalFilters(GlobalFilters.Filters);
            RouteConfig.RegisterRoutes(RouteTable.Routes);
            
        }
        public Task HandleReceived(MqttApplicationMessageReceivedEventArgs e)
        {
            Console.WriteLine("Received application message.");
            e.DumpToConsole();
            var data = e.ApplicationMessage.PayloadSegment;
            var str = Encoding.UTF8.GetString(data.Array);
            var req= JsonConvert.DeserializeObject<PrintRequest>(str);
            if (req.TemplateFile != null)
            {
                var filePath = Path.Combine(tmpDir, $"{Guid.NewGuid().ToString()}.btw");
                if (!Directory.Exists(tmpDir))
                {
                    Directory.CreateDirectory(tmpDir);
                }

                using (var stream = new FileStream(filePath, FileMode.Create))
                {
                    stream.Write(req.TemplateFile,0,req.TemplateFile.Length);
                    stream.Flush();
                }
                string message = "";
                BartenderUtil.Print(filePath, req, ref message);
            }
            return Task.CompletedTask;
        }
        private void Start()
        {
            client = new MQTTClient(Environment.GetEnvironmentVariable("PRINT_MQTT_HOST") == null ? "127.0.0.1" : Environment.GetEnvironmentVariable("PRINT_MQTT_HOST"), 1883, HandleReceived, new string[] { "task/print" });
        }
    }
}
