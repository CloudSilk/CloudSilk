using System.Collections.Generic;

namespace PrintService.Models
{
    public class PrintRequest
    {
        public IDictionary<string, dynamic> Parameters;
        public string PrinterName;
        public int PrintCopies;
        public string PrintJobName;
        public byte[] TemplateFile;
    }
}