using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{
    public class AccountSegmentValueNPFA
    {
        public int Id { get; set; }
        public string SegmentValue { get; set; }
        public bool Active { get; set; }
        public int SegmentId { get; set; }
        public string ExternalId { get; set; }
        public string SegmentExternalId { get; set; }
        public string SegmentValueDescription { get; set; }
        public DateTime? ValidFromDate { get; set; }
        public DateTime? ValidToDate { get; set; }
    }
}
