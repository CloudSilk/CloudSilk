using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace PrintService.Models
{
    public class PrintRequest
    {
        public IDictionary<string, dynamic> Parameters;
        public string PrinterName;
        public int PrintCopies;
        public string PrintJobName;
    }
}