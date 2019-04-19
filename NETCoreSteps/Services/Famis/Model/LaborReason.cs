using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class LaborReason
    {
        public int Id { get; set; }
        public string Description { get; set; }
        public int TabOrder { get; set; }
        public bool Active { get; set; }
    }
}