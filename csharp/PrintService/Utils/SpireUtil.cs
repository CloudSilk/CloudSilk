namespace PrintService.Utils
{
    using Spire.Pdf;
    using System;
    using System.Drawing.Printing;
    using System.Linq;
    using System.Threading.Tasks;

    /// <summary>
    /// 
    /// </summary>
    public class SpireUtil
    {
        private class PrintResult
        {
            /// <summary>
            /// 
            /// </summary>
            public bool Printed { get; set; }
            /// <summary>
            /// 
            /// </summary>
            public string Message { get; set; }

            /// <summary>
            /// 
            /// </summary>
            public static PrintResult Default = new PrintResult()
            {
                Printed = true,
                Message = "已将打印作业成功发送到后台处理程序"
            };
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="filePath"></param>
        /// <param name="printerName"></param>
        /// <param name="printCopies"></param>
        /// <param name="message"></param>
        /// <param name="printJobName"></param>
        /// <param name="waitForCompletionTimeout"></param>
        /// <returns></returns>
        public static bool Print(string filePath, string printerName, int printCopies, ref string message, int waitForCompletionTimeout = 10000)
        {
            var printed = false;
            try
            {
                var installed = PrinterSettings.InstalledPrinters.Cast<string>().Any(c => c == printerName);
                if (!installed) throw new Exception($@"无法找到名称为{printerName}的打印机");

                var task = Task<PrintResult>.Factory.StartNew(() => {
                    var result = PrintResult.Default;
                    try
                    {
                        using (var document = new PdfDocument())
                        {
                            document.LoadFromFile(filePath);
                            document.PrintSettings.PrinterName = printerName;
                            document.PrintSettings.Copies = (short)printCopies;
                            document.PrintSettings.PrintController = new StandardPrintController();

                            document.PrintSettings.SelectSinglePageLayout(Spire.Pdf.Print.PdfSinglePageScalingMode.FitSize, true);
                            document.PrintSettings.SetPaperMargins(9, 12, 12, 12);
                            document.Print();
                        }
                    }
                    catch (Exception ex)
                    {
                        result.Printed = false;
                        result.Message = ex.Message;
                    }
                    return result;
                });

                var inLimitedTimeComplete = task.Wait(waitForCompletionTimeout);
                if (!inLimitedTimeComplete) throw new Exception("打印作业发送超时");

                message = task.Result.Message;
                printed = task.Result.Printed;
            }
            catch (Exception ex)
            {
                message = ex.Message;
            }
            return printed;
        }
    }
}
