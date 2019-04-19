using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.Serialization;
using System.Web;

namespace Famis.Model
{
    public class Receipt
    {
        
        public int Id { get; set; }
        
        public int PoId { get; set; }
        
        public string PoNumber { get; set; }
        
        public int? RequestId { get; set; }
        
        public string RequestExternalId { get; set; }
        
        public int? PoDetailId { get; set; }
        
        public int? PoLineNumber { get; set; }
        
        public string PackingSlipNumber { get; set; }
        
        public DateTimeOffset ReceiptDate { get; set; }
        
        public int? QuantityReceived { get; set; }
        
        public decimal TaxAmount { get; set; }
        
        public decimal TaxRate { get; set; }
        
        public decimal UnitCost { get; set; }
        
        public decimal ShippingAndHandling { get; set; }
        
        public decimal TotalCost { get; set; }
        
        public int? BinId { get; set; }
        
        public int? WarehouseMaterialId { get; set; }
        
        public int? OtherCostTypeId { get; set; }
        
        public decimal MarkupPercent { get; set; }
        
        public int? VendorId { get; set; }
        
        public string VendorExternalId { get; set; }
        
        public string Description { get; set; }
       
        public string Message { get; set; }
       
        public string File { get; set; }

    }
}