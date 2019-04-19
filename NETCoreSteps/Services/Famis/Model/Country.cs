using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Famis.Model
{
    public class Country
    {
        public string Id { get; set; }
        public string Name { get; set; }
        public bool ActiveFlag { get; set; }
        public string Description { get; set; }
        public string Abbreviation { get; set; }
        public DateTime UpdateDate { get; set; }
        public string UpdatedByName { get; set; }
        public string TabOrder { get; set; }
        public bool DefaultFlag { get; set; }
    }
}