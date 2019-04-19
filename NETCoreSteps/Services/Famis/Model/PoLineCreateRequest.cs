using Newtonsoft.Json;
using System;

namespace Famis.Model
{
    public class POLineCreateRequest
    {
        public int? POId { get; set; }
        public string PONumber { get; set; }
        public decimal QuantityOrdered { get; set; }
        public string Description { get; set; }
        public DateTimeOffset? DateRequired { get; set; }
        public int PropertyId { get; set; }
        public bool TaxableFlag { get; set; }
        public decimal TaxRate { get; set; }
        public decimal ShippingHandling { get; set; }
        public decimal UnitCost { get; set; }
        public bool OtherCostsFlag { get; set; }
        public bool MaterialsFlag { get; set; }
        public int? RequestId { get; set; }
        public int? PRLineId { get; set; }
        public int? OtherCostId { get; set; }
        public int? UOMId { get; set; }
        
        public bool ShouldSerializeUOMId() {
            return OtherCostsFlag;
        }
    }
}