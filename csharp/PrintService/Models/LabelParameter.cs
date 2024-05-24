using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace PrintService.model
{
    public class LabelParameter : IKeyID
    {
        /// <summary>
        /// 
        /// </summary>
        [Key]
        public int ID { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "名称")]
        [StringLength(50)]
        [Required]
        public string Name { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "预设值")]
        [StringLength(1000)]
        public string DefaultValue { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "标签模板")]
        public int? LabelTemplateID { get; set; }
        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "标签模板")]
        public virtual LabelTemplate LabelTemplate { get; set; }
    }
    public class LabelTemplate : IKeyID
    {
        /// <summary>
        /// 
        /// </summary>
        [Key]
        public int ID { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "代号")]
        [StringLength(50)]
        [Required]
        public string Code { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "描述")]
        [StringLength(500)]
        public string Description { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "文件路径")]
        [StringLength(1000)]
        public string FilePath { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "备注")]
        [StringLength(500)]
        public string Remark { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "标签类型")]
        public int? LabelTypeID { get; set; }
        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "标签类型")]
        public virtual LabelType LabelType { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "标签参数")]
        public virtual ICollection<LabelParameter> LabelParameters { get; set; }
    }
    public interface IKeyID
    {
        /// <summary>
        /// 
        /// </summary>
        int ID { get; set; }

    }
    public class LabelType : IKeyID
    {
        /// <summary>
        /// 
        /// </summary>
        [Key]
        public int ID { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "类型")]
        [StringLength(50)]
        [Required]
        public string Code { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "描述")]
        [StringLength(500)]
        public string Description { get; set; }

        /// <summary>
        /// 
        /// </summary>
        [Display(Name = "备注")]
        [StringLength(500)]
        public string Remark { get; set; }
    }
}
