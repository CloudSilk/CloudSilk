namespace PrintService.Utils
{
    using Seagull.BarTender.Print;
    using System;
    using System.Collections.Generic;
    using System.Drawing.Printing;
    using System.Linq;
    using PrintService.Models;

    /// <summary>
    /// 
    /// </summary>
    public class BartenderUtil
    {
        private static Engine Engine = null;
        private static Object Locker = new Object();

        /// <summary>
        /// 
        /// </summary>
        /// <param name="fileName"></param>
        /// <param name="keyValues"></param>
        /// <param name="printerName"></param>
        /// <param name="printCopies"></param>
        /// <param name="message"></param>
        /// <param name="printJobName"></param>
        /// <returns></returns>
        public static bool Print(string fileName, PrintRequest printRequest, ref string message)
        {
            var printed = false;
            try
            {
                var installed = PrinterSettings.InstalledPrinters.Cast<string>().Any(c => c == printRequest.PrinterName);
                if (!installed) throw new Exception($@"无法找到名称为{printRequest.PrinterName}的打印机");

                if (null == BartenderUtil.Engine || !BartenderUtil.Engine.IsAlive)
                {
                    BartenderUtil.Engine?.Stop();
                    BartenderUtil.Engine?.Dispose();

                    BartenderUtil.Engine = new Engine(true);
                }

                //加载模板文件
                var document = BartenderUtil.Engine.Documents.FirstOrDefault(c => c.FileName == fileName);
                if (null != document)
                {
                    BartenderUtil.Engine.Documents.Close(fileName, SaveOptions.DoNotSaveChanges);
                }
                document = BartenderUtil.Engine.Documents.Open(fileName);

                //设置模板打印参数
                document.SubStrings.ForEach(subString => {
                    var defined = null != printRequest.Parameters && printRequest.Parameters.ContainsKey(subString.Name);
                    if (defined)
                    {
                        var value = printRequest.Parameters[subString.Name];
                        subString.Value = value;
                    }
                });
                //设置模板打印设置
                document.PrintSetup.PrinterName = printRequest.PrinterName;
                document.PrintSetup.IdenticalCopiesOfLabel = printRequest.PrintCopies;

                var messages = new Messages();
                var printResult = document.Print(printRequest.PrintJobName, out messages);

                //关闭模板文件
                BartenderUtil.Engine.Documents.Close(fileName, SaveOptions.DoNotSaveChanges);

                if (Result.Success != printResult)
                {
                    var text = messages.FirstOrDefault()?.Text;
                    if (!text.Contains("已将打印作业成功发送到后台处理程序")) throw new Exception(text);
                }
                printed = true;
            }
            catch (Exception ex)
            {
                message = ex.Message;
            }
            return printed;
        }

        public static bool Preview(string filePath,string imageFile)
        {
            using (var engine = new Engine(true))
            {
                var document = engine.Documents.Open(filePath);
                engine.ActiveDocument.ExportImageToFile(imageFile, ImageType.JPEG, ColorDepth.ColorDepth24bit, new Resolution(500, 500), OverwriteOptions.Overwrite);
            }
            return true;
        }
    }

}
