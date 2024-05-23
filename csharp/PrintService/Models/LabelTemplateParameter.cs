namespace PrintService.model
{
    using System.ComponentModel.DataAnnotations;

    /// <summary>
    /// 
    /// </summary>
    public class LabelTemplateParameter
    {
        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "名称")]
        public string Name { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "预设值")]
        public string DefaultValue { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "设定值")]
        public string FixedValue { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "备注")]
        public string Remark { get; set; }
    }
}
