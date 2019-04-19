using System;
using System.Collections.Generic;

namespace Famis.Model
{
    public class PrHeader
    {
        
        public int Id { get; set; }
        
        public string Number { get; set; }
        
        public string RequestorName { get; set; }
        
        public string RequestorEmail { get; set; }
        
        public string RequestorPhone { get; set; }
        
        public int? CreatedById { get; set; }
        
        public string CreatedByExternalId { get; set; }
        
        public string AttentionTo { get; set; }

        public string AttentionToExternalId { get; set; }

        public string Description { get; set; }
        
        public int? TypeId { get; set; }
        
        public float TotalAmount { get; set; }
        
        public DateTimeOffset CreateDate { get; set; }
        
        public DateTimeOffset UpdateDate { get; set; }
        
        public int? StatusId { get; set; }
        
        public int? PropertyId { get; set; }
        
        public string PropertyExternalId { get; set; }
        
        public string ShipToAddress { get; set; }
        
        public int? RequestId { get; set; }
        
        public string RequestExternalId { get; set; }
        
        public decimal? NTE { get; set; }
        
        public decimal? NTEApproved { get; set; }
       
        public string Message { get; set; }
       
        public string File { get; set; }

        public List<PrLine> Lines { get; set; }

        public bool ShouldSerializeId()
        {
            return (false);
        }

        public bool ShouldSerializeLines()
        {
            return (false);
        }
        public bool ShouldSerializeMessage()
        {
            return (false);
        }
        public bool ShouldSerializeFile()
        {
            return (false);
        }
        public bool ShouldSerializeVendorExternalId()
        {
            return (false);
        }
    }
}