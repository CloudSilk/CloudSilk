using System;
using System.Web;

namespace PrintService
{
    public class StartupModule : IHttpModule
    {
        public void Init(HttpApplication context)
        {
            Console.WriteLine("应用程序已启动，模块已激活！");
        }

        public void Dispose()
        {
            // 清理代码，如果有的话
        }
    }
}