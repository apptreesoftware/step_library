using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class Uom
    {
        public int Id { get; set; }
        public string Description { get; set; }
        public int UpdatedById { get; set; }
        public object UpdatedByIdExternal { get; set; }
        public DateTime UpdateDate { get; set; }
        public bool ActiveFlag { get; set; }
        public int TabOrder { get; set; }
    }
}